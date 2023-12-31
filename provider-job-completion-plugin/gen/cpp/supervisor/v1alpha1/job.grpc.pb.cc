// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: supervisor/v1alpha1/job.proto

#include "supervisor/v1alpha1/job.pb.h"
#include "supervisor/v1alpha1/job.grpc.pb.h"

#include <functional>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/impl/channel_interface.h>
#include <grpcpp/impl/client_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/rpc_service_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/sync_stream.h>
namespace supervisor {
namespace v1alpha1 {

static const char* JobAPI_method_names[] = {
  "/supervisor.v1alpha1.JobAPI/SetJobStatus",
};

std::unique_ptr< JobAPI::Stub> JobAPI::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< JobAPI::Stub> stub(new JobAPI::Stub(channel, options));
  return stub;
}

JobAPI::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_SetJobStatus_(JobAPI_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status JobAPI::Stub::SetJobStatus(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest& request, ::supervisor::v1alpha1::SetJobStatusResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::supervisor::v1alpha1::SetJobStatusRequest, ::supervisor::v1alpha1::SetJobStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_SetJobStatus_, context, request, response);
}

void JobAPI::Stub::async::SetJobStatus(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest* request, ::supervisor::v1alpha1::SetJobStatusResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::supervisor::v1alpha1::SetJobStatusRequest, ::supervisor::v1alpha1::SetJobStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetJobStatus_, context, request, response, std::move(f));
}

void JobAPI::Stub::async::SetJobStatus(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest* request, ::supervisor::v1alpha1::SetJobStatusResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_SetJobStatus_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SetJobStatusResponse>* JobAPI::Stub::PrepareAsyncSetJobStatusRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::supervisor::v1alpha1::SetJobStatusResponse, ::supervisor::v1alpha1::SetJobStatusRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_SetJobStatus_, context, request);
}

::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SetJobStatusResponse>* JobAPI::Stub::AsyncSetJobStatusRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetJobStatusRaw(context, request, cq);
  result->StartCall();
  return result;
}

JobAPI::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      JobAPI_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< JobAPI::Service, ::supervisor::v1alpha1::SetJobStatusRequest, ::supervisor::v1alpha1::SetJobStatusResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](JobAPI::Service* service,
             ::grpc::ServerContext* ctx,
             const ::supervisor::v1alpha1::SetJobStatusRequest* req,
             ::supervisor::v1alpha1::SetJobStatusResponse* resp) {
               return service->SetJobStatus(ctx, req, resp);
             }, this)));
}

JobAPI::Service::~Service() {
}

::grpc::Status JobAPI::Service::SetJobStatus(::grpc::ServerContext* context, const ::supervisor::v1alpha1::SetJobStatusRequest* request, ::supervisor::v1alpha1::SetJobStatusResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace supervisor
}  // namespace v1alpha1

