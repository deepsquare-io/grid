#include "job_api.h"

extern "C" {
#include "src/common/slurm_jobcomp.h"
}

using supervisor::v1alpha1::SendJobFailedRequest;
using supervisor::v1alpha1::SendJobFailedResponse;
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

SendJobFailedRequest MakeSendJobFailedRequestFromReport(
    const report_t& report) {
  SendJobFailedRequest req;
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

bool JobAPIClient::SendJobFailed(const SendJobFailedRequest& req) {
  grpc::ClientContext context;

  SendJobFailedResponse rep;
  grpc::Status status = stub_->SendJobFailed(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    error("%s: error %d: %s", plugin_type, (int)status.error_code(),
          status.error_message().c_str());
  }
  return rc;
}
