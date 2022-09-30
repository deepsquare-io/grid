// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: supervisor/v1alpha1/job.proto
#ifndef GRPC_supervisor_2fv1alpha1_2fjob_2eproto__INCLUDED
#define GRPC_supervisor_2fv1alpha1_2fjob_2eproto__INCLUDED

#include "supervisor/v1alpha1/job.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_generic_service.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/client_context.h>
#include <grpcpp/impl/codegen/completion_queue.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace supervisor {
namespace v1alpha1 {

// A job handler API
class JobAPI final {
 public:
  static constexpr char const* service_full_name() {
    return "supervisor.v1alpha1.JobAPI";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::supervisor::v1alpha1::SendJobStartResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>> AsyncSendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>>(AsyncSendJobStartRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>> PrepareAsyncSendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>>(PrepareAsyncSendJobStartRaw(context, request, cq));
    }
    virtual ::grpc::Status SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::supervisor::v1alpha1::SendJobResultResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>> AsyncSendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>>(AsyncSendJobResultRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>> PrepareAsyncSendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>>(PrepareAsyncSendJobResultRaw(context, request, cq));
    }
    virtual ::grpc::Status SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::supervisor::v1alpha1::SendJobFailResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>> AsyncSendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>>(AsyncSendJobFailRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>> PrepareAsyncSendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>>(PrepareAsyncSendJobFailRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>* AsyncSendJobStartRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobStartResponse>* PrepareAsyncSendJobStartRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>* AsyncSendJobResultRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobResultResponse>* PrepareAsyncSendJobResultRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>* AsyncSendJobFailRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::supervisor::v1alpha1::SendJobFailResponse>* PrepareAsyncSendJobFailRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::supervisor::v1alpha1::SendJobStartResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>> AsyncSendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>>(AsyncSendJobStartRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>> PrepareAsyncSendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>>(PrepareAsyncSendJobStartRaw(context, request, cq));
    }
    ::grpc::Status SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::supervisor::v1alpha1::SendJobResultResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>> AsyncSendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>>(AsyncSendJobResultRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>> PrepareAsyncSendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>>(PrepareAsyncSendJobResultRaw(context, request, cq));
    }
    ::grpc::Status SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::supervisor::v1alpha1::SendJobFailResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>> AsyncSendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>>(AsyncSendJobFailRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>> PrepareAsyncSendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>>(PrepareAsyncSendJobFailRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response, std::function<void(::grpc::Status)>) override;
      void SendJobStart(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response, std::function<void(::grpc::Status)>) override;
      void SendJobResult(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response, std::function<void(::grpc::Status)>) override;
      void SendJobFail(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>* AsyncSendJobStartRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobStartResponse>* PrepareAsyncSendJobStartRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobStartRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>* AsyncSendJobResultRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobResultResponse>* PrepareAsyncSendJobResultRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobResultRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>* AsyncSendJobFailRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::supervisor::v1alpha1::SendJobFailResponse>* PrepareAsyncSendJobFailRaw(::grpc::ClientContext* context, const ::supervisor::v1alpha1::SendJobFailRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_SendJobStart_;
    const ::grpc::internal::RpcMethod rpcmethod_SendJobResult_;
    const ::grpc::internal::RpcMethod rpcmethod_SendJobFail_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status SendJobStart(::grpc::ServerContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response);
    virtual ::grpc::Status SendJobResult(::grpc::ServerContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response);
    virtual ::grpc::Status SendJobFail(::grpc::ServerContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_SendJobStart() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobStart(::grpc::ServerContext* context, ::supervisor::v1alpha1::SendJobStartRequest* request, ::grpc::ServerAsyncResponseWriter< ::supervisor::v1alpha1::SendJobStartResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_SendJobResult() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobResult(::grpc::ServerContext* context, ::supervisor::v1alpha1::SendJobResultRequest* request, ::grpc::ServerAsyncResponseWriter< ::supervisor::v1alpha1::SendJobResultResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_SendJobFail() {
      ::grpc::Service::MarkMethodAsync(2);
    }
    ~WithAsyncMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobFail(::grpc::ServerContext* context, ::supervisor::v1alpha1::SendJobFailRequest* request, ::grpc::ServerAsyncResponseWriter< ::supervisor::v1alpha1::SendJobFailResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(2, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_SendJobStart<WithAsyncMethod_SendJobResult<WithAsyncMethod_SendJobFail<Service > > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_SendJobStart() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobStartRequest, ::supervisor::v1alpha1::SendJobStartResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::supervisor::v1alpha1::SendJobStartRequest* request, ::supervisor::v1alpha1::SendJobStartResponse* response) { return this->SendJobStart(context, request, response); }));}
    void SetMessageAllocatorFor_SendJobStart(
        ::grpc::MessageAllocator< ::supervisor::v1alpha1::SendJobStartRequest, ::supervisor::v1alpha1::SendJobStartResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobStartRequest, ::supervisor::v1alpha1::SendJobStartResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobStart(
      ::grpc::CallbackServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_SendJobResult() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobResultRequest, ::supervisor::v1alpha1::SendJobResultResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::supervisor::v1alpha1::SendJobResultRequest* request, ::supervisor::v1alpha1::SendJobResultResponse* response) { return this->SendJobResult(context, request, response); }));}
    void SetMessageAllocatorFor_SendJobResult(
        ::grpc::MessageAllocator< ::supervisor::v1alpha1::SendJobResultRequest, ::supervisor::v1alpha1::SendJobResultResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobResultRequest, ::supervisor::v1alpha1::SendJobResultResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobResult(
      ::grpc::CallbackServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_SendJobFail() {
      ::grpc::Service::MarkMethodCallback(2,
          new ::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobFailRequest, ::supervisor::v1alpha1::SendJobFailResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::supervisor::v1alpha1::SendJobFailRequest* request, ::supervisor::v1alpha1::SendJobFailResponse* response) { return this->SendJobFail(context, request, response); }));}
    void SetMessageAllocatorFor_SendJobFail(
        ::grpc::MessageAllocator< ::supervisor::v1alpha1::SendJobFailRequest, ::supervisor::v1alpha1::SendJobFailResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(2);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::supervisor::v1alpha1::SendJobFailRequest, ::supervisor::v1alpha1::SendJobFailResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobFail(
      ::grpc::CallbackServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_SendJobStart<WithCallbackMethod_SendJobResult<WithCallbackMethod_SendJobFail<Service > > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_SendJobStart() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_SendJobResult() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_SendJobFail() {
      ::grpc::Service::MarkMethodGeneric(2);
    }
    ~WithGenericMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_SendJobStart() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobStart(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_SendJobResult() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobResult(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_SendJobFail() {
      ::grpc::Service::MarkMethodRaw(2);
    }
    ~WithRawMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendJobFail(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(2, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_SendJobStart() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->SendJobStart(context, request, response); }));
    }
    ~WithRawCallbackMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobStart(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_SendJobResult() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->SendJobResult(context, request, response); }));
    }
    ~WithRawCallbackMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobResult(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_SendJobFail() {
      ::grpc::Service::MarkMethodRawCallback(2,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->SendJobFail(context, request, response); }));
    }
    ~WithRawCallbackMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendJobFail(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SendJobStart : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_SendJobStart() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::supervisor::v1alpha1::SendJobStartRequest, ::supervisor::v1alpha1::SendJobStartResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::supervisor::v1alpha1::SendJobStartRequest, ::supervisor::v1alpha1::SendJobStartResponse>* streamer) {
                       return this->StreamedSendJobStart(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_SendJobStart() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SendJobStart(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobStartRequest* /*request*/, ::supervisor::v1alpha1::SendJobStartResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSendJobStart(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::supervisor::v1alpha1::SendJobStartRequest,::supervisor::v1alpha1::SendJobStartResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SendJobResult : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_SendJobResult() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::supervisor::v1alpha1::SendJobResultRequest, ::supervisor::v1alpha1::SendJobResultResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::supervisor::v1alpha1::SendJobResultRequest, ::supervisor::v1alpha1::SendJobResultResponse>* streamer) {
                       return this->StreamedSendJobResult(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_SendJobResult() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SendJobResult(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobResultRequest* /*request*/, ::supervisor::v1alpha1::SendJobResultResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSendJobResult(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::supervisor::v1alpha1::SendJobResultRequest,::supervisor::v1alpha1::SendJobResultResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SendJobFail : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_SendJobFail() {
      ::grpc::Service::MarkMethodStreamed(2,
        new ::grpc::internal::StreamedUnaryHandler<
          ::supervisor::v1alpha1::SendJobFailRequest, ::supervisor::v1alpha1::SendJobFailResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::supervisor::v1alpha1::SendJobFailRequest, ::supervisor::v1alpha1::SendJobFailResponse>* streamer) {
                       return this->StreamedSendJobFail(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_SendJobFail() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SendJobFail(::grpc::ServerContext* /*context*/, const ::supervisor::v1alpha1::SendJobFailRequest* /*request*/, ::supervisor::v1alpha1::SendJobFailResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSendJobFail(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::supervisor::v1alpha1::SendJobFailRequest,::supervisor::v1alpha1::SendJobFailResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_SendJobStart<WithStreamedUnaryMethod_SendJobResult<WithStreamedUnaryMethod_SendJobFail<Service > > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_SendJobStart<WithStreamedUnaryMethod_SendJobResult<WithStreamedUnaryMethod_SendJobFail<Service > > > StreamedService;
};

}  // namespace v1alpha1
}  // namespace supervisor


#endif  // GRPC_supervisor_2fv1alpha1_2fjob_2eproto__INCLUDED
