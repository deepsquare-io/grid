cmake_minimum_required(VERSION 3.24.0)

project(
  jobcomp-provider
  VERSION 1.0.0
  DESCRIPTION "Job completion plugin for providers"
  LANGUAGES CXX)

set(CMAKE_C_STANDARD 17)
set(CMAKE_CXX_STANDARD 17)

# Availables options
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
if(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)
  set(CMAKE_INSTALL_PREFIX
      /usr/lib64/slurm
      CACHE PATH "Install path" FORCE)
endif(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)
set(SLURM_SRC_DIR ${CMAKE_SOURCE_DIR}/externals/slurm)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_STATIC_LIBRARY_PREFIX "")
set(THREADS_PREFER_PTHREAD_FLAG ON)

option(USE_SYSTEM_GRPC "Use system installed gRPC" OFF)
if(USE_SYSTEM_GRPC)
  find_package(gRPC CONFIG REQUIRED)
else()
  include(FetchContent)
  FetchContent_Declare(
    gRPC
    GIT_REPOSITORY https://github.com/grpc/grpc
    GIT_TAG v1.49.1)
  set(FETCHCONTENT_QUIET OFF)
  FetchContent_MakeAvailable(gRPC)
endif()

find_package(Threads REQUIRED)

add_compile_options(
  -O3
  -m64
  -Wall
  -Wextra
  -Wshadow
  -Wformat=2
  -Wpointer-arith
  -Winit-self
  -Wcast-align
  -Wcast-qual
  -Wunreachable-code
  -fno-common
  -Wno-unused-parameter)

# Include generated *.pb.h files
include_directories("gen/cpp")

file(GLOB_RECURSE sources src/*.cc src/*.h)
file(GLOB_RECURSE grpc_sources gen/cpp/*.cc gen/cpp/*.h)

add_custom_command(
  OUTPUT ${SLURM_SRC_DIR}/slurm/slurm_version.h
  COMMAND ${SLURM_SRC_DIR}/configure
  WORKING_DIRECTORY ${SLURM_SRC_DIR}
  COMMENT "Configure slurm"
  VERBATIM)

add_custom_target(slurm_configure
                  DEPENDS ${SLURM_SRC_DIR}/slurm/slurm_version.h)

add_library(jobcomp_provider SHARED ${sources} ${grpc_sources})
add_library(jobcomp_provider_static STATIC ${sources} ${grpc_sources})

add_dependencies(jobcomp_provider slurm_configure)
add_dependencies(jobcomp_provider_static slurm_configure)

set_target_properties(jobcomp_provider PROPERTIES PREFIX "")
set_target_properties(jobcomp_provider_static
                      PROPERTIES PREFIX "" OUTPUT_NAME jobcomp_provider)

target_include_directories(jobcomp_provider PUBLIC ${SLURM_SRC_DIR})
target_include_directories(jobcomp_provider_static PUBLIC ${SLURM_SRC_DIR})

target_link_libraries(jobcomp_provider PRIVATE Threads::Threads m grpc++)
target_link_libraries(jobcomp_provider_static PRIVATE Threads::Threads m grpc++)

install(
  TARGETS jobcomp_provider jobcomp_provider_static
  LIBRARY DESTINATION ${CMAKE_INSTALL_PREFIX}
  ARCHIVE DESTINATION ${CMAKE_INSTALL_PREFIX})

include(package.cmake)
