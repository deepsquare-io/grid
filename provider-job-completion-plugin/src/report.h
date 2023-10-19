#ifndef REPORT_H
#define REPORT_H

#include <cstdint>
#include <ctime>
#include <string>

typedef struct report {
  /** @brief A Slurm Job ID. */
  uint32_t job_id;
  /** @brief The name of the job. */
  std::string job_name;
  /** @brief A Slurm Job state. */
  uint32_t job_state;
  /** @brief The job start timestamp. */
  time_t start_time;
  /** @brief The job end timestamp. */
  time_t end_time;
  /** @brief The job duration. */
  time_t elapsed;
  /** @brief The job comment. Used to filter. */
  std::string comment;
  /** @brief The job exit code. */
  uint32_t exit_code;
} report_t;

#endif // REPORT_H
