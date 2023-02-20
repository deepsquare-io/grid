# Workflow files
To run your workload on deepsquare, you have to provide a yaml workflow file.
It is a convenient way to define your workload in a human readable format. This workflow file will later be sent to a service which will translate it into a Slurm compatible job definition.  
A workflow file can contain multiple steps, some of which can be configured to run in parallel, on a different set of nodes/resources and each of them can run in different container images if needed.

Here is an example of the workflow file used to create the upscaling app currently available on DeepSquare:

```yaml
job:
  enableLogging: true
  resources:
    tasks: 4 # TO TEMPLATE: NTASKS, 1 NTASKS if not $IS_VIDEO, else 4 NTASKS
    cpusPerTask: 8 # TO TEMPLATE
    memPerCpu: 8000 # TO TEMPLATE
    gpusPerTask: 1 # TO TEMPLATE
  env:
    - key: IS_VIDEO
      value: 'false' # TO TEMPLATE
    - key: IS_FACE
      value: 'false' # TO TEMPLATE
    - key: IS_ANIME
      value: 'false' # TO TEMPLATE
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
  steps:
    - name: compute-frames
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
          gpusPerTask: 0 # TO TEMPLATE
        shell: /bin/bash
        command: |-
          set -ex

          rm -rf "${STORAGE_PATH}/input_frames/"
          rm -rf "${STORAGE_PATH}/input_video/"
          mkdir -p "${STORAGE_PATH}/input_frames/"
          mkdir -p "${STORAGE_PATH}/output_frames/"
          videosFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | wc -l)"
          imagesFound="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | wc -l)"

          if "${IS_VIDEO}"; then
            if [[ ${videosFound} -ge "1" ]]; then
                mkdir -p "${STORAGE_PATH}/input_video/"
                #Only consider the first video on the list
                videoSourceFile="$(find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: video/[^:]*$!!p' | head -1)"
                source_fps="$(ffmpeg -i "${videoSourceFile}" 2>&1 | sed -n "s/.*, \(.*\) fp.*/\1/p")"
                #Only take the first 5 minutes
                ffmpeg -ss 00:00 -accurate_seek -t 05:00 -i "${videoSourceFile}" -c:v libx264 -crf "${source_fps}" -c:a aac "${STORAGE_PATH}/input_video/input_video_trimmed.mp4"
                #Extract all the frames
                ffmpeg -i "${STORAGE_PATH}/input_video/input_video_trimmed.mp4" -qscale:v 1 -qmin 1 -qmax 1 -vsync 0 "${STORAGE_PATH}/input_frames/frame%08d.png"
            else
                echo "No input video found, exiting" && exit 1
            fi
          else
            if [[ ${imagesFound} -ge "1" ]]; then
                find "${DEEPSQUARE_INPUT}" -maxdepth 1 -type f -exec file -N -i -- {} + | sed -n 's!: image/[^:]*$!!p' | xargs -I{} cp "{}" "${STORAGE_PATH}/input_frames/" || (echo "Zero picture found" && exit 1)
            else
                echo "No input picture found, exiting" && exit 1
            fi
          fi
          totalFrames=$(find "${STORAGE_PATH}/input_frames/" -type f | wc -l)
          framesPerTask=$(( (totalFrames + NTASKS -1) /NTASKS))
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
          end: 4 # TO TEMPLATE: NTASKS
        steps:
          - name: upscale
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
                gpusPerTask: 1 # TO TEMPLATE
              shell: /bin/bash
              command: |-
                set -ex

                echo "Upscaling batch ${index}"
                /bin/bash /opt/Real-ESRGAN/upscale.sh "${STORAGE_PATH}/input_frames/batch-${index}"
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
          gpusPerTask: 0 # TO TEMPLATE
        command: |-
          set -ex

          if "${IS_VIDEO}"; then
            source_fps="$(ffmpeg -i "${STORAGE_PATH}/input_video/input_video_trimmed.mp4" 2>&1 | sed -n "s/.*, \(.*\) fp.*/\1/p")"
            ffmpeg -r "${source_fps}" -i "${STORAGE_PATH}/output_frames/frame%08d_out.png" -i "${STORAGE_PATH}/input_video/input_video_trimmed.mp4" -map 0:v:0 -map 1:a:0 -c:a copy -c:v libx264 -r "${source_fps}" -pix_fmt yuv420p "${DEEPSQUARE_OUTPUT}/result.mp4"
          fi
```
