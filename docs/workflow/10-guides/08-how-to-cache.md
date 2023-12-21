# How to use caches

Caches are hard-coded in specific paths. The two most important one are:

- `$DEEPSQUARE_SHARED_TMP`, the cache shared between nodes.
- `$DEEPSQUARE_DISK_TMP`, the cache on one node.

These directories are ALWAYS defined, and are free to use. Beware that the caches are managed by the provider, and therefore, can be susceptible of sudden deletion. You may report for abuse to us if this happens too often.

There are two ways to use these caches: via symbolic links, or via mount binds.

## Symbolic Link

A symbolic link is a simple directory redirection. On Linux, to create a symbolic link, you do `ln -s <source> <dest>`. This creates a special file at `<dest>` which redirect to source when writing at `<dest>/<child>`.

Because it is creating a file, make sure that there is no exiting directory of file at the `<dest>` path.

**Example:**

If you wish to cache `/path/to/be/cached` and this directory already exists, you should run:

```shell title="command"
# Delete path
rm -rf /path/to/be/cached
mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
ln -s $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

If there are existing files in the `/path/to/be/cached`, you may want to save them:

```shell title="command"
# Save path
mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
# Save files
mv /path/to/be/cached* /path/to/be/cached || true
rm -rf /path/to/be/cached
ln -s $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

## Mount binds

When using the caches or `$STORAGE_PATH`, it is very possible that you are not able to delete a directory. You can use a "mount namespace", which is able to isolate the mounts table, and allows you to execute `mount` even if you are not privileged.

Right now, to use a mount namespace, you can either use a network namespace or `mapUid: 0`.

**Before**

```yaml title="Workflow"
steps:
  - name: 'example'
    run:
      command: |
        rm -rf /path/to/be/cached
        mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
        ln -s $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

**After, with mount namespace via MapUid**

```yaml title="Workflow"
steps:
  - name: 'example'
    run:
      mapUid: 0
      command: |
        mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
        mount --bind $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

**After, with mount namespace via network namespace**

```yaml title="Workflow"
steps:
  - name: 'example'
    run:
      mapUid: 0
      mapGid: 0
      network: slirp4netns
      command: |
        mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
        mount --bind $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

This avoids to use `rm` which may throw a `Device or resource busy`.

If there are existing files in the `/path/to/be/cached`, you may want to save them:

```shell title="command"
mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
# Save files
mv /path/to/be/cached/* $DEEPSQUARE_SHARED_TMP/my-cache
mount --bind $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

:::warning

This feature might be incomplete. Please report any issue or request for feature.

:::
