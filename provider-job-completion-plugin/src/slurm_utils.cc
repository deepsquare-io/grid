#include "slurm_utils.h"

void parse_slurm_job_info(const job_record_t &job, report_t &report) {
  report.job_id = job.job_id;
  if (job.name && job.name[0])
    report.job_name = std::string(job.name);
  report.job_state = job.job_state;
  report.start_time = job.start_time;
  report.end_time = job.end_time;
  report.elapsed = job.end_time - job.start_time;
  if (job.comment && job.comment[0])
    report.comment = std::string(job.comment);
  report.exit_code = job.exit_code;
}
