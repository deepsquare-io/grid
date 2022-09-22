# A SLURM Job Completion plugin

The plugin send the actual job duration and allocated resources.

## Regenerate the gRPC Buf

You have to install protoc. Use the scripts inside the `protos` directory.

Then, configure the `CMakeLists.txt` and choose a gRPC version compatible with the version of protoc you have installed.

Then, compile gRPC using CMake, Make/Ninja and your favorite compiler.

```shell
mkdir -p build
cd build
cmake .. -G <your generator> # Ninja, or by default, Make
<your generator> grpc++
# For example: ninja -j 8 grpc++
# The grpc plugin will be installed at _deps/grpc-build
```

After compiling the plugin you can add them to the PATH environment variable or move it to `$HOME/.local/bin`.

My recommendation is to set the PATH environment variable:

```shell
# still in the build directory
export PATH="$(pwd)/_deps/grpc-build:$PATH"
```

After that, you can call `buf generate` in the `protos` directory.
