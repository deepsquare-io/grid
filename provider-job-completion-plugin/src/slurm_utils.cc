#include "slurm_utils.h"

#include <cmath>

extern "C" {
#include "src/common/gres.h"
#include "src/common/xstring.h"
}

void parse_slurm_job_info(job_record_t* job, report_t* report) {
  report->job_id = job->job_id;
  debug("%s: report->job_id %u", plugin_type, report->job_id);
  if (job->name && job->name[0])
    report->job_name = xstrdup(job->name);
  else
    report->job_name = NULL;
  debug("%s: report->qos_name %s", plugin_type, report->qos_name);
  if (job->account && job->account[0])
    report->account = xstrdup(job->account);
  else
    report->account = NULL;
  debug("%s: report->account %s", plugin_type, report->account);
  if (job->assoc_ptr && job->assoc_ptr->cluster && job->assoc_ptr->cluster[0])
    report->cluster = xstrdup(job->assoc_ptr->cluster);
  else
    report->cluster = NULL;
  debug("%s: report->cluster %s", plugin_type, report->cluster);
  if (job->part_ptr && job->part_ptr->name && job->part_ptr->name[0])
    report->partition = xstrdup(job->part_ptr->name);
  else
    report->partition = NULL;
  debug("%s: report->partition %s", plugin_type, report->partition);
  report->job_state = job->job_state;
  debug("%s: report->job_state %s", plugin_type,
        job_state_string(job->job_state));
  report->user_id = job->user_id;
  debug("%s: report->user_id %u", plugin_type, report->user_id);
  report->start_time = job->start_time;
  debug("%s: report->start_time %ld", plugin_type, report->start_time);
  report->end_time = job->end_time;
  debug("%s: report->end_time %ld", plugin_type, report->end_time);
  report->elapsed = job->end_time - job->start_time;
  debug("%s: report->elapsed %ld", plugin_type, report->elapsed);
  if (job->qos_ptr) report->usage_factor = job->qos_ptr->usage_factor;
  debug("%s: report->usage_factor %lf", plugin_type, report->usage_factor);
  if (job->qos_ptr && job->qos_ptr->name && job->qos_ptr->name[0])
    report->qos_name = xstrdup(job->qos_ptr->name);
  else
    report->qos_name = NULL;
  debug("%s: report->qos_name %s", plugin_type, report->qos_name);
  if (job->tres_alloc_cnt) {
    report->billing = job->tres_alloc_cnt[TRES_ARRAY_BILLING];
    debug("%s: report->billing %lu", plugin_type, report->billing);
    report->mem = job->tres_alloc_cnt[TRES_ARRAY_MEM];
    debug("%s: report->mem %lu", plugin_type, report->mem);
    report->cpu = job->tres_alloc_cnt[TRES_ARRAY_CPU];
    debug("%s: report->cpu %lu", plugin_type, report->cpu);
  }
  report->total_cost =
      ((uint64_t)round(((double)report->billing * (double)report->elapsed *
                        report->usage_factor) /
                       60.0l));
  debug("%s: report->total_cost %lu", plugin_type, report->total_cost);

  report->priority = job->priority;
  debug("%s: report->priority %u", plugin_type, report->priority);

  // Trying to find the gres gpu. Default to 0.
  // See: https://slurm.schedmd.com/gres_design.html
  debug("%s: gres_list_alloc", plugin_type);
  report->gpu = 0;
  if (job->gres_list_alloc && !list_is_empty(job->gres_list_alloc)) {
    gres_state_t* gres_state = NULL;
    ListIterator itr = list_iterator_create(job->gres_list_alloc);

    while ((gres_state = (gres_state_t*)list_next(itr))) {
      debug("%s: found gres %s", plugin_type, gres_state->gres_name);
      if (xstrncmp(gres_state->gres_name, "gpu", 3) == 0) {
        gres_job_state_t* gres_job_state =
            (gres_job_state_t*)gres_state->gres_data;
        report->gpu = gres_job_state->total_gres;
        debug("%s: report->gpu %lu", plugin_type, report->gpu);
        break;
      }
    }
    slurm_list_iterator_destroy(itr);
  }
}
