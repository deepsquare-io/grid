#ifndef REPORT_H
#define REPORT_H

#include <cstdint>
#include <ctime>
#include <string>

extern "C" {
#include <slurm/slurm.h>
}

typedef struct report {
  /** @brief A Slurm Job ID. */
  uint32_t job_id;
  /** @brief The name of the job. */
  std::string job_name;
  /** @brief A Slurm Job state. */
  enum job_states job_state;
  /** @brief The job start timestamp. */
  time_t start_time;
  /** @brief The job end timestamp. */
  time_t end_time;
  /** @brief The job duration. */
  time_t elapsed;
  /** @brief The job comment. Used to filter. */
  std::string comment;
} report_t;

#endif  // REPORT_H
