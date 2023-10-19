// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: supervisor/v1alpha1/ssh.proto

#include "supervisor/v1alpha1/ssh.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = ::PROTOBUF_NAMESPACE_ID::internal;
namespace supervisor {
namespace v1alpha1 {
template <typename>
PROTOBUF_CONSTEXPR FetchAuthorizedKeysRequest::FetchAuthorizedKeysRequest(
    ::_pbi::ConstantInitialized) {}
struct FetchAuthorizedKeysRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR FetchAuthorizedKeysRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~FetchAuthorizedKeysRequestDefaultTypeInternal() {}
  union {
    FetchAuthorizedKeysRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 FetchAuthorizedKeysRequestDefaultTypeInternal _FetchAuthorizedKeysRequest_default_instance_;
template <typename>
PROTOBUF_CONSTEXPR FetchAuthorizedKeysResponse::FetchAuthorizedKeysResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.authorized_keys_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_._cached_size_)*/{}} {}
struct FetchAuthorizedKeysResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR FetchAuthorizedKeysResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~FetchAuthorizedKeysResponseDefaultTypeInternal() {}
  union {
    FetchAuthorizedKeysResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 FetchAuthorizedKeysResponseDefaultTypeInternal _FetchAuthorizedKeysResponse_default_instance_;
}  // namespace v1alpha1
}  // namespace supervisor
static ::_pb::Metadata file_level_metadata_supervisor_2fv1alpha1_2fssh_2eproto[2];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_supervisor_2fv1alpha1_2fssh_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_supervisor_2fv1alpha1_2fssh_2eproto = nullptr;
const ::uint32_t TableStruct_supervisor_2fv1alpha1_2fssh_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::FetchAuthorizedKeysRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::FetchAuthorizedKeysResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::supervisor::v1alpha1::FetchAuthorizedKeysResponse, _impl_.authorized_keys_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        { 0, -1, -1, sizeof(::supervisor::v1alpha1::FetchAuthorizedKeysRequest)},
        { 8, -1, -1, sizeof(::supervisor::v1alpha1::FetchAuthorizedKeysResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::supervisor::v1alpha1::_FetchAuthorizedKeysRequest_default_instance_._instance,
    &::supervisor::v1alpha1::_FetchAuthorizedKeysResponse_default_instance_._instance,
};
const char descriptor_table_protodef_supervisor_2fv1alpha1_2fssh_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\035supervisor/v1alpha1/ssh.proto\022\023supervi"
    "sor.v1alpha1\"\034\n\032FetchAuthorizedKeysReque"
    "st\"F\n\033FetchAuthorizedKeysResponse\022\'\n\017aut"
    "horized_keys\030\001 \001(\tR\016authorizedKeys2\204\001\n\006S"
    "shAPI\022z\n\023FetchAuthorizedKeys\022/.superviso"
    "r.v1alpha1.FetchAuthorizedKeysRequest\0320."
    "supervisor.v1alpha1.FetchAuthorizedKeysR"
    "esponse\"\000B\366\001\n\027com.supervisor.v1alpha1B\010S"
    "shProtoP\001Zagithub.com/deepsquare-io/grid"
    "/supervisorapis/protos/gen/go/supervisor"
    "/v1alpha1;supervisorv1alpha1\370\001\000\242\002\003SXX\252\002\023"
    "Supervisor.V1alpha1\312\002\023Supervisor\\V1alpha"
    "1\342\002\037Supervisor\\V1alpha1\\GPBMetadata\352\002\024Su"
    "pervisor::V1alpha1b\006proto3"
};
static ::absl::once_flag descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto = {
    false,
    false,
    546,
    descriptor_table_protodef_supervisor_2fv1alpha1_2fssh_2eproto,
    "supervisor/v1alpha1/ssh.proto",
    &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_once,
    nullptr,
    0,
    2,
    schemas,
    file_default_instances,
    TableStruct_supervisor_2fv1alpha1_2fssh_2eproto::offsets,
    file_level_metadata_supervisor_2fv1alpha1_2fssh_2eproto,
    file_level_enum_descriptors_supervisor_2fv1alpha1_2fssh_2eproto,
    file_level_service_descriptors_supervisor_2fv1alpha1_2fssh_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_getter() {
  return &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_supervisor_2fv1alpha1_2fssh_2eproto(&descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto);
namespace supervisor {
namespace v1alpha1 {
// ===================================================================

class FetchAuthorizedKeysRequest::_Internal {
 public:
};

FetchAuthorizedKeysRequest::FetchAuthorizedKeysRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase(arena) {
  // @@protoc_insertion_point(arena_constructor:supervisor.v1alpha1.FetchAuthorizedKeysRequest)
}
FetchAuthorizedKeysRequest::FetchAuthorizedKeysRequest(const FetchAuthorizedKeysRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase() {
  FetchAuthorizedKeysRequest* const _this = this; (void)_this;
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:supervisor.v1alpha1.FetchAuthorizedKeysRequest)
}





const ::PROTOBUF_NAMESPACE_ID::Message::ClassData FetchAuthorizedKeysRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl,
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl,
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*FetchAuthorizedKeysRequest::GetClassData() const { return &_class_data_; }







::PROTOBUF_NAMESPACE_ID::Metadata FetchAuthorizedKeysRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_getter, &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_once,
      file_level_metadata_supervisor_2fv1alpha1_2fssh_2eproto[0]);
}
// ===================================================================

class FetchAuthorizedKeysResponse::_Internal {
 public:
};

FetchAuthorizedKeysResponse::FetchAuthorizedKeysResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
}
FetchAuthorizedKeysResponse::FetchAuthorizedKeysResponse(const FetchAuthorizedKeysResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  FetchAuthorizedKeysResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.authorized_keys_) {}

    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.authorized_keys_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.authorized_keys_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_authorized_keys().empty()) {
    _this->_impl_.authorized_keys_.Set(from._internal_authorized_keys(), _this->GetArenaForAllocation());
  }
  // @@protoc_insertion_point(copy_constructor:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
}

