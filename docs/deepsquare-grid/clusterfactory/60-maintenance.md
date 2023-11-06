# Part 4: Maintaining a DeepSquare cluster

As you can see, ClusterFactory is quite heavy. So here are the recommended practices to maintain your cluster with ClusterFactory.

## Updating, Backup, Restore, Ejecting a controller

About maintaining the Kubernetes Cluster, read that [here](https://docs.clusterfactory.io/docs/guides/maintenance/updating-k0s-cluster).

Upgrading is safe and seamless so don't hesitate.

To update ClusterFactory, `git fetch upstream` and `git merge upstream/<ref>`. **Always solve the merge conflits. Updates to the `core` and `argo` directories will go to `core.example` and `argo.example`.**

## Updating the software stack

**To update the containers:**

In each `values.yaml` file, you can override the tag of the container.

If you are using `dev` containers, just kill the pods. If the `imagePullPolicy` is `Always`, it will pull a new image.

**To update a helm subchart:**

You can update the Helm subchart by bumping the version inside the `Chart.yaml`.

**To update the smart-contract address:**

Update the LDAP connector secrets and supervisor values.

If you are using `dev` containers, the smart-contract address is always the [latest `smart-contract` release](https://github.com/deepsquare-io/grid/releases).

If you are using a stable version, the smart-contract address will be the one used on [app.deepsquare.run](https://app.deepsquare.run).

**To update the compute plane:**

If you are using `root=live:https://sos-ch-dk-2.exo.io/osimages/squareos-9.2/squareos-9.2.squashfs` in your kernel command line parameters, rebooting the nodes will always update the OS image.

**To update the postscripts:**

Update the git repository that you've created. Rebooting the nodes will execute the latest postscripts.

**In case of major upgrade (new OS image + new container images):**

Make sure the SLURM version matches (`slurmd --version`) between the SLURM controller and SLURM daemon. Besides that, everything is decoupled enough that there is no problem to update each component individually.

**To customize the OS image:**

ClusterFactory provides the [Packer recipe for building the SquareOS image](https://github.com/deepsquare-io/ClusterFactory/tree/main/packer-recipes/rocky9.2). Edit and use it to add your software.

## Monitoring

We recommend to deploy a [prometheus stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack) on Kubernetes and deploy node exporters on the compute nodes.

Use the stack with the [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md) so you can use the `ServiceMonitor` resources to easily configure the services to be monitored instead of using a static configuration.

## Single Responsibility

ClusterFactory tries to follow GitOps. You must force yourself to not use `kubectl` and `cfctl` directly.

Core resources (Volumes, Namespaces, AppProject) will always need to be deployed with `kubectl`. The rest can be integrated inside a Kustomization bundle or a Helm application.

You can expose the ArgoCD Dashboard so that Developpers can use the ArgoCD Dashboard to deploy their application.

You can also share the public certificate of `sealed-secret` by running:

```shell
kubeseal --fetch-cert --controller-namespace sealed-secrets --controller-name sealed-secrets > tls.crt
```

Then, developers can seal a secret by using `kubeseal` with `tls.crt`:

```yaml
kubeseal --cert tls.crt -o yaml -f my-secret.yaml.local
```

# Having issues ?

Feel free to ask us questions or ask for help on Discord!

The product is still young and the documentation certainly needs some fine-tuning.
