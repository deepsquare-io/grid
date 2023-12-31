# Image Upscaler

DeepSquare integrates the neural network model optimised image upscaler [Real-ESRGAN](https://github.com/xinntao/Real-ESRGAN).

This guide presents how we designed and implmented the workflow to execute Real-ESRGAN on the DeepSquare Grid.

## Design

We designed the workflow as follows:

- The input can be an image or a video
  - If the input is a video, the frames are extracted.
  - The extracted images are distributed evenly among the GPUs.
  - The scaled frames are reassembled into the original video.
- The output will be the scaled image or video.

The docker image is already compiled from [github.com/xinntao/Real-ESRGAN](https://github.com/xinntao/Real-ESRGAN) and is exported as `registry-1.deepsquare.run/library/upscaling:latest`.

The steps are as follows:

1. Compute the number of images (extract and distribute the frames into batches if the input is a video).

2. Upscaling of the images.

3. If the input is a video, reassemble the images into a video.

```mermaid
---
title: Upscaling Workflow Architecture
---
flowchart LR
  video[Video] --> extract_frames

  subgraph Compute the number of images
    extract_frames[[Frames extraction]]
    split_images_into_batches[[Creation of batches of images]]
    extract_frames --> split_images_into_batches
  end

  split_images_into_batches --> upscaling1
  split_images_into_batches --> upscaling2
  split_images_into_batches --> upscaling3

  subgraph Upscale in Parallel
    upscaling1[[Upscaling Task 1]]
    upscaling2[[Upscaling Task 2]]
    upscaling3[[Upscaling Task 3]]
  end

  upscaling1 --> combine_the_frames
  upscaling2 --> combine_the_frames
  upscaling3 --> combine_the_frames

  subgraph Combine the frames
    combine_the_frames[[Frames Concatenation]]
  end

  combine_the_frames --> output

  output[New Upscaled Video]
```

## Implementation

Let's start with the input, output and resources.

### Resource allocation, environment variable and input/output

We will use 4 parallel tasks. Each task has 8 CPUs, 8 GB of RAM, 1 GPU.

```yaml
enableLogging: true

resources:
  tasks: 4
  cpusPerTask: 8
  memPerCpu: 8000
  gpus: 4

input:
  s3:
    region: region
    bucketUrl: s3://test
    path: '/test'
    accessKeyId: accessKeyId
    secretAccessKey: secretAccessKey
    endpointUrl: https://example

output:
  s3:
    region: region
    bucketUrl: s3://test
    path: '/test'
    accessKeyId: accessKeyId
    secretAccessKey: secretAccessKey
    endpointUrl: https://example
continuousOutputSync: true

env:
  - key: IS_VIDEO
    value: 'false' # Change this value if you want to render a video or an image
  - key: IS_FACE
    value: 'false'
  - key: IS_ANIME
    value: 'false'
```

### 1. Compute the number of frames

Let's implement the first step:

```yaml
steps:
  - name: compute-frames
    run:
      container:
        deepsquareHosted: true
        apptainer: true
        registry: registry-1.deepsquare.run
        image: library/upscaling:latest
      shell: /bin/bash
      command: |-
        set -e

        mkdir -p "${STORAGE_PATH}/input_frames/"
        mkdir -p "${STORAGE_PATH}/output_frames/"

        # Look for video and images
        videosFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | wc -l)"
        imagesFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | wc -l)"

        if "${IS_VIDEO}"; then
          if [[ ${videosFound} -ge "1" ]]; then
            mkdir -p "${STORAGE_PATH}/input_video/"

            # Only consider the first video on the list
            videoSourceFile="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | head -1)"

            # Extract all the frames in the input_frames directory
            ffmpeg -i "${videoSourceFile}" -qscale:v 1 -qmin 1 -qmax 1 -vsync 0 "${STORAGE_PATH}/input_frames/frame%08d.png"
          else
              echo "No input video found, exiting" && exit 1
          fi
        else
          if [[ ${imagesFound} -ge "1" ]]; then
            # Copy the frames in the input_frames directory
            find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | xargs -I{} cp "{}" "${STORAGE_PATH}/input_frames/" || (echo "Zero picture found" && exit 1)
          else
            echo "No input picture found, exiting" && exit 1
          fi
        fi

        # Compute the number of frames and frames per task
        totalFrames=$(find "${STORAGE_PATH}/input_frames/" -type f | wc -l)
        framesPerTask=$(( (totalFrames + NTASKS -1) /NTASKS))

        # Distribute the frames into batches
        for i in $(seq 1 "$NTASKS"); do
          cd "${STORAGE_PATH}/input_frames/"
          mkdir -p "${STORAGE_PATH}/input_frames/batch-${i}"
          if "${IS_VIDEO}"; then
            mv $(find . -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | head -${framesPerTask}) "${STORAGE_PATH}/input_frames/batch-${i}/"
          else
            find . -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | head -${framesPerTask} | while IFS= read -r file; do
              mv "$file" "${STORAGE_PATH}/input_frames/batch-${i}/"
            done
          fi

          cd -
        done
```

1. We use `file -i` which displays the MIME type, and filters out all files that are not videos or images.
2. We extract and store the images in the `input_frames` directory.
3. Finally, we calculate the number of images and distribute them to the tasks.

### 2. Upscale the frames

Let's implement the second step. We need to launch multiple substeps in parallel. We can use the `for` directive, and the variable `$index` to select the batch directory:

```yaml
steps:
  # ...
  - name: upscaling-loop
    for:
      parallel: true
      range:
        begin: 1
        end: 4 # NTASKS
      steps:
        - name: upscale
          run:
            resources:
              gpusPerTask: 1
            container:
              deepsquareHosted: true
              apptainer: true
              registry: registry-1.deepsquare.run
              image: library/upscaling:latest
            shell: /bin/bash
            command: |-
              set -e

              echo "Upscaling batch ${index}"
              /opt/Real-ESRGAN/upscale.sh  "${STORAGE_PATH}/input_frames/batch-${index}"
```

After executing the `upscale.sh` script, the frames will be generated inside the `output_frames` directory.

### 3. (Video) Re-assemble the frames

Using `ffmpeg`, we can determine the FPS of the original video and reassemble all the frames into a new video stream:

```yaml
steps:
  # ...
  - name: re-encode-video
    run:
      container:
        deepsquareHosted: true
        apptainer: true
        registry: registry-1.deepsquare.run
        image: library/upscaling:latest
      resources:
        tasks: 1
        cpusPerTask: 8 # TO TEMPLATE
        memPerCpu: 8000 # TO TEMPLATE
      command: |-
        set -e

        videoSourceFile="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | head -1)"

        if "${IS_VIDEO}"; then
          source_fps="$(ffmpeg -i "${videoSourceFile}" 2>&1 | sed -n "s/.*, \(.*\) fp.*/\1/p")"
          ffmpeg -r "${source_fps}" -i "${STORAGE_PATH}/output_frames/frame%08d_out.png" -i "${videoSourceFile}" -map 0:v:0 -map 1:a:0 -c:a copy -c:v libx264 -r "${source_fps}" -pix_fmt yuv420p "${DEEPSQUARE_OUTPUT}/result.mp4"
        fi
```

## Conclusion

We've developped this workflow similarly to [blender](blender). It always follows the same pattern:

1. Divide the data into tasks
2. Execute a `for` loop that runs one task per GPU
3. Reassemble the output

<details>

<summary>Complete Workflow</summary>

```yaml
enableLogging: true
resources:
  tasks: 4
  cpusPerTask: 8
  memPerCpu: 8000
  gpus: 4
input:
  s3:
    region: region
    bucketUrl: s3://test
    path: '/test'
    accessKeyId: accessKeyId
    secretAccessKey: secretAccessKey
    endpointUrl: https://example
output:
  s3:
    region: region
    bucketUrl: s3://test
    path: '/test'
    accessKeyId: accessKeyId
    secretAccessKey: secretAccessKey
    endpointUrl: https://example
continuousOutputSync: true
env:
  - key: IS_VIDEO
    value: 'false' # Change this value if you want to render a video or an image
  - key: IS_FACE
    value: 'false'
  - key: IS_ANIME
    value: 'false'
steps:
  - name: compute-frames
    run:
      container:
        deepsquareHosted: true
        apptainer: true
        registry: registry-1.deepsquare.run
        image: library/upscaling:latest
      shell: /bin/bash
      command: |-
        set -e

        mkdir -p "${STORAGE_PATH}/input_frames/"
        mkdir -p "${STORAGE_PATH}/output_frames/"

        # Look for video and images
        videosFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | wc -l)"
        imagesFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | wc -l)"

        if "${IS_VIDEO}"; then
          if [[ ${videosFound} -ge "1" ]]; then
            mkdir -p "${STORAGE_PATH}/input_video/"

            # Only consider the first video on the list
            videoSourceFile="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | head -1)"

            # Extract all the frames in the input_frames directory
            ffmpeg -i "${videoSourceFile}" -qscale:v 1 -qmin 1 -qmax 1 -vsync 0 "${STORAGE_PATH}/input_frames/frame%08d.png"
          else
              echo "No input video found, exiting" && exit 1
          fi
        else
          if [[ ${imagesFound} -ge "1" ]]; then
            # Copy the frames in the input_frames directory
            find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | xargs -I{} cp "{}" "${STORAGE_PATH}/input_frames/" || (echo "Zero picture found" && exit 1)
          else
            echo "No input picture found, exiting" && exit 1
          fi
        fi

        # Compute the number of frames and frames per task
        totalFrames=$(find "${STORAGE_PATH}/input_frames/" -type f | wc -l)
        framesPerTask=$(( (totalFrames + NTASKS -1) /NTASKS))

        # Distribute the frames into batches
        for i in $(seq 1 "$NTASKS"); do
          cd "${STORAGE_PATH}/input_frames/"
          mkdir -p "${STORAGE_PATH}/input_frames/batch-${i}"
          if "${IS_VIDEO}"; then
            mv $(find . -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | head -${framesPerTask}) "${STORAGE_PATH}/input_frames/batch-${i}/"
          else
            find . -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | head -${framesPerTask} | while IFS= read -r file; do
              mv "$file" "${STORAGE_PATH}/input_frames/batch-${i}/"
            done
          fi

          cd -
        done
  - name: upscaling-loop
    for:
      parallel: true
      range:
        begin: 1
        end: 4 # NTASKS
      steps:
        - name: upscale
          run:
            container:
              deepsquareHosted: true
              apptainer: true
              registry: registry-1.deepsquare.run
              image: library/upscaling:latest
            resources:
              gpusPerTask: 1
            shell: /bin/bash
            command: |-
              set -e

              echo "Upscaling batch ${index}"
              /opt/Real-ESRGAN/upscale.sh  "${STORAGE_PATH}/input_frames/batch-${index}"
  - name: re-encode-video
    run:
      container:
        deepsquareHosted: true
        apptainer: true
        registry: registry-1.deepsquare.run
        image: library/upscaling:latest
      resources:
        tasks: 1
        cpusPerTask: 8 # TO TEMPLATE
        memPerCpu: 8000 # TO TEMPLATE
      command: |-
        set -e

        videoSourceFile="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | head -1)"

        if "${IS_VIDEO}"; then
          source_fps="$(ffmpeg -i "${videoSourceFile}" 2>&1 | sed -n "s/.*, \(.*\) fp.*/\1/p")"
          ffmpeg -r "${source_fps}" -i "${STORAGE_PATH}/output_frames/frame%08d_out.png" -i "${videoSourceFile}" -map 0:v:0 -map 1:a:0 -c:a copy -c:v libx264 -r "${source_fps}" -pix_fmt yuv420p "${DEEPSQUARE_OUTPUT}/result.mp4"
        fi
```

</details>
