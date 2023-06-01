# Usage

Instead of using a `run` step, you can use a `use` step to invoke a module.

Specify:

- The `source` of the module, which is the git repository accessible via HTTPS. You can specify a git tag by adding `@<git-tag>` as a suffix. You can also specify the git SHA commit `@<git-commit-sha>`, which you can shorten to 7 characters. Otherwise, the HEAD commit will be used.
- The `args`, which must match the `inputs` field of the module.
- The `exportEnvAs`, which allows to use the `outputs` of the module as environment variables. The variables will be prefixed with the value of `exportEnvAs`.

Example of usage:

```json title="Hello-world workflow"
{
  "enableLogging": false,
  "resources": {
    "tasks": 1,
    "cpusPerTask": 8,
    "memPerCpu": 8000,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "name": "hello-world",
      "use": {
        "source": "github.com/deepsquare-io/workflow-module-example@v1",
        "args": [
          {
            "key": "WHO",
            "value": "me"
          }
        ],
        "exportEnvAs": "HELLO_WORLD"
      }
    },
    {
      "name": "repeat",
      "run": {
        "command": "echo ${HELLO_WORLD_RESULT}"
      }
    }
  ]
}
```

You can check the [module `github.com/deepsquare-io/workflow-module-example@v1` git repository](https://github.com/deepsquare-io/workflow-module-example/tree/v1).

:::caution

Be careful, when using a tag like `v1` or `v1.0.0`, make sure the author is trustworthy, otherwise you may fall victim to a supply chain attack.

There is a risk of falling victim to a supply chain attack by having an attacker alter the module code by "moving" the git tag, which compromises any workflow that uses the module.

Specifying the full SHA commit is the safest option, but specifying a tag may be more practical.

**A module has access to the entire Workflow specification! Beware of malicious modules!**

:::
