project(
  spank-provider
  VERSION 1.0.5
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

include(cmake/dependencies.cmake)

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
