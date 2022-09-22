#ifndef REPORT_H
#define REPORT_H

#include <cstdint>
#include <ctime>

typedef struct report {
  /** @brief A Slurm Job ID. */
  uint32_t job_id;
  /** @brief The name of the job. */
  char *job_name;
  /** @brief A UNIX user ID. */
  uint32_t user_id;
  /** @brief A Slurm account. */
  char *account;
  /** @brief A Slurm Cluster name. */
  char *cluster;
  /** @brief A Slurm Partition name. */
  char *partition;
  /** @brief A Slurm Job state. */
  uint32_t job_state;
  /** @brief The allocated CPUs. */
  uint64_t cpu;
  /** @brief The alloctaed memory in MB. */
  uint64_t mem;
  /** @brief The allocated GPUs. */
  uint64_t gpu;
  /** @brief The billing tres factor. */
  uint64_t billing;
  /** @brief The job start timestamp. */
  time_t start_time;
  /** @brief The job end timestamp. */
  time_t end_time;
  /** @brief The job duration. */
  time_t elapsed;
  /** @brief The name of the Qos. */
  char *qos_name;
  /** @brief The usage factor of the Qos. */
  double usage_factor;
  /**
   * @brief The total cost.
   *
   * total_cost = round((usage_factor * elapsed * billing)/60.0)
   */
  uint64_t total_cost;
  /** @brief The priority of the job. */
  uint32_t priority;
} report_t;

void free_report_members(report_t *report);

#endif  // REPORT_H
