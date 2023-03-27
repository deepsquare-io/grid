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

My recommendation is to install the plugin.

```shell
cp ./build/_deps/grpc-build/grpc_*_plugin ~/.local/bin
```

After that, you can call `buf generate` in the `protos` directory.

## Getting the include files

The project uses FetchContent. In the `.vscode` directory, the include paths should already be configured, it only need to fetch the files:

```shell
mkdir -p build
cd build
cmake .. -G <your generator> # Ninja, or by default, Make
```

Should fetch and generate all the headers file.

## Building

```shell
mkdir -p build
cd build
cmake .. -G <your generator> # Ninja, or by default, Make
<your generator>
# The files are stored in the build directory
```

## Install

```shell
mkdir -p build
cd build
cmake .. -G <your generator> # Ninja, or by default, Make
<your generator> install
```

The .so file is installed in the `/usr/lib64/slurm` directory. To load the plugin, add these parameters in the slurm.conf:

```conf
JobCompType=jobcomp/provider
```

## Packaging (RPM/DEB)

```shell
mkdir -p build
cd build
cmake .. -G <your generator> # Ninja, or by default, Make
<your generator> package
```
