#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='2'
export MEM_PER_CPU='4000'
export GPUS_PER_TASK='1'
export GPUS='1'
export CPUS='2'
export MEM='8000'
STORAGE_PATH="/opt/cache/shared/$(id -u)/$SLURM_JOB_NAME.$SLURM_JOB_ID"
export STORAGE_PATH
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
DEEPSQUARE_TMP="/opt/cache/persistent/user-$(id -u)"
export DEEPSQUARE_TMP
DEEPSQUARE_SHARED_TMP="/opt/cache/persistent/user-$(id -u)"
export DEEPSQUARE_SHARED_TMP
DEEPSQUARE_SHARED_WORLD_TMP="/opt/cache/world-tmp"
export DEEPSQUARE_SHARED_WORLD_TMP
DEEPSQUARE_DISK_TMP="/opt/cache/disk/tmp/user-$(id -u)"
export DEEPSQUARE_DISK_TMP
DEEPSQUARE_DISK_WORLD_TMP="/opt/cache/disk/world-tmp"
export DEEPSQUARE_DISK_WORLD_TMP
ENROOT_RUNTIME_PATH="/run/enroot/user-$(id -u)"
export ENROOT_RUNTIME_PATH
ENROOT_CACHE_PATH="/opt/cache/enroot/group-$(id -g)"
export ENROOT_CACHE_PATH
ENROOT_DATA_PATH="/mnt/scratch/tmp/enroot/containers/user-$(id -u)"
export ENROOT_DATA_PATH
export APPTAINER_TMPDIR="/mnt/scratch/tmp/apptainer"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT" "$DEEPSQUARE_TMP"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chmod 700 "$DEEPSQUARE_TMP"

# Create mountpoints
/usr/bin/mkdir -p "$STORAGE_PATH/tmp"
/usr/bin/mkdir -p "$STORAGE_PATH/world-tmp"
/usr/bin/mkdir -p "$STORAGE_PATH/disk/tmp"
/usr/bin/mkdir -p "$STORAGE_PATH/disk/world-tmp"

/usr/bin/chown -R "$(id -u):$(id -g)" "$STORAGE_PATH"

for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  srun --job-name="prepare-dir" -N 1-1 -n 1 -w "$node" sh -c 'mkdir -p "$DEEPSQUARE_DISK_TMP" && /usr/bin/chmod 700 "$DEEPSQUARE_DISK_TMP"'
done

cleanup() {
  /bin/rm -rf "$STORAGE_PATH"
}
trap cleanup EXIT INT TERM

cd "$STORAGE_PATH/"
loadDeepsquareEnv() {
  while IFS= read -r envvar; do
    printf ',%s' "$envvar"
  done < <(/usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=')
}

declare -A EXIT_SIGNALS
/usr/bin/echo 'Running: ''kasmvnc'

IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry-1.docker.io#kasmweb/desktop:1.14.0-rolling' &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
if [ $? -ne 0 ]; then
  cat "/tmp/enroot.import.$SLURM_JOB_ID.log"
fi
set -e
tries=1; while [ "$tries" -lt 10 ]; do
  if /usr/bin/file "$IMAGE_PATH" | /usr/bin/grep -q "Squashfs filesystem"; then
    break
  fi
  /usr/bin/echo "Image is not complete. Wait a few seconds... ($tries/10)"
  /usr/bin/sleep 10
  tries=$((tries+1))
done
if [ "$tries" -ge 10 ]; then
  /usr/bin/echo "Image import failure (corrupted image). Please try again."
  exit 1
fi
/usr/bin/echo "Image successfully imported!"
# shellcheck disable=SC2097,SC2098,SC1078
/usr/bin/srun --job-name='kasmvnc' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=2 \
  --mem-per-cpu=4000M \
  --gpus-per-task=1 \
  --ntasks=1 \
  --gpu-bind=none \
  /bin/sh -c '/usr/bin/enroot create --name "container-$SLURM_JOB_ID" -- "$IMAGE_PATH"
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID"
}
trap enrootClean EXIT INT TERM
''set -e