inline void FetchAuthorizedKeysResponse::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.authorized_keys_) {}

    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.authorized_keys_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.authorized_keys_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

FetchAuthorizedKeysResponse::~FetchAuthorizedKeysResponse() {
  // @@protoc_insertion_point(destructor:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void FetchAuthorizedKeysResponse::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.authorized_keys_.Destroy();
}

void FetchAuthorizedKeysResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void FetchAuthorizedKeysResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.authorized_keys_.ClearToEmpty();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* FetchAuthorizedKeysResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string authorized_keys = 1 [json_name = "authorizedKeys"];
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_authorized_keys();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys"));
        } else {
          goto handle_unusual;
        }
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

::uint8_t* FetchAuthorizedKeysResponse::_InternalSerialize(
    ::uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string authorized_keys = 1 [json_name = "authorizedKeys"];
  if (!this->_internal_authorized_keys().empty()) {
    const std::string& _s = this->_internal_authorized_keys();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  return target;
}

::size_t FetchAuthorizedKeysResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string authorized_keys = 1 [json_name = "authorizedKeys"];
  if (!this->_internal_authorized_keys().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_authorized_keys());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData FetchAuthorizedKeysResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    FetchAuthorizedKeysResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*FetchAuthorizedKeysResponse::GetClassData() const { return &_class_data_; }


void FetchAuthorizedKeysResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<FetchAuthorizedKeysResponse*>(&to_msg);
  auto& from = static_cast<const FetchAuthorizedKeysResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_authorized_keys().empty()) {
    _this->_internal_set_authorized_keys(from._internal_authorized_keys());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void FetchAuthorizedKeysResponse::CopyFrom(const FetchAuthorizedKeysResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool FetchAuthorizedKeysResponse::IsInitialized() const {
  return true;
}

void FetchAuthorizedKeysResponse::InternalSwap(FetchAuthorizedKeysResponse* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.authorized_keys_, lhs_arena,
                                       &other->_impl_.authorized_keys_, rhs_arena);
}

::PROTOBUF_NAMESPACE_ID::Metadata FetchAuthorizedKeysResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_getter, &descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto_once,
      file_level_metadata_supervisor_2fv1alpha1_2fssh_2eproto[1]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace v1alpha1
}  // namespace supervisor
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::supervisor::v1alpha1::FetchAuthorizedKeysRequest*
Arena::CreateMaybeMessage< ::supervisor::v1alpha1::FetchAuthorizedKeysRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::supervisor::v1alpha1::FetchAuthorizedKeysRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::supervisor::v1alpha1::FetchAuthorizedKeysResponse*
Arena::CreateMaybeMessage< ::supervisor::v1alpha1::FetchAuthorizedKeysResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::supervisor::v1alpha1::FetchAuthorizedKeysResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
