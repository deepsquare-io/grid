// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: supervisor/v1alpha1/ssh.proto

#include "supervisor/v1alpha1/ssh.pb.h"
#include "supervisor/v1alpha1/ssh.grpc.pb.h"

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

static const char* SshAPI_method_names[] = {
  "/supervisor.v1alpha1.SshAPI/FetchAuthorizedKeys",
};

std::unique_ptr< SshAPI::Stub> SshAPI::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< SshAPI::Stub> stub(new SshAPI::Stub(channel, options));
  return stub;
}

SshAPI::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_FetchAuthorizedKeys_(SshAPI_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status SshAPI::Stub::FetchAuthorizedKeys(::grpc::ClientContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest& request, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::supervisor::v1alpha1::FetchAuthorizedKeysRequest, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_FetchAuthorizedKeys_, context, request, response);
}

void SshAPI::Stub::async::FetchAuthorizedKeys(::grpc::ClientContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest* request, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::supervisor::v1alpha1::FetchAuthorizedKeysRequest, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_FetchAuthorizedKeys_, context, request, response, std::move(f));
}

void SshAPI::Stub::async::FetchAuthorizedKeys(::grpc::ClientContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest* request, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_FetchAuthorizedKeys_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::FetchAuthorizedKeysResponse>* SshAPI::Stub::PrepareAsyncFetchAuthorizedKeysRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::supervisor::v1alpha1::FetchAuthorizedKeysResponse, ::supervisor::v1alpha1::FetchAuthorizedKeysRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_FetchAuthorizedKeys_, context, request);
}

::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::FetchAuthorizedKeysResponse>* SshAPI::Stub::AsyncFetchAuthorizedKeysRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncFetchAuthorizedKeysRaw(context, request, cq);
  result->StartCall();
  return result;
}

SshAPI::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      SshAPI_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< SshAPI::Service, ::supervisor::v1alpha1::FetchAuthorizedKeysRequest, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](SshAPI::Service* service,
             ::grpc::ServerContext* ctx,
             const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest* req,
             ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* resp) {
               return service->FetchAuthorizedKeys(ctx, req, resp);
             }, this)));
}

SshAPI::Service::~Service() {
}

::grpc::Status SshAPI::Service::FetchAuthorizedKeys(::grpc::ServerContext* context, const ::supervisor::v1alpha1::FetchAuthorizedKeysRequest* request, ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace supervisor
}  // namespace v1alpha1

