// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: supervisor/v1alpha1/supervisor.proto

#include "supervisor/v1alpha1/supervisor.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace supervisor {
namespace v1alpha1 {
PROTOBUF_CONSTEXPR JobResult::JobResult(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.job_name_)*/{&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}}
  , /*decltype(_impl_.job_id_)*/uint64_t{0u}
  , /*decltype(_impl_.job_duration_)*/uint64_t{0u}
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct JobResultDefaultTypeInternal {
  PROTOBUF_CONSTEXPR JobResultDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~JobResultDefaultTypeInternal() {}
  union {
    JobResult _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 JobResultDefaultTypeInternal _JobResult_default_instance_;
PROTOBUF_CONSTEXPR SendJobResultRequest::SendJobResultRequest(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.job_result_)*/nullptr
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct SendJobResultRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR SendJobResultRequestDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~SendJobResultRequestDefaultTypeInternal() {}
  union {
    SendJobResultRequest _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 SendJobResultRequestDefaultTypeInternal _SendJobResultRequest_default_instance_;
PROTOBUF_CONSTEXPR SendJobResultResponse::SendJobResultResponse(
    ::_pbi::ConstantInitialized) {}
struct SendJobResultResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR SendJobResultResponseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~SendJobResultResponseDefaultTypeInternal() {}
  union {
    SendJobResultResponse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 SendJobResultResponseDefaultTypeInternal _SendJobResultResponse_default_instance_;
}  // namespace v1alpha1
}  // namespace supervisor
static ::_pb::Metadata file_level_metadata_supervisor_2fv1alpha1_2fsupervisor_2eproto[3];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_supervisor_2fv1alpha1_2fsupervisor_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_supervisor_2fv1alpha1_2fsupervisor_2eproto = nullptr;

const uint32_t TableStruct_supervisor_2fv1alpha1_2fsupervisor_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::JobResult, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::JobResult, _impl_.job_name_),
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::JobResult, _impl_.job_id_),
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::JobResult, _impl_.job_duration_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::SendJobResultRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::SendJobResultRequest, _impl_.job_result_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::SendJobResultResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::supervisor::v1alpha1::JobResult)},
  { 9, -1, -1, sizeof(::supervisor::v1alpha1::SendJobResultRequest)},
  { 16, -1, -1, sizeof(::supervisor::v1alpha1::SendJobResultResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::supervisor::v1alpha1::_JobResult_default_instance_._instance,
  &::supervisor::v1alpha1::_SendJobResultRequest_default_instance_._instance,
  &::supervisor::v1alpha1::_SendJobResultResponse_default_instance_._instance,
};

const char descriptor_table_protodef_supervisor_2fv1alpha1_2fsupervisor_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n$supervisor/v1alpha1/supervisor.proto\022\023"
  "supervisor.v1alpha1\"`\n\tJobResult\022\031\n\010job_"
  "name\030\001 \001(\tR\007jobName\022\025\n\006job_id\030\002 \001(\004R\005job"
  "Id\022!\n\014job_duration\030\003 \001(\004R\013jobDuration\"U\n"
  "\024SendJobResultRequest\022=\n\njob_result\030\001 \001("
  "\0132\036.supervisor.v1alpha1.JobResultR\tjobRe"
  "sult\"\027\n\025SendJobResultResponse2r\n\006JobAPI\022"
  "h\n\rSendJobResult\022).supervisor.v1alpha1.S"
  "endJobResultRequest\032*.supervisor.v1alpha"
  "1.SendJobResultResponse\"\000B\201\002\n\027com.superv"
  "isor.v1alpha1B\017SupervisorProtoP\001Zegithub"
  ".com/deepsquare-io/the-grid/supervisorap"
  "is/protos/gen/go/supervisor/v1alpha1;sup"
  "ervisorv1alpha1\370\001\000\242\002\003SXX\252\002\023Supervisor.V1"
  "alpha1\312\002\023Supervisor\\V1alpha1\342\002\037Superviso"
  "r\\V1alpha1\\GPBMetadata\352\002\024Supervisor::V1a"
  "lpha1b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto = {
    false, false, 653, descriptor_table_protodef_supervisor_2fv1alpha1_2fsupervisor_2eproto,
    "supervisor/v1alpha1/supervisor.proto",
    &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_once, nullptr, 0, 3,
    schemas, file_default_instances, TableStruct_supervisor_2fv1alpha1_2fsupervisor_2eproto::offsets,
    file_level_metadata_supervisor_2fv1alpha1_2fsupervisor_2eproto, file_level_enum_descriptors_supervisor_2fv1alpha1_2fsupervisor_2eproto,
    file_level_service_descriptors_supervisor_2fv1alpha1_2fsupervisor_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_getter() {
  return &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_supervisor_2fv1alpha1_2fsupervisor_2eproto(&descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto);
namespace supervisor {
namespace v1alpha1 {

// ===================================================================

class JobResult::_Internal {
 public:
};

JobResult::JobResult(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:supervisor.v1alpha1.JobResult)
}
JobResult::JobResult(const JobResult& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  JobResult* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.job_name_){}
    , decltype(_impl_.job_id_){}
    , decltype(_impl_.job_duration_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.job_name_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.job_name_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_job_name().empty()) {
    _this->_impl_.job_name_.Set(from._internal_job_name(), 
      _this->GetArenaForAllocation());
  }
  ::memcpy(&_impl_.job_id_, &from._impl_.job_id_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.job_duration_) -
    reinterpret_cast<char*>(&_impl_.job_id_)) + sizeof(_impl_.job_duration_));
  // @@protoc_insertion_point(copy_constructor:supervisor.v1alpha1.JobResult)
}

