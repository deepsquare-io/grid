#ifndef JOB_API_H
#define JOB_API_H

#include <grpcpp/channel.h>

#include "report.h"
#include "supervisor/v1alpha1/supervisor.grpc.pb.h"

using supervisor::v1alpha1::JobAPI;
using supervisor::v1alpha1::JobResult;
using supervisor::v1alpha1::SendJobResultRequest;

SendJobResultRequest MakeSendJobResultRequestFromReport(const report_t& report);

class JobAPIClient {
 private:
  std::unique_ptr<JobAPI::Stub> stub_;
  bool _SendJobResult(const SendJobResultRequest& req);

 public:
  JobAPIClient(std::shared_ptr<grpc::Channel> channel)
      : stub_(JobAPI::NewStub(channel)) {}

  bool SendJobResult(const SendJobResultRequest& req) {
    return _SendJobResult(req);
  }
};

#endif  // JOB_API_H
