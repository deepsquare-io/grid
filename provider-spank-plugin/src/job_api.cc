#include "job_api.h"

extern "C" {
#include <slurm/spank.h>

extern const char plugin_type[];
}

using supervisor::v1alpha1::SendJobStartRequest;
using supervisor::v1alpha1::SendJobStartResponse;

SendJobStartRequest MakeSendJobStartRequestFromReport(const report_t& report) {
  SendJobStartRequest req;
  req.set_job_name(report.job_name);
  req.set_job_id(report.job_id);

  return req;
}

bool JobAPIClient::SendJobStart(const SendJobStartRequest& req) {
  grpc::ClientContext context;

  SendJobStartResponse rep;
  grpc::Status status = stub_->SendJobStart(&context, req, &rep);
  bool rc = status.ok();
  if (!rc) {
    slurm_error("%s: error %d: %s", plugin_type, (int)status.error_code(),
                status.error_message().c_str());
  }
  return rc;
}
