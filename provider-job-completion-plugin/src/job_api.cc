#include "job_api.h"

using supervisor::v1alpha1::SendJobResultRequest;
using supervisor::v1alpha1::SendJobResultResponse;

SendJobResultRequest MakeSendJobResultRequestFromReport(
    const report_t& report) {
  SendJobResultRequest req;
  auto r = req.mutable_job_result();
  r->set_job_name(report.job_name);
  r->set_job_id(report.job_id);
  r->set_job_duration(report.elapsed);

  return req;
}

bool JobAPIClient::_SendJobResult(const SendJobResultRequest& req) {
  grpc::ClientContext context;

  SendJobResultResponse rep;
  grpc::Status status = stub_->SendJobResult(&context, req, &rep);
  return status.ok();
}