nsenter_flags() {
  pid="$1"
  flags="--target=${pid}"
  userns="$(readlink "/proc/${pid}/ns/user")"
  mntns="$(readlink "/proc/${pid}/ns/mnt")"
  netns="$(readlink "/proc/${pid}/ns/net")"

  self_userns="$(readlink /proc/self/ns/user)"
  self_mntns="$(readlink /proc/self/ns/mnt)"
  self_netns="$(readlink /proc/self/ns/net)"

  if [ "${userns}" != "${self_userns}" ]; then
    flags="$flags --preserve-credentials -U"
  fi
  if [ "${mntns}" != "${self_mntns}" ]; then
    flags="$flags -m"
  fi
  if [ "${netns}" != "${self_netns}" ]; then
    flags="$flags -n"
  fi
  echo "${flags}"
}

wait_for_network_namespace() {
  # Wait that the namespace is ready.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    flags=$(nsenter_flags "$1")
    if echo "$flags" | grep -qvw -- -n; then
      flags="$flags -n"
    fi
    # shellcheck disable=SC2086
    if nsenter ${flags} true >/dev/null 2>&1; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

# shellcheck disable=SC2016,SC1078,SC1079
/usr/bin/unshare --user --net --mount /bin/sh -c '"'"'
set -e

nsenter_flags() {
  pid=$1
  flags="--target=${pid}"
  userns="$(readlink "/proc/${pid}/ns/user")"
  mntns="$(readlink "/proc/${pid}/ns/mnt")"
  netns="$(readlink "/proc/${pid}/ns/net")"

  self_userns="$(readlink /proc/self/ns/user)"
  self_mntns="$(readlink /proc/self/ns/mnt)"
  self_netns="$(readlink /proc/self/ns/net)"

  if [ "${userns}" != "${self_userns}" ]; then
    flags="$flags --preserve-credentials -U"
  fi
  if [ "${mntns}" != "${self_mntns}" ]; then
    flags="$flags -m"
  fi
  if [ "${netns}" != "${self_netns}" ]; then
    flags="$flags -n"
  fi
  /usr/bin/echo "${flags}"
}

wait_for_network_device() {
  # Wait that the device appears.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    if nsenter $(nsenter_flags "$1") ip addr show "$2"; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

wait_for_network_device $$ tap0

/usr/bin/bore -s bore.deepsquare.run -p 2200 -ls localhost -lp 6901 -r &

'"'"''"'"'/usr/bin/cat <<'"'"'"'"'"'"'"'"'EOFenroot'"'"'"'"'"'"'"'"' >"$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
  /usr/bin/echo "VNC_PW='"'"'"'"'"'"'"'"'password'"'"'"'"'"'"'"'"'"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
}

hooks() {
  /usr/bin/cat << '"'"'"'"'"'"'"'"'EOFrclocal'"'"'"'"'"'"'"'"' > "${ENROOT_ROOTFS}/etc/rc.local"
cd "/deepsquare" || { echo "change dir to working directory failed"; exit 1; }
exec "$@"
EOFrclocal
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
  /bin/sh -c '"'"'"'"'"'"'"'"'mkdir -p $HOME/Desktop
cd $HOME
/dockerstartup/kasm_default_profile.sh
cp /dockerstartup/vnc_startup.sh /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_audio_out_websocket$/#start_audio_out_websocket/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_audio_out$/#start_audio_out/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_audio_in$/#start_audio_in/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_upload$/#start_upload/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_webcam$/#start_webcam/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_printer$/#start_printer/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/start_gamepad$/#start_gamepad/g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/ --ssl//g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s| --cert \$\{HOME\}/\.vnc/self\.pem --certkey \$\{HOME\}/\.vnc/self\.pem||g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s/ -sslOnly//g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' /tmp/vnc_startup.sh
mkdir -p "$HOME/.vnc"
cp /usr/share/kasmvnc/kasmvnc_defaults.yaml "$HOME/.vnc/kasmvnc.yaml"
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s|require_ssl: true|require_ssl: false|g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' "$HOME/.vnc/kasmvnc.yaml"
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s|pem_certificate: .*|pem_certificate: \$\{HOME\}/.vnc/self.pem|g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' "$HOME/.vnc/kasmvnc.yaml"
sed -Ei '"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'s|pem_key: .*|pem_key: \$\{HOME\}/.vnc/self.pem|g'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"'"' "$HOME/.vnc/kasmvnc.yaml"
/tmp/vnc_startup.sh --wait
'"'"'"'"'"'"'"'"''"'"' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
}
trap cleanup EXIT INT TERM

wait $child
'

for pid in "${!EXIT_SIGNALS[@]}"; do
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done