inline void JobResult::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.job_name_){}
    , decltype(_impl_.job_id_){uint64_t{0u}}
    , decltype(_impl_.job_duration_){uint64_t{0u}}
    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.job_name_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    _impl_.job_name_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

JobResult::~JobResult() {
  // @@protoc_insertion_point(destructor:supervisor.v1alpha1.JobResult)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void JobResult::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.job_name_.Destroy();
}

void JobResult::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void JobResult::Clear() {
// @@protoc_insertion_point(message_clear_start:supervisor.v1alpha1.JobResult)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.job_name_.ClearToEmpty();
  ::memset(&_impl_.job_id_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.job_duration_) -
      reinterpret_cast<char*>(&_impl_.job_id_)) + sizeof(_impl_.job_duration_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* JobResult::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string job_name = 1 [json_name = "jobName"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_job_name();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "supervisor.v1alpha1.JobResult.job_name"));
        } else
          goto handle_unusual;
        continue;
      // uint64 job_id = 2 [json_name = "jobId"];
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _impl_.job_id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // uint64 job_duration = 3 [json_name = "jobDuration"];
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 24)) {
          _impl_.job_duration_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* JobResult::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:supervisor.v1alpha1.JobResult)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string job_name = 1 [json_name = "jobName"];
  if (!this->_internal_job_name().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_job_name().data(), static_cast<int>(this->_internal_job_name().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "supervisor.v1alpha1.JobResult.job_name");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_job_name(), target);
  }

  // uint64 job_id = 2 [json_name = "jobId"];
  if (this->_internal_job_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(2, this->_internal_job_id(), target);
  }

  // uint64 job_duration = 3 [json_name = "jobDuration"];
  if (this->_internal_job_duration() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(3, this->_internal_job_duration(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:supervisor.v1alpha1.JobResult)
  return target;
}

size_t JobResult::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:supervisor.v1alpha1.JobResult)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string job_name = 1 [json_name = "jobName"];
  if (!this->_internal_job_name().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_job_name());
  }

  // uint64 job_id = 2 [json_name = "jobId"];
  if (this->_internal_job_id() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_job_id());
  }

  // uint64 job_duration = 3 [json_name = "jobDuration"];
  if (this->_internal_job_duration() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_job_duration());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData JobResult::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    JobResult::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*JobResult::GetClassData() const { return &_class_data_; }


void JobResult::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<JobResult*>(&to_msg);
  auto& from = static_cast<const JobResult&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:supervisor.v1alpha1.JobResult)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_job_name().empty()) {
    _this->_internal_set_job_name(from._internal_job_name());
  }
  if (from._internal_job_id() != 0) {
    _this->_internal_set_job_id(from._internal_job_id());
  }
  if (from._internal_job_duration() != 0) {
    _this->_internal_set_job_duration(from._internal_job_duration());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void JobResult::CopyFrom(const JobResult& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:supervisor.v1alpha1.JobResult)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool JobResult::IsInitialized() const {
  return true;
}

void JobResult::InternalSwap(JobResult* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &_impl_.job_name_, lhs_arena,
      &other->_impl_.job_name_, rhs_arena
  );
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(JobResult, _impl_.job_duration_)
      + sizeof(JobResult::_impl_.job_duration_)
      - PROTOBUF_FIELD_OFFSET(JobResult, _impl_.job_id_)>(
          reinterpret_cast<char*>(&_impl_.job_id_),
          reinterpret_cast<char*>(&other->_impl_.job_id_));
}

::PROTOBUF_NAMESPACE_ID::Metadata JobResult::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_getter, &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_once,
      file_level_metadata_supervisor_2fv1alpha1_2fsupervisor_2eproto[0]);
}

// ===================================================================

class SendJobResultRequest::_Internal {
 public:
  static const ::supervisor::v1alpha1::JobResult& job_result(const SendJobResultRequest* msg);
};

