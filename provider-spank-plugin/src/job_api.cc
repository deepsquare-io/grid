#include "job_api.h"

extern "C" {
#include <slurm/spank.h>

extern const char plugin_type[];
}

using supervisor::v1alpha1::JobStatus;
using supervisor::v1alpha1::SetJobStatusRequest;
using supervisor::v1alpha1::SetJobStatusResponse;

SetJobStatusRequest MakeSetJobRunning(const report_t& report) {
  SetJobStatusRequest req;
  req.set_name(report.job_name);
  req.set_id(report.job_id);
  req.set_duration(0);
  req.set_status(JobStatus::JOB_STATUS_RUNNING);

  return req;
}

bool JobAPIClient::SetJobStatus(const SetJobStatusRequest& req) {
  grpc::ClientContext context;

  SetJobStatusResponse rep;
  grpc::Status status = stub_->SetJobStatus(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    slurm_error("%s: error %d: %s", plugin_type, (int)status.error_code(),
                status.error_message().c_str());
  }
  return rc;
}
