// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: supervisor/v1alpha1/ssh.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_supervisor_2fv1alpha1_2fssh_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_supervisor_2fv1alpha1_2fssh_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021006 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_bases.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_supervisor_2fv1alpha1_2fssh_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_supervisor_2fv1alpha1_2fssh_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_supervisor_2fv1alpha1_2fssh_2eproto;
namespace supervisor {
namespace v1alpha1 {
class FetchAuthorizedKeysRequest;
struct FetchAuthorizedKeysRequestDefaultTypeInternal;
extern FetchAuthorizedKeysRequestDefaultTypeInternal _FetchAuthorizedKeysRequest_default_instance_;
class FetchAuthorizedKeysResponse;
struct FetchAuthorizedKeysResponseDefaultTypeInternal;
extern FetchAuthorizedKeysResponseDefaultTypeInternal _FetchAuthorizedKeysResponse_default_instance_;
}  // namespace v1alpha1
}  // namespace supervisor
PROTOBUF_NAMESPACE_OPEN
template<> ::supervisor::v1alpha1::FetchAuthorizedKeysRequest* Arena::CreateMaybeMessage<::supervisor::v1alpha1::FetchAuthorizedKeysRequest>(Arena*);
template<> ::supervisor::v1alpha1::FetchAuthorizedKeysResponse* Arena::CreateMaybeMessage<::supervisor::v1alpha1::FetchAuthorizedKeysResponse>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace supervisor {
namespace v1alpha1 {

// ===================================================================

class FetchAuthorizedKeysRequest final :
    public ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase /* @@protoc_insertion_point(class_definition:supervisor.v1alpha1.FetchAuthorizedKeysRequest) */ {
 public:
  inline FetchAuthorizedKeysRequest() : FetchAuthorizedKeysRequest(nullptr) {}
  explicit PROTOBUF_CONSTEXPR FetchAuthorizedKeysRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  FetchAuthorizedKeysRequest(const FetchAuthorizedKeysRequest& from);
  FetchAuthorizedKeysRequest(FetchAuthorizedKeysRequest&& from) noexcept
    : FetchAuthorizedKeysRequest() {
    *this = ::std::move(from);
  }

  inline FetchAuthorizedKeysRequest& operator=(const FetchAuthorizedKeysRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline FetchAuthorizedKeysRequest& operator=(FetchAuthorizedKeysRequest&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const FetchAuthorizedKeysRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const FetchAuthorizedKeysRequest* internal_default_instance() {
    return reinterpret_cast<const FetchAuthorizedKeysRequest*>(
               &_FetchAuthorizedKeysRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(FetchAuthorizedKeysRequest& a, FetchAuthorizedKeysRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(FetchAuthorizedKeysRequest* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(FetchAuthorizedKeysRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  FetchAuthorizedKeysRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<FetchAuthorizedKeysRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyFrom;
  inline void CopyFrom(const FetchAuthorizedKeysRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::CopyImpl(*this, from);
  }
  using ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeFrom;
  void MergeFrom(const FetchAuthorizedKeysRequest& from) {
    ::PROTOBUF_NAMESPACE_ID::internal::ZeroFieldsBase::MergeImpl(*this, from);
  }
  public:

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "supervisor.v1alpha1.FetchAuthorizedKeysRequest";
  }
  protected:
  explicit FetchAuthorizedKeysRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // @@protoc_insertion_point(class_scope:supervisor.v1alpha1.FetchAuthorizedKeysRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
  };
  friend struct ::TableStruct_supervisor_2fv1alpha1_2fssh_2eproto;
};
// -------------------------------------------------------------------

class FetchAuthorizedKeysResponse final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:supervisor.v1alpha1.FetchAuthorizedKeysResponse) */ {
 public:
  inline FetchAuthorizedKeysResponse() : FetchAuthorizedKeysResponse(nullptr) {}
  ~FetchAuthorizedKeysResponse() override;
  explicit PROTOBUF_CONSTEXPR FetchAuthorizedKeysResponse(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  FetchAuthorizedKeysResponse(const FetchAuthorizedKeysResponse& from);
  FetchAuthorizedKeysResponse(FetchAuthorizedKeysResponse&& from) noexcept
    : FetchAuthorizedKeysResponse() {
    *this = ::std::move(from);
  }

  inline FetchAuthorizedKeysResponse& operator=(const FetchAuthorizedKeysResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline FetchAuthorizedKeysResponse& operator=(FetchAuthorizedKeysResponse&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const FetchAuthorizedKeysResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const FetchAuthorizedKeysResponse* internal_default_instance() {
    return reinterpret_cast<const FetchAuthorizedKeysResponse*>(
               &_FetchAuthorizedKeysResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(FetchAuthorizedKeysResponse& a, FetchAuthorizedKeysResponse& b) {
    a.Swap(&b);
  }
  inline void Swap(FetchAuthorizedKeysResponse* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(FetchAuthorizedKeysResponse* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  FetchAuthorizedKeysResponse* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<FetchAuthorizedKeysResponse>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const FetchAuthorizedKeysResponse& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const FetchAuthorizedKeysResponse& from) {
    FetchAuthorizedKeysResponse::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(FetchAuthorizedKeysResponse* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "supervisor.v1alpha1.FetchAuthorizedKeysResponse";
  }
  protected:
  explicit FetchAuthorizedKeysResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kAuthorizedKeysFieldNumber = 1,
  };
  // string authorized_keys = 1 [json_name = "authorizedKeys"];
  void clear_authorized_keys();
  const std::string& authorized_keys() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_authorized_keys(ArgT0&& arg0, ArgT... args);
  std::string* mutable_authorized_keys();
  PROTOBUF_NODISCARD std::string* release_authorized_keys();
  void set_allocated_authorized_keys(std::string* authorized_keys);
  private:
  const std::string& _internal_authorized_keys() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_authorized_keys(const std::string& value);
  std::string* _internal_mutable_authorized_keys();
  public:

  // @@protoc_insertion_point(class_scope:supervisor.v1alpha1.FetchAuthorizedKeysResponse)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr authorized_keys_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_supervisor_2fv1alpha1_2fssh_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// FetchAuthorizedKeysRequest

// -------------------------------------------------------------------

// FetchAuthorizedKeysResponse

// string authorized_keys = 1 [json_name = "authorizedKeys"];
inline void FetchAuthorizedKeysResponse::clear_authorized_keys() {
  _impl_.authorized_keys_.ClearToEmpty();
}
inline const std::string& FetchAuthorizedKeysResponse::authorized_keys() const {
  // @@protoc_insertion_point(field_get:supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys)
  return _internal_authorized_keys();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void FetchAuthorizedKeysResponse::set_authorized_keys(ArgT0&& arg0, ArgT... args) {
 
 _impl_.authorized_keys_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys)
}
inline std::string* FetchAuthorizedKeysResponse::mutable_authorized_keys() {
  std::string* _s = _internal_mutable_authorized_keys();
  // @@protoc_insertion_point(field_mutable:supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys)
  return _s;
}
inline const std::string& FetchAuthorizedKeysResponse::_internal_authorized_keys() const {
  return _impl_.authorized_keys_.Get();
}
inline void FetchAuthorizedKeysResponse::_internal_set_authorized_keys(const std::string& value) {
  
  _impl_.authorized_keys_.Set(value, GetArenaForAllocation());
}
inline std::string* FetchAuthorizedKeysResponse::_internal_mutable_authorized_keys() {
  
  return _impl_.authorized_keys_.Mutable(GetArenaForAllocation());
}
inline std::string* FetchAuthorizedKeysResponse::release_authorized_keys() {
  // @@protoc_insertion_point(field_release:supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys)
  return _impl_.authorized_keys_.Release();
}
inline void FetchAuthorizedKeysResponse::set_allocated_authorized_keys(std::string* authorized_keys) {
  if (authorized_keys != nullptr) {
    
  } else {
    
  }
  _impl_.authorized_keys_.SetAllocated(authorized_keys, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.authorized_keys_.IsDefault()) {
    _impl_.authorized_keys_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:supervisor.v1alpha1.FetchAuthorizedKeysResponse.authorized_keys)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1alpha1
}  // namespace supervisor

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_supervisor_2fv1alpha1_2fssh_2eproto