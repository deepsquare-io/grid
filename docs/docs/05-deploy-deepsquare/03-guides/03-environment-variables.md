# Using environment variables

Environment variables are the best way to customize your environment without having to rebuild your container image.

## Priority

Environment variables can be set at the task level and at the step level.

From lowest to highest, the order of priority is:

- Job level
- DeepSquare defined variables (`DEEPSQUARE_INPUT`, `DEEPSQUARE_OUTPUT`, `DEEPSQUARE_TMP` `STORAGE_PATH`, `DEEPSQUARE_ENV`)
- Container environment variables (`<container rootfs>/etc/environment`)
- Step level

## Isolation

A step-level environment variable will not be transferred to the next step.

A job-level environment variable that has been overridden will not be transferred to the next step.

Variables defined by DeepSquare that has been overridden will not be transferred to the next step.

The only way to pass information between steps is to use shared storage and DeepSquare variables, like `STORAGE_PATH`.

## DeepSquare defined variables

The environement variable `$STORAGE_PATH` stores the path of the shared storage path. That shared storage only lives during the job, it is a scratch storage.

The environment variable `$DEEPSQUARE_TMP` stores the path to a per-site cache. The cache is flushed periodically according to the infrastructure provider's policy.

If running a container, `$STORAGE_PATH` will be defined at `/deepsquare`.

The other environment variables defined by DeepSquare are subdirectories in the `$STORAGE_PATH`.

| Environment variables | Value                              |
| --------------------- | ---------------------------------- |
| STORAGE_PATH          | /deepsquare if running a container |
| DEEPSQUARE_INPUT      | /deepsquare/input                  |
| DEEPSQUARE_OUTPUT     | /deepsquare/output                 |
| DEEPSQUARE_ENV        | /deepsquare/env                    |
| DEEPSQUARE_TMP        | /deepsquare/tmp                    |

**Always use environment variables and not hard-code the values to be future-proof.**

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

You can add you own files if prefer no to use `$DEEPSQUARE_ENV`.

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
