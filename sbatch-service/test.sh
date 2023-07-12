#!/bin/bash -l

set -e

export NTASKS='1'
export CPUS_PER_TASK='1'
export MEM_PER_CPU='512'
export GPUS_PER_TASK='0'
export GPUS='0'
export CPUS='1'
export MEM='512'
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
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT" "$DEEPSQUARE_TMP" "$DEEPSQUARE_DISK_TMP"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chmod 700 "$DEEPSQUARE_TMP"
/usr/bin/chmod 700 "$DEEPSQUARE_DISK_TMP"
/usr/bin/chown -R "$(id -u):$(id -g)" "$STORAGE_PATH"

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
/usr/bin/echo 'Running: ''cache'
(
export CACHED_PATH='$DEEPSQUARE_INPUT'
export KEY='cache'
export KEY='my-cache'

( # CATCH FINALLY
finally() {
set -e
/usr/bin/echo 'Running: ''Push to cache'
/usr/bin/srun --job-name='Push to cache' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=512M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'CACHED_PATH="$(eval "echo \"$CACHED_PATH\"")"
KEY_PATH="$(eval "echo \"$KEY\"")"
if [ "$KEY_PATH" = "" ]; then
  echo "Key is empty"
  exit 1
fi
if [ "$DISK" = "true" ]; then
  CACHE_TMP="$DEEPSQUARE_DISK_TMP/$KEY_PATH/"
else
  CACHE_TMP="$DEEPSQUARE_SHARED_TMP/$KEY_PATH/"
fi
mkdir -p -m700 "$CACHE_TMP"
rsync -avP "$CACHED_PATH" "$CACHE_TMP"
'
}
trap finally EXIT INT TERM
/usr/bin/echo 'Running: ''Cache'
/usr/bin/echo 'Running: ''Pull from cache'
/usr/bin/srun --job-name='Pull from cache' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=512M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'CACHED_PATH="$(eval "echo \"$CACHED_PATH\"")"
KEY_PATH="$(eval "echo \"$KEY\"")"
if [ "$KEY_PATH" = "" ]; then
  echo "Key is empty"
  exit 1
fi
if [ "$DISK" = "true" ]; then
  CACHE_TMP="$DEEPSQUARE_DISK_TMP/$KEY_PATH/"
else
  CACHE_TMP="$DEEPSQUARE_SHARED_TMP/$KEY_PATH/"
fi

if [ -d "$CACHE_TMP" ]; then
  rsync -avP "$CACHE_TMP" "$CACHED_PATH"
else
  echo "There is no cache."
fi
'
/usr/bin/echo 'Running: ''test'
/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=512M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'echo test >> $DEEPSQUARE_INPUT/test'
) # CATCH FINALLY
)


for pid in "${!EXIT_SIGNALS[@]}"; do
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done

