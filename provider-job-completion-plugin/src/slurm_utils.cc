#include "slurm_utils.h"

void parse_slurm_job_info(const job_record_t& job, report_t& report) {
  report.job_id = job.job_id;
  debug("%s: report.job_id %u", plugin_type, report.job_id);

  if (job.name && job.name[0]) report.job_name = std::string(job.name);
  debug("%s: report.job_name %s", plugin_type, report.job_name.c_str());

  report.job_state = job.job_state;
  debug("%s: report.job_state %s", plugin_type,
        job_state_string(job.job_state));

  report.start_time = job.start_time;
  debug("%s: report.start_time %ld", plugin_type, report.start_time);

  report.end_time = job.end_time;
  debug("%s: report.end_time %ld", plugin_type, report.end_time);

  report.elapsed = job.end_time - job.start_time;
  debug("%s: report.elapsed %ld", plugin_type, report.elapsed);
}
