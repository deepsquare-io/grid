package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanStepRunResources = model.StepRunResources{
	Tasks:       utils.Ptr(1),
	CpusPerTask: utils.Ptr(1),
	MemPerCPU:   utils.Ptr(1),
	GpusPerTask: utils.Ptr(0),
}

func cleanStepRun(command string) *model.StepRun {
	return &model.StepRun{
		Resources: &cleanStepRunResources,
		Container: &model.ContainerRun{
			Image:    "image",
			Registry: utils.Ptr("registry"),
			Username: utils.Ptr("username"),
			Password: utils.Ptr("password"),
			X11:      utils.Ptr(true),
			Mounts: []*model.Mount{
				{
					HostDir:      "/host",
					ContainerDir: "/container",
					Options:      "ro",
				},
			},
		},
		Env: []*model.EnvVar{
			{
				Key:   "test",
				Value: "value",
			},
		},
		Command: command,
	}
}

func TestRenderStepRun(t *testing.T) {
	tests := []struct {
		input         model.StepRun
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanStepRun("hostname"),
			expected: `/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' > /dev/null
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.Apptainer = utils.Ptr(true)
				return r
			}(),
			expected: `export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sif"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/apptainer --silent pull --disable-cache "$IMAGE_PATH" 'docker://registry/image'
/usr/bin/echo "Image successfully imported!"
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
  "$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.MapRoot = utils.Ptr(true)
				return r
			}(),
			expected: `/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' > /dev/null
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir=/deepsquare \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with pyxis maproot",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.WorkDir = utils.Ptr("/home")
				return r
			}(),
			expected: `/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' > /dev/null
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
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --no-container-remap-root \
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
  --container-workdir='/home' \
  --container-image="$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with pyxis workdir",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.Apptainer = utils.Ptr(true)
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
  "$STORAGE_PATH"'/test/my.sqshfs' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer absolute path image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.DeepsquareHosted = utils.Ptr(true)
				return r
			}(),
			expected: `export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
  '/opt/software/registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with deepsquare-hosted image",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command: `hostname
