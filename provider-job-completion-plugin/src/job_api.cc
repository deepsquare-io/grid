#include "job_api.h"

extern "C" {
#include "src/common/slurm_jobcomp.h"
}

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
  bool rc = status.ok();
  if (!rc) {
    error("%s: error %d: %s", plugin_type, (int)status.error_code(),
          status.error_message().c_str());
  }
  return rc;
}
