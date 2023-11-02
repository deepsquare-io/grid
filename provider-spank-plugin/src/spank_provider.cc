extern "C" {
#include <slurm/slurm.h>
#include <slurm/spank.h>
}

#include <grpcpp/channel.h>
#include <grpcpp/create_channel.h>

#include "job_api.h"
#include "report.h"
#include "slurm_utils.h"

extern "C" {

extern const char plugin_name[] = "provider";
extern const char plugin_type[] = "spank";
extern const unsigned int plugin_version = SLURM_VERSION_NUMBER;
extern const unsigned int spank_plugin_version = SLURM_VERSION_NUMBER;

/*
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
int slurm_spank_job_prolog(spank_t spank, int ac, char *argv[]) {
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
