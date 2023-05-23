# Unity Render Streaming

The DeepSquare Grid can run workloads requiring x11. One usecase is to run [Unity Render Streaming](https://docs.unity3d.com/Packages/com.unity.renderstreaming@3.1/manual/index.html) which is a solution for peer-to-peer streaming.

Unity uses WebRTC which permits [**NAT traversal via STUN**](https://www.rfc-editor.org/rfc/rfc7635.html), so no need to use a tunnel to expose the video stream.

## Design

The workflow is not the most complicated part, so let's talk about how we can run a graphical session inside a container.

The conditions to run a graphical session is the following:

- Mount a x11 socket
- Select a x11 DISPLAY via the `DISPLAY` environment variable
- The container MUST have all the shared libraries to run a x11 session.

The shared libraries are:

```
libglvnd0
libgl1
libglx0
libegl1
libgles2
libxcb1-dev
libx11-xcb-dev
libglib2.0-0
libc++1-10
mesa-utils
libx11-dev
libxkbcommon-dev
libwayland-dev
libxrandr-dev
libegl1-mesa-dev
vulkan-utils
libgl1-mesa-glx
libvulkan1
libvulkan-dev
ocl-icd-opencl-dev
```

You should also install the Vulkan SDK if you are using Vulkan inside your Unity application.

The DeepSquare Grid should already mount the devices and shared libraries of the NVIDIA drivers.

Since this is not an HPC workload, parallel computing is not possible. Therefore, we can only run the service with a single GPU, as there is no parallel decomposition.

The process of a HPC application is typically:

1. Decomposition: Decompose the problem into smaller sub-problems.
2. Mapping: Map each sub-problem to a specific CPU.
3. Communication: Establish communication channels between CPUs.
4. Execution: Execute the workload on the assigned CPUs.
5. Aggregation: Aggregate the results obtained from each CPU to obtain the final solution.

Without decomposition, there is no parallel processing, and therefore no horizontal scaling of the application.

Basically, the application runs in one step: "Run the application".

## Implementation

For the sake of readability, we use YAML. You can use an online YAML to JSON converter if you wish to execute the workflow on the [dev environment](https://app.deepsquare.run/sandbox).

The implementation is quite immediate:

```yaml
job:
  enableLogging: true
  env:
    - key: DISPLAY
      value: ':99'
    - key: XDG_RUNTIME_DIR
      value: '/tmp'
  resources:
    tasks: 1
    cpusPerTask: 4
    memPerCpu: 4096
    gpusPerTask: 1
  steps:
    - name: tdp
      run:
        container:
          x11: true
          deepsquareHosted: true
          apptainer: true
          registry: registry-1.deepsquare.run
          image: 'library/tdp:latest'
        resources:
          tasks: 1
          cpusPerTask: 4
          memPerCpu: 4096
          gpusPerTask: 1
        command: |-
          /unityapp/DeepSquareURSSample.x86_64 \
            -maxplayers 8 \
            -webserverurl 'wss://tdp.deepsquare.run' \
            -renderfps 60 \
            -streamfps 30 \
            -sessionpin 123 \
            -minbitrate 10000 \
            -maxbitrate 50000 \
            -adaptativeupscaling \
            -minimumrenderscale 0.5 \
            -renderscalestepcount 0.1 \
            -displayfpscounter \
            --logFile -
```

However, most of the settings depends on how you have configured your Unity application and web server.

## Conclusion

While this isn't an HPC application, this sample can guide you on how to launch graphical application. To access to the graphical session, you may:

- Use WebRTC like Unity Render Streaming.
- Use [wireguard](/docs/deploy-deepsquare/guides/connecting-wireguard) to expose a VNC server which requires the port 5900 TCP and UDP.
- Use [bore](/docs/deploy-deepsquare/guides/connecting-bore) and a noVNC server, which is a web server which act as a VNC client.
