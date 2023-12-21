# Usage

Instead of using a `run` step, you can use a `use` step to invoke a module.

Specify:

- The `source` of the module, which is the git repository accessible via HTTPS.
  - You can specify a git tag by adding `@<git-tag>` as a suffix.
  - You can also specify the git SHA commit `@<git-commit-sha>`, which you can shorten to 7 characters. Otherwise, the HEAD commit will be used.
- The `args`, which must match the `inputs` field of the module.
- The `exportEnvAs`, which allows to use the `outputs` of the module as environment variables. The variables will be prefixed with the value of `exportEnvAs`.

:::info

**About monorepos of modules**

To specify the source of a module inside a monorepo:

- The format of the source is the following: `<host>/<owner>/<repo>[/<path to directory containing module.yaml>]`. If the git repository is a monorepo of modules, you can specify the path to the directory containing a `module.yaml` by appending `/path`
  - Example: `github.com/deepsquare-io/workflow-module-example/other-module-example`
- The git tag format should be `<path>/<ref>` and the source `<host>/<owner>/<repo>/<path>@<ref>`.
  - Example: The git tag is `other-module-example/v1`.
    Therefore, the source is `github.com/deepsquare-io/workflow-module-example/other-module-example@v1`

:::

Example of usage:

```yaml title="Hello-world workflow"
enableLogging: false

resources:
  tasks: 1
  cpusPerTask: 8
  memPerCpu: 8000
  gpus: 0

steps:
  - name: hello-world
    ## Use a module
    use:
      ## Address of the module.
      source: github.com/deepsquare-io/workflow-module-example@v1
      ## Pass environment variable.
      args:
        - key: WHO
          value: me
      ## Export the outputs of the module.
      exportEnvAs: HELLO_WORLD
  - name: repeat
    run:
      ## Use the outputs of the module.
      command: echo ${HELLO_WORLD_RESULT}
```

You can check the [module `github.com/deepsquare-io/workflow-module-example@v1` git repository](https://github.com/deepsquare-io/workflow-module-example/tree/v1).

:::warning

Be careful, when using a tag like `v1` or `v1.0.0`, make sure the author is trustworthy, otherwise you may fall victim to a supply chain attack.

There is a risk of falling victim to a supply chain attack by having an attacker alter the module code by "moving" the git tag, which compromises any workflow that uses the module.

Specifying the full SHA commit is the safest option, but specifying a tag may be more practical.

**A module has access to the entire Workflow specification! Beware of malicious modules!**

:::
