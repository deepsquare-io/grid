#include "job_api.h"

extern "C" {
#include "src/common/slurm_jobcomp.h"
}

using supervisor::v1alpha1::SendJobFailRequest;
using supervisor::v1alpha1::SendJobFailResponse;
using supervisor::v1alpha1::SendJobResultRequest;
using supervisor::v1alpha1::SendJobResultResponse;

SendJobResultRequest MakeSendJobResultRequestFromReport(
    const report_t& report) {
  SendJobResultRequest req;
  req.set_job_name(report.job_name);
  req.set_job_id(report.job_id);
  req.set_job_duration(report.elapsed);

  return req;
}

SendJobFailRequest MakeSendJobFailRequestFromReport(const report_t& report) {
  SendJobFailRequest req;
  req.set_job_name(report.job_name);
  req.set_job_id(report.job_id);

  return req;
}

bool JobAPIClient::SendJobResult(const SendJobResultRequest& req) {
  grpc::ClientContext context;

  SendJobResultResponse rep;
  grpc::Status status = stub_->SendJobResult(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    error("%s: error %d: %s", plugin_type, (int)status.error_code(),
          status.error_message().c_str());
  }
  return rc;
}

bool JobAPIClient::SendJobFail(const SendJobFailRequest& req) {
  grpc::ClientContext context;

  SendJobFailResponse rep;
  grpc::Status status = stub_->SendJobFail(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    error("%s: error %d: %s", plugin_type, (int)status.error_code(),
          status.error_message().c_str());
  }
  return rc;
}
