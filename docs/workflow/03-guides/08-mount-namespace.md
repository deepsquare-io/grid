# How to use `mount --bind` instead of symbolic link

When using the caches or `$STORAGE_PATH`, it is very possible that you are not able to delete a directory. You can use a "mount namespace", which is able to isolate the mounts table.

Right now, to use a mount namespace, you can either use a network namespace or `mapUid: 0`.

**Before**

```yaml
steps:
  - name: 'example'
    run:
      command: |
        rm -rf /path/to/be/cached
        mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
        ln -s $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

**After, with mount namespace via mapUid**

```yaml
steps:
  - name: 'example'
    run:
      mapUid: 0
      command: |
        mkdir -p $DEEPSQUARE_SHARED_TMP/my-cache
        mount --bind $DEEPSQUARE_SHARED_TMP/my-cache /path/to/be/cached
```

**After, with mount namespace via network namespace**

```yaml
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

:::warning

This feature might be incomplete.

:::