/usr/bin/echo "test"`,
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname
/usr/bin/echo "test"'`,
			title: "Positive test with multiline command",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				WorkDir:   utils.Ptr("/dir"),
				Command:   `hostname`,
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'cd '"'"'/dir'"'"' || { echo "change dir to working directory failed"; exit 1; };''hostname'`,
			title: "Positive test with workdir",
		},
		{
			input: model.StepRun{
				Env:               cleanStepRun("").Env,
				Resources:         &cleanStepRunResources,
				Command:           "hostname",
				DisableCPUBinding: utils.Ptr(true),
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --cpu-bind=none \
  /bin/sh -c 'hostname'`,
			title: "Positive test with disable cpu-bind",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
				MapRoot:   utils.Ptr(true),
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /usr/bin/unshare --user --map-root-user --mount \
  /bin/sh -c 'hostname'`,
			title: "Positive test with map root",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
				Mpi:       utils.Ptr("none"),
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --mpi=none \
  /bin/sh -c 'hostname'`,
			title: "Positive test with mpi",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
				Network:   utils.Ptr("slirp4netns"),
				DNS:       []string{"1.1.1.1"},
				CustomNetworkInterfaces: []*model.NetworkInterface{
					&cleanWireguardNIC,
				},
			},
			expected: `test='value' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'set -e

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
/usr/bin/unshare --user --net --mount --map-root-user /bin/sh -c '"'"'
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

/usr/sbin/ip link add dev net0 type wireguard
/usr/sbin/ip -4 address add 10.0.0.1/32 dev net0
/usr/bin/echo '"'"'"'"'"'"'"'"'abc'"'"'"'"'"'"'"'"' > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 private-key "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' allowed-ips 0.0.0.0/0,172.10.0.0/32
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' endpoint 10.0.0.0:30
/usr/bin/echo '"'"'"'"'"'"'"'"'sha'"'"'"'"'"'"'"'"' > "$(pwd)/wg-preshared-0"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' preshared-key "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' persistent-keepalive 20
/usr/sbin/ip link set mtu 1420 up dev net0
/usr/sbin/ip -4 route add 0.0.0.0/0 dev net0
/usr/sbin/ip -4 route add 172.10.0.0/32 dev net0

/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

'"'"''"'"'hostname'"'"' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
}
trap cleanup EXIT INT TERM

wait $child
'`,
			title: "Positive test with wireguard tunnel",
		},
		{
			input: model.StepRun{
				Container: cleanApptainerStepRun("").Container,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
				Network:   utils.Ptr("slirp4netns"),
				DNS:       []string{"1.1.1.1"},
				CustomNetworkInterfaces: []*model.NetworkInterface{
					&cleanWireguardNIC,
				},
			},
			expected: `export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sif"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/apptainer --silent pull --disable-cache "$IMAGE_PATH" 'docker://registry/image'
/usr/bin/echo "Image successfully imported!"
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /bin/sh -c 'set -e

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
/usr/bin/unshare --user --net --mount --map-root-user /bin/sh -c '"'"'
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

/usr/sbin/ip link add dev net0 type wireguard
/usr/sbin/ip -4 address add 10.0.0.1/32 dev net0
/usr/bin/echo '"'"'"'"'"'"'"'"'abc'"'"'"'"'"'"'"'"' > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 private-key "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' allowed-ips 0.0.0.0/0,172.10.0.0/32
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' endpoint 10.0.0.0:30
/usr/bin/echo '"'"'"'"'"'"'"'"'sha'"'"'"'"'"'"'"'"' > "$(pwd)/wg-preshared-0"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' preshared-key "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' persistent-keepalive 20
/usr/sbin/ip link set mtu 1420 up dev net0
/usr/sbin/ip -4 route add 0.0.0.0/0 dev net0
/usr/sbin/ip -4 route add 172.10.0.0/32 dev net0

/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

'"'"''"'"'/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
  "$IMAGE_PATH" \
  /bin/sh -c '"'"'"'"'"'"'"'"'hostname'"'"'"'"'"'"'"'"''"'"' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
}
trap cleanup EXIT INT TERM

wait $child
'`,
			title: "Positive test with wireguard tunnel and apptainer",
		},
		{
			input: model.StepRun{
				Container: cleanStepRun("").Container,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
				Network:   utils.Ptr("slirp4netns"),
				DNS:       []string{"1.1.1.1"},
				CustomNetworkInterfaces: []*model.NetworkInterface{
					&cleanWireguardNIC,
				},
			},
			expected: `/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/enroot import -o "$IMAGE_PATH" -- 'docker://registry#image' > /dev/null
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
/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
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
/usr/bin/unshare --user --net --mount --map-root-user /bin/sh -c '"'"'
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

/usr/sbin/ip link add dev net0 type wireguard
/usr/sbin/ip -4 address add 10.0.0.1/32 dev net0
/usr/bin/echo '"'"'"'"'"'"'"'"'abc'"'"'"'"'"'"'"'"' > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 private-key "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' allowed-ips 0.0.0.0/0,172.10.0.0/32
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' endpoint 10.0.0.0:30
/usr/bin/echo '"'"'"'"'"'"'"'"'sha'"'"'"'"'"'"'"'"' > "$(pwd)/wg-preshared-0"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' preshared-key "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'"'"'"'"'"'"'pub'"'"'"'"'"'"'"'"' persistent-keepalive 20
/usr/sbin/ip link set mtu 1420 up dev net0
/usr/sbin/ip -4 route add 0.0.0.0/0 dev net0
/usr/sbin/ip -4 route add 172.10.0.0/32 dev net0

/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

'"'"''"'"'/usr/bin/cat << '"'"'"'"'"'"'"'"'EOFenroot'"'"'"'"'"'"'"'"' > "$STORAGE_PATH/enroot.conf"
#ENROOT_REMAP_ROOT=n
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/env"
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
  /usr/bin/echo '"'"'"'"'"'"'"'"'/host /container bind,ro'"'"'"'"'"'"'"'"'
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
  /bin/sh -c '"'"'"'"'"'"'"'"'hostname'"'"'"'"'"'"'"'"''"'"' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
}
trap cleanup EXIT INT TERM

wait $child
'`,
			title: "Positive test with wireguard tunnel and enroot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStepRun(&cleanJob, cleanStepWithRun(&tt.input))

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
				require.NoError(t, renderer.Shellcheck(actual))
			}
		})
	}
}
