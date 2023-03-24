#include "job_api.h"

extern "C" {
#include "src/common/slurm_jobcomp.h"
}

using supervisor::v1alpha1::JobStatus;
using supervisor::v1alpha1::SetJobStatusRequest;
using supervisor::v1alpha1::SetJobStatusResponse;

SetJobStatusRequest MakeSetJobStatusRequest(const report_t &report) {
  SetJobStatusRequest req;
  req.set_name(report.job_name);
  req.set_id(report.job_id);
  req.set_duration(report.elapsed);

  switch (report.job_state & JOB_STATE_BASE) {
  case JOB_FAILED:
    req.set_status(JobStatus::JOB_STATUS_FAILED);
    break;
  case JOB_TIMEOUT:
    req.set_status(JobStatus::JOB_STATUS_OUT_OF_CREDITS);
    break;
  case JOB_CANCELLED:
    req.set_status(JobStatus::JOB_STATUS_CANCELLED);
    break;
  default:
    req.set_status(JobStatus::JOB_STATUS_FINISHED);
    break;
  }

  return req;
}

bool JobAPIClient::SetJobStatus(const SetJobStatusRequest &req) {
  grpc::ClientContext context;
  auto deadline = std::chrono::system_clock::now() + std::chrono::seconds(30);
  context.set_deadline(deadline);

  SetJobStatusResponse rep;
  grpc::Status status = stub_->SetJobStatus(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    error("%s: error %d: %s", plugin_type, (int)status.error_code(),
          status.error_message().c_str());
  }
  return rc;
}