const ::supervisor::v1alpha1::JobResult&
SendJobResultRequest::_Internal::job_result(const SendJobResultRequest* msg) {
  return *msg->_impl_.job_result_;
}
SendJobResultRequest::SendJobResultRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:supervisor.v1alpha1.SendJobResultRequest)
}
SendJobResultRequest::SendJobResultRequest(const SendJobResultRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  SendJobResultRequest* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.job_result_){nullptr}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_job_result()) {
    _this->_impl_.job_result_ = new ::supervisor::v1alpha1::JobResult(*from._impl_.job_result_);
  }
  // @@protoc_insertion_point(copy_constructor:supervisor.v1alpha1.SendJobResultRequest)
}

inline void SendJobResultRequest::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.job_result_){nullptr}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

SendJobResultRequest::~SendJobResultRequest() {
  // @@protoc_insertion_point(destructor:supervisor.v1alpha1.SendJobResultRequest)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void SendJobResultRequest::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  if (this != internal_default_instance()) delete _impl_.job_result_;
}

void SendJobResultRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void SendJobResultRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:supervisor.v1alpha1.SendJobResultRequest)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  if (GetArenaForAllocation() == nullptr && _impl_.job_result_ != nullptr) {
    delete _impl_.job_result_;
  }
  _impl_.job_result_ = nullptr;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* SendJobResultRequest::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .supervisor.v1alpha1.JobResult job_result = 1 [json_name = "jobResult"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ctx->ParseMessage(_internal_mutable_job_result(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* SendJobResultRequest::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:supervisor.v1alpha1.SendJobResultRequest)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .supervisor.v1alpha1.JobResult job_result = 1 [json_name = "jobResult"];
  if (this->_internal_has_job_result()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(1, _Internal::job_result(this),
        _Internal::job_result(this).GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:supervisor.v1alpha1.SendJobResultRequest)
  return target;
}

size_t SendJobResultRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:supervisor.v1alpha1.SendJobResultRequest)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .supervisor.v1alpha1.JobResult job_result = 1 [json_name = "jobResult"];
  if (this->_internal_has_job_result()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.job_result_);
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData SendJobResultRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    SendJobResultRequest::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*SendJobResultRequest::GetClassData() const { return &_class_data_; }


void SendJobResultRequest::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<SendJobResultRequest*>(&to_msg);
  auto& from = static_cast<const SendJobResultRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:supervisor.v1alpha1.SendJobResultRequest)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_has_job_result()) {
    _this->_internal_mutable_job_result()->::supervisor::v1alpha1::JobResult::MergeFrom(
        from._internal_job_result());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void SendJobResultRequest::CopyFrom(const SendJobResultRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:supervisor.v1alpha1.SendJobResultRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool SendJobResultRequest::IsInitialized() const {
  return true;
}

void SendJobResultRequest::InternalSwap(SendJobResultRequest* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_.job_result_, other->_impl_.job_result_);
}

::PROTOBUF_NAMESPACE_ID::Metadata SendJobResultRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_getter, &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_once,
      file_level_metadata_supervisor_2fv1alpha1_2fsupervisor_2eproto[1]);
}

// ===================================================================

class SendJobResultResponse::_Internal {
 public:
};

SendJobResultResponse::SendJobResultResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase(arena, is_message_owned) {
  // @@protoc_insertion_point(arena_constructor:supervisor.v1alpha1.SendJobResultResponse)
}
SendJobResultResponse::SendJobResultResponse(const SendJobResultResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase() {
  SendJobResultResponse* const _this = this; (void)_this;
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:supervisor.v1alpha1.SendJobResultResponse)
}





const ::PROTOBUF_NAMESPACE_ID::Message::ClassData SendJobResultResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl,
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl,
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*SendJobResultResponse::GetClassData() const { return &_class_data_; }







::PROTOBUF_NAMESPACE_ID::Metadata SendJobResultResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_getter, &descriptor_table_supervisor_2fv1alpha1_2fsupervisor_2eproto_once,
      file_level_metadata_supervisor_2fv1alpha1_2fsupervisor_2eproto[2]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1alpha1
}  // namespace supervisor
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::supervisor::v1alpha1::JobResult*
Arena::CreateMaybeMessage< ::supervisor::v1alpha1::JobResult >(Arena* arena) {
  return Arena::CreateMessageInternal< ::supervisor::v1alpha1::JobResult >(arena);
}
template<> PROTOBUF_NOINLINE ::supervisor::v1alpha1::SendJobResultRequest*
Arena::CreateMaybeMessage< ::supervisor::v1alpha1::SendJobResultRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::supervisor::v1alpha1::SendJobResultRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::supervisor::v1alpha1::SendJobResultResponse*
Arena::CreateMaybeMessage< ::supervisor::v1alpha1::SendJobResultResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::supervisor::v1alpha1::SendJobResultResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
