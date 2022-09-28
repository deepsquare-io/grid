#include <grpcpp/channel.h>
#include <grpcpp/create_channel.h>

#include "job_api.h"
#include "report.h"
extern "C" {
#include "src/common/slurm_jobcomp.h"
}
#include "slurm_utils.h"
#include "supervisor/v1alpha1/supervisor.grpc.pb.h"

using supervisor::v1alpha1::SendJobResultRequest;

extern "C" {
/*
 * These variables are required by the generic plugin interface.  If they
 * are not found in the plugin, the plugin loader will ignore it.
 *
 * plugin_name - a string giving a human-readable description of the
 * plugin.  There is no maximum length, but the symbol must refer to
 * a valid string.
 *
 * plugin_type - a string suggesting the type of the plugin or its
 * applicability to a particular form of data or method of data handling.
 * If the low-level plugin API is used, the contents of this string are
 * unimportant and may be anything.  Slurm uses the higher-level plugin
 * interface which requires this string to be of the form
 *
 *	<application>/<method>
 *
 * where <application> is a description of the intended application of
 * the plugin (e.g., "jobcomp" for Slurm job completion logging) and <method>
 * is a description of how this plugin satisfies that application.  Slurm will
 * only load job completion logging plugins if the plugin_type string has a
 * prefix of "jobcomp/".
 *
 * plugin_version - an unsigned 32-bit integer containing the Slurm version
 * (major.minor.micro combined into a single number).
 */
extern const char plugin_name[] = "Job completion plugin for providers";
extern const char plugin_type[] = "jobcomp/provider";
extern const uint32_t plugin_version = SLURM_VERSION_NUMBER;

/* File descriptor used for logging */
static char *report_url = NULL;

/**
 * @brief Called when the plugin is loaded, before any other functions are
 * called. Put global initialization here.
 *
 * @return int SLURM_SUCCESS on success, or SLURM_ERROR on failure.
 */
extern int init(void) {
  slurm_info("%s: Initializing %s", plugin_type, plugin_name);
  return SLURM_SUCCESS;
}

/**
 * @brief Called when the plugin is removed. Clear any allocated storage here.
 *
 * @return int SLURM_SUCCESS on success, or SLURM_ERROR on failure.
 */
extern int fini(void) {
  slurm_info("%s: Finishing %s", plugin_type, plugin_name);
  xfree(report_url);
  return SLURM_SUCCESS;
}

/**
 * @brief Specify the location to be used for job logging.
 *
 * @param location (input) specification of where logging should be done. The
 * interpretation of this string is at the discretion of the plugin
 * implementation.
 * @return int SLURM_SUCCESS if successful. On failure, the plugin should return
 * SLURM_ERROR and set the errno to an appropriate value to indicate the reason
 * for failure.
 */
extern int jobcomp_p_set_location(char *location) {
  slurm_info("%s: Set location %s", plugin_type, location);
  int rc = SLURM_SUCCESS;

  if (location == NULL || location[0] == '\0') {
    return error("%s: JobCompLoc was either not set or blank", plugin_type);
  }
  xfree(report_url);
  report_url = xstrdup(location);

  return rc;
}

/**
 * @brief Note that a job is about to terminate or change size. The job's state
 * will include the JOB_RESIZING flag if and only if it is about to change size.
 * Otherwise the job is terminating. Note the existence of resize_time in the
 * job record if one wishes to record information about a job at each size (i.e.
 * a history of the job as its size changes through time).
 *
 * @param job_ptr (input) Pointer to job record as defined in
 * src/slurmctld/slurmctld.h
 * @return int SLURM_SUCCESS if successful. On failure, the plugin should return
 * SLURM_ERROR and set the errno to an appropriate value to indicate the reason
 * for failure.
 */
extern int jobcomp_p_log_record(job_record_t *job_ptr) {
  debug("%s: start %s %u", plugin_type, __func__, job_ptr->job_id);
  if (job_ptr == NULL) return error("%s: job_ptr is NULL", plugin_type);

  // Assert the job state
  if (!IS_JOB_COMPLETE(job_ptr) && !IS_JOB_TIMEOUT(job_ptr) &&
      !IS_JOB_FAILED(job_ptr) && !IS_JOB_COMPLETING(job_ptr)) {
    debug(
        "%s: job %u is not COMPLETED but was %s, "
        "ignoring...",
        plugin_type, job_ptr->job_id, job_state_string(job_ptr->job_state));
    return SLURM_SUCCESS;
  }
  int rc = SLURM_SUCCESS;

  debug("%s: fetch report", plugin_type);

  // Parsing the job_ptr
  report_t report;
  parse_slurm_job_info(job_ptr, &report);

  JobAPIClient job_api(
      grpc::CreateChannel(report_url, grpc::InsecureChannelCredentials()));

  auto req = MakeSendJobResultRequestFromReport(report);
  if (!job_api.SendJobResult(req)) {
    error("%s: publish failed", plugin_type);
  }

  free_report_members(&report);

  debug("%s: end %s %u", plugin_type, __func__, job_ptr->job_id);
  return rc;
}

/**
 * @brief Get completed job info from storage.
 *
 * @param job_cond (input) specification of filters to identify the jobs we
 * wish information about (start time, end time, cluster name, user id, etc).
 * acct_job_cond_t is defined in common/slurm_accounting_storage.h.
 * @return List A list of job records or NULL on error. Elements on the list
 * are of type jobcomp_job_rec_t, which is defined in common/slurm_jobcomp.h.
 * Any returned list must be destroyed to avoid memory leaks.
 */
extern List jobcomp_p_get_jobs(void *job_cond) {
  debug("%s: %s function is not implemented", plugin_type, __func__);
  return NULL;
}

/**
 * @brief Used to archive old data.
 *
 * @param arch_cond
 * @return int Error number for the last failure encountered by the job
 * completion plugin.
 */
extern int jobcomp_p_archive(void *arch_cond) {
  debug("%s: %s function is not implemented", plugin_type, __func__);
  return SLURM_SUCCESS;
}
}
