#include "report.h"

extern "C" {
#include "src/common/slurm_jobcomp.h"
}

void free_report_members(report_t *report) {
  if (report->job_name) xfree(report->job_name);
  if (report->account) xfree(report->account);
  if (report->cluster) xfree(report->cluster);
  if (report->qos_name) xfree(report->qos_name);
  if (report->partition) xfree(report->partition);
}
