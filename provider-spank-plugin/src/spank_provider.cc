extern "C" {
#include <slurm/slurm.h>
#include <slurm/spank.h>
}

#include <grpcpp/channel.h>
#include <grpcpp/create_channel.h>

#include "job_api.h"
#include "report.h"
#include "slurm_utils.h"
#include "supervisor/v1alpha1/job.grpc.pb.h"

using supervisor::v1alpha1::SetJobStatusRequest;

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

/**
 * @brief Called at the same time as the job prolog. If this function returns
 * a non-zero value and the SPANK plugin that contains it is required in the
 * plugstack.conf, the node that this is run on will be drained.
 *
 * @param spank (input) SPANK handle which must be passed back to Slurm when the
 * plugin calls functions like spank_get_item and spank_getenv.
 * @param ac Argument count
 * @param argv Argument vector
 * @return int Error code
 */
extern int slurm_spank_job_prolog(spank_t spank, int ac, char *argv[]) {
  // Fetch the job
  unsigned int jobid = 0;
  job_info_msg_t *job_info = NULL;
  if (spank_get_item(spank, S_JOB_ID, &jobid) != 0) {
    slurm_error("%s: couldn't find the job ID", plugin_type);
    return SLURM_SUCCESS;
  }
  slurm_debug("%s: start %s %d", plugin_type, __func__, jobid);
  if (slurm_load_job(&job_info, jobid, SHOW_ALL) != 0) {
    slurm_error("%s: couldn't load the job %u", plugin_type, jobid);
    return SLURM_SUCCESS;
  }

  slurm_job_info_t job = job_info->job_array[0];
  report_t report;
  parse_slurm_job_info(job, report);

  // Filter
  if (report.comment.find("supervisor") != 0) {
    slurm_debug("%s: won't report job %d", plugin_type, report.job_id);
    return SLURM_SUCCESS;
  }

  // Get endpoint
  auto endpoint = report.comment.substr(11);

  // Output
  grpc::experimental::TlsChannelCredentialsOptions options;
  options.set_verify_server_certs(false);
  JobAPIClient job_api(grpc::CreateChannel(
      endpoint, grpc::experimental::TlsCredentials(options)));

  auto req = MakeSetJobRunning(report);
  if (!job_api.SetJobStatus(req)) {
    slurm_error("%s: SetJobStatus failed", plugin_type);
  }

  return SLURM_SUCCESS;
}
}
