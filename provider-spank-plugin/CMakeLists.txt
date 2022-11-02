cmake_minimum_required(VERSION 3.24.0)

project(
  spank-provider
  VERSION 1.0.4
  DESCRIPTION "SPANK plugin for providers"
  LANGUAGES C CXX)

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

include(FetchContent)
set(FETCHCONTENT_QUIET OFF)

FetchContent_Declare(
  slurm
  GIT_REPOSITORY https://github.com/SchedMD/slurm
  GIT_TAG slurm-22-05-3-1)
FetchContent_GetProperties(slurm)
if(NOT slurm_POPULATED)
  FetchContent_Populate(slurm)
  execute_process(COMMAND ${slurm_SOURCE_DIR}/configure
                  WORKING_DIRECTORY ${slurm_SOURCE_DIR})
endif()

FetchContent_Declare(
  grpc
  GIT_REPOSITORY https://github.com/grpc/grpc
  GIT_TAG v1.49.1)
FetchContent_GetProperties(grpc)
if(NOT grpc_POPULATED)
  FetchContent_Populate(grpc)
  add_subdirectory(${grpc_SOURCE_DIR} ${grpc_BINARY_DIR} EXCLUDE_FROM_ALL)
endif()

find_package(Threads REQUIRED)

add_compile_options(
  -O3
  -m64
  -Wall
  -Wextra
  -Wformat=2
  -Wpointer-arith
  -Winit-self
  -Wcast-align
  -Wcast-qual
  -Wunreachable-code
  -fno-common
  -Wno-unused-parameter)

file(GLOB_RECURSE sources src/*.cc src/*.c src/*.h)
file(GLOB_RECURSE grpc_sources gen/cpp/*.cc gen/cpp/*.h)

add_library(spank_provider SHARED ${sources} ${grpc_sources})

set_target_properties(spank_provider PROPERTIES PREFIX "")

target_include_directories(
  spank_provider
  PUBLIC ${slurm_SOURCE_DIR}
  PUBLIC "gen/cpp")

target_link_libraries(spank_provider Threads::Threads m grpc++)

install(
  TARGETS spank_provider
  LIBRARY DESTINATION ${CMAKE_INSTALL_PREFIX}
  ARCHIVE DESTINATION ${CMAKE_INSTALL_PREFIX})

include(package.cmake)
