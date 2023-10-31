include(FetchContent)
set(FETCHCONTENT_QUIET OFF)

FetchContent_Declare(
  slurm
  GIT_REPOSITORY https://github.com/SchedMD/slurm
  GIT_TAG slurm-23.02.6-1)
FetchContent_GetProperties(slurm)
if(NOT slurm_POPULATED)
  FetchContent_Populate(slurm)
  execute_process(COMMAND ${slurm_SOURCE_DIR}/configure
                  WORKING_DIRECTORY ${slurm_SOURCE_DIR})
endif()

set(ABSL_ENABLE_INSTALL ON)
FetchContent_Declare(
  grpc
  GIT_REPOSITORY https://github.com/grpc/grpc
  GIT_TAG v1.59.2)
FetchContent_GetProperties(grpc)
if(NOT grpc_POPULATED)
  FetchContent_Populate(grpc)
  add_subdirectory(${grpc_SOURCE_DIR} ${grpc_BINARY_DIR} EXCLUDE_FROM_ALL)
endif()

find_package(Threads REQUIRED)
