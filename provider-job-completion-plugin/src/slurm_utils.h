#ifndef SLURM_UTIL_H
#define SLURM_UTIL_H

#include "report.h"

extern "C" {
#include "src/interfaces/jobcomp.h"
}

/**
 * @brief Parse the job info and put in the report struct.
 *
 * @param job (input) The slurm job info.
 * @param report (output) The report.
 * @return int Error code.
 */
void parse_slurm_job_info(const job_record_t &job, report_t &report);

#endif // SLURM_UTIL_H
