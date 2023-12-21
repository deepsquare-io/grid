# Leveraging environment variables to configure a workload

Environment variables are the best way to customize your environment without having to rebuild your container image.

## Priority

Environment variables can be set at the task level and at the step level.

From lowest to highest, the order of priority is:

- Job level
- DeepSquare defined variables (`DEEPSQUARE_INPUT`, `DEEPSQUARE_OUTPUT`, `DEEPSQUARE_TMP` `STORAGE_PATH`, `DEEPSQUARE_ENV`, ...)
- Container environment variables (`<container rootfs>/etc/environment`)
- Step level

## Isolation

A step-level environment variable will not be transferred to the next step.

A job-level environment variable that has been overridden will not be transferred to the next step.

Variables defined by DeepSquare that has been overridden will not be transferred to the next step.

The only way to pass information between steps is to use shared storage and DeepSquare variables, like `STORAGE_PATH`.

## DeepSquare defined variables

The environement variable `$STORAGE_PATH` stores the path of the shared storage path. That shared storage only lives during the job, it is a scratch storage.

There are many caches that can be used during a job:

- `$DEEPSQUARE_TMP` or `DEEPSQUARE_SHARED_TMP` stores the path to a shared file system cache isolated by user. The cache is periodically cleared according to the infrastructure provider's policy.
- `$DEEPSQUARE_SHARED_WORLD_TMP` stores the path to a shared file system cache **not** isolated. The cache is periodically cleared according to the infrastructure provider's policy.
- `$DEPSQUARE_DISK_TMP` stores the path to a per-site, per-node, per-user cache. The cache is cleared when the node is restarted, or periodically according to the infrastructure provider's policy.
- `$DEEPSQUARE_DISK_WORLD_TMP` stores the path to a cache per site, per node and for all users. The cache is cleared when the node is restarted, or periodically according to the infrastructure provider's policy.

**When running a container**, `$STORAGE_PATH` will be set to `/deepsquare`.

The other environment variables defined by DeepSquare are subdirectories in the `$STORAGE_PATH`:

| Environment variables                   | Value                      |
| --------------------------------------- | -------------------------- |
| STORAGE_PATH                            | /deepsquare                |
| DEEPSQUARE_INPUT                        | /deepsquare/input          |
| DEEPSQUARE_OUTPUT                       | /deepsquare/output         |
| DEEPSQUARE_ENV                          | /deepsquare/env            |
| DEEPSQUARE_TMP or DEEPSQUARE_SHARED_TMP | /deepsquare/tmp            |
| DEEPSQUARE_SHARED_WORLD_TMP             | /deepsquare/world-tmp      |
| DEEPSQUARE_DISK_TMP                     | /deepsquare/disk/tmp       |
| DEEPSQUARE_DISK_WORLD_TMP               | /deepsquare/disk/world-tmp |

**When running host**, all the variables are set to:

| Environment variables                   | Value                                 |
| --------------------------------------- | ------------------------------------- |
| STORAGE_PATH                            | /opt/cache/shared/&lt;user&gt;        |
| DEEPSQUARE_INPUT                        | /opt/cache/shared/&lt;user&gt;/input  |
| DEEPSQUARE_OUTPUT                       | /opt/cache/shared/&lt;user&gt;/output |
| DEEPSQUARE_ENV                          | /opt/cache/shared/&lt;user&gt;/env    |
| DEEPSQUARE_TMP or DEEPSQUARE_SHARED_TMP | /opt/cache/persistent/&lt;user&gt;    |
| DEEPSQUARE_SHARED_WORLD_TMP             | /opt/cache/world-tmp                  |
| DEEPSQUARE_DISK_TMP                     | /opt/cache/disk/tmp/&lt;user&gt;      |
| DEEPSQUARE_DISK_WORLD_TMP               | /opt/cache/disk/world-tmp             |

These values **will** change. **Always use environment variables and not hard-code the values to be future-proof.**

## Transfering variables from one step to another

### The DeepSquare way

You can add `KEY=value` in the `$DEEPSQUARE_ENV` file.

For example:

```shell title="step.command"
echo "KEY=value" >> "$DEEPSQUARE_ENV"
```

And this environment variable will be loaded at the next step.

### The manual way

Remember that `$STORAGE_PATH` is a shared storage between all nodes and steps?

You can add you own files if prefer not to use `$DEEPSQUARE_ENV`.

For example:

```shell title="step.command"
cat << EOF > "$STORAGE_PATH/my-env"
KEY=value
EOF
```

In the next step, you have to load the environment variables by yourself:

```shell title="step2.command"
set -o allexport; source "$STORAGE_PATH/my-env"; set +o allexport
```
