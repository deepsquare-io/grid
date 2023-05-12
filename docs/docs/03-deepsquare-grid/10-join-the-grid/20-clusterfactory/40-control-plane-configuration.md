# 2. Deploying the control plane services

This guide will only indicates how to deploy the software stack. It won't help you to learn what one software do. Please read the documentation of each software if you wish to learn more.

## Add the Git repository to ArgoCD

Argo CD can retrieve your repository from your Git hosting server, synchronize changes and deploy your Kubernetes manifests.

1. Create a local secret containing [an SSH deploy key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys#set-up-deploy-keys) and the git URL:

```yaml title="argo/default/secrets/my-repository-secret.yaml.local"
apiVersion: v1
kind: Secret
metadata:
  name: my-repository-secret
  namespace: argocd
  labels:
    argocd.argoproj.io/secret-type: repository
type: Opaque
stringData:
  sshPrivateKey: |
    -----BEGIN RSA PRIVATE KEY-----
    -----END RSA PRIVATE KEY-----
  type: git
  url: git@github.com:<your account>/<your repo>.git
```

2. Seal the secret and apply it:

   ```shell
   cfctl kubeseal
   kubectl apply -f argo/default/secrets/my-repository-sealed-secret.yaml
   ```

:::caution

Never commit a plaintext Secret. Always seal the secret!

:::

## Configure the storage

Before continuing futher, we need some storage provisionners to host our data.

By default, ClusterFactory providers two providers: nfs and local-path.

**Deploy the local-path-storage provisionner:**

```shell title="user@local:/ClusterFactory"
kubectl apply -f argo/local-path-storage
kubectl apply -f argo/local-path-storage/apps
```

The local-path provisionner creates volumes per node in the /opt/local-path-provisioner on the host. If you wish to change that path, you can edit the [local-path-storage.yaml](https://github.com/rancher/local-path-provisioner/blob/master/deploy/local-path-storage.yaml) and change the `source` of the ArgoCD Application `argo/local-path-storage/apps/local-path-storage-app.yaml`.

**Deploy the nfs StorageClass:**

Edit the `argo/volumes/dynamic-nfs.yaml` StorageClass and apply it:

```shell title="user@local:/ClusterFactory"
kubectl apply -f argo/volumes/dynamic-nfs.yaml
```

:::info

If you are using an another distributed file system, you can check the available Container Storage Interface (CSI) drivers [here](https://kubernetes-csi.github.io/docs/drivers.html).

:::

## 1. Deploy the LDAP server 389ds

1. Deploy Namespace and AppProject

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/ldap/
   ```

2. Configure the admin password in the secret. Create a `-secret.yaml.local` file:

   ```yaml title="argo/ldap/secrets/389ds-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: 389ds-secret
     namespace: ldap
   stringData:
     dm-password: <a password>
   ```

   Seal and apply the secret:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/ldap/secrets/389ds-sealed-secret.yaml
   ```

3. Deploy the LDAP IngressRoutes:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/ldap/ingresses/ldap-ingresses.yaml
   ```

4. Create a `Certificate` for LDAPS:

   ```yaml title="argo/ldap/certificates/ldap.example.com-cert.yaml"
   apiVersion: cert-manager.io/v1
   kind: Certificate
   metadata:
     name: ldap.example.com-cert
     namespace: ldap
   spec:
     secretName: ldap.example.com-secret
     issuerRef:
       name: private-cluster-issuer
       kind: ClusterIssuer
     commonName: ldap.example.com
     subject:
       countries: [CH]
       localities: [Lonay]
       organizationalUnits: []
       organizations: [Example Org]
       postalCodes: ['1027']
       provinces: [Laud]
       streetAddresses: [Chemin des Mouettes 1]
     duration: 8760h
     dnsNames:
       - ldap.example.com
       - dirsrv-389ds.ldap.svc.cluster.local
     privateKey:
       size: 4096
       algorithm: RSA
   ```

   Apply it:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/ldap/certificates/ldap.example.com-cert.yaml
   ```

5. Edit the ArgoCD Application to use our private fork:

   ```yaml title="argo/ldap/apps/389ds-app.yaml > spec > source"
   source:
     # You should have forked this repo. Change the URL to your fork.
     repoURL: git@github.com:<your account>/ClusterFactory.git
     # You should use your branch too.
     targetRevision: <your branch>
     path: helm/389ds
     helm:
       releaseName: 389ds

       # Create a values file inside your fork and change the values.
       valueFiles:
         - values-production.yaml
   ```

6. Edit the values file to use our certificate, secret and storage provisionner.

   ```yaml title="helm/389ds/values-production.yaml"
   tls:
     secretName: ldap.example.com-secret

   config:
     dmPassword:
       secretName: '389ds-secret'
       key: 'dm-password'
     suffixName: 'dc=example,dc=com'

   initChownData:
     enabled: true

   persistence:
     storageClassName: 'dynamic-nfs'
   ```

   Edit the `suffixName` according to your need.

7. Commit and push:

   ```shell title="user@local:/ClusterFactory"
   git add .
   git commit -m "Added 389ds service"
   git push
   ```

   And deploy the Argo CD application:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/ldap/apps/389ds-app.yaml
   ```

   Check the [ArgoCD dashboard](https://argocd.internal) to see if everything went well.

8. Enter the shell of the container and executes these command:

   Enter the shell with:

   ```shell title="user@local:/ClusterFactory"
   kubectl exec -i -t -n ldap dirsrv-389ds-0 -c dirsrv-389ds -- sh -c "clear; (bash || ash || sh)"
   ```

   Initialize the database:

   ```shell title="pod: dirsrv-389ds-0 (namespace: ldap)
   dsconf localhost backend create --suffix dc=example,dc=com --be-name example_backend
   dsidm localhost initialise
   ```

   Add plugins:

   ```shell title="pod: dirsrv-389ds-0 (namespace: ldap)
   # Unique mail
   dsconf localhost plugin attr-uniq add "mail attribute uniqueness" --attr-name mail --subtree "opu=people,dc=example,dc=com"
   # Unique uid
   dsconf localhost plugin attr-uniq add "uid attribute uniqueness" --attr-name uid --subtree "ou=people,dc=example,dc=com"
   # Unique uid number
   dsconf localhost plugin attr-uniq add "uidNumber attribute uniqueness" --attr-name uidNumber --subtree "dc=example,dc=com"
   # Unique gid number
   dsconf localhost plugin attr-uniq add "gidNumber attribute uniqueness" --attr-name gidNumber --subtree "ou=groups,dc=example,dc=com"

   # Enable the plugins
   dsconf localhost plugin attr-uniq enable "mail attribute uniqueness"
   dsconf localhost plugin attr-uniq enable "uid attribute uniqueness"
   dsconf localhost plugin attr-uniq enable "uidNumber attribute uniqueness"
   dsconf localhost plugin attr-uniq enable "gidNumber attribute uniqueness"

   # Generate UID and GID when the value is -1
   dsconf localhost plugin dna config "UID and GID numbers" add \
   --type gidNumber uidNumber \
   --filter "(|(objectclass=posixAccount)(objectclass=posixGroup))" \
   --scope dc=example,dc=run\
   --next-value 1601 \
   --magic-regen -1
   dsconf localhost plugin dna enable
   ```

   Restart the server after the changes:

   ```shell title="user@local:/ClusterFactory"
   kubectl delete pod -n ldap dirsrv-389ds-0
   ```

   Re-enter the shell and add indexes:

   ```shell title="pod: dirsrv-389ds-0 (namespace: ldap)"
   dsconf localhost backend index add --index-type eq --attr uidNumber example_backend
   dsconf localhost backend index add --index-type eq --attr gidNumber example_backend
   dsconf localhost backend index add --index-type eq --attr nsSshPublicKey example_backend
   dsconf localhost backend index reindex example_backend
   ```

   Add necessary users and groups for SLURM:

   ```shell title="pod: dirsrv-389ds-0 (namespace: ldap)"
   dsidm -b "dc=example,dc=com" localhost group create \
     --cn cluster-users
   dsidm -b "dc=example,dc=com" localhost group modify cluster-users \
     "add:objectClass:posixGroup" \
     "add:gidNumber:1600"
   dsidm -b "dc=example,dc=com" localhost group create \
     --cn slurm
   dsidm -b "dc=example,dc=com" localhost group modify slurm \
     "add:objectClass:posixGroup" \
     "add:gidNumber:1501"

   dsidm -b "dc=example,dc=com" localhost user create \
     --uid slurm \
     --cn slurm \
     --displayName slurm \
     --homeDirectory "/dev/shm" \
     --uidNumber 1501 \
     --gidNumber 1501
   dsidm -b "dc=example,dc=com" localhost group add_member \
     slurm \
     uid=slurm,ou=people,dc=example,dc=com
   dsidm -b "dc=example,dc=com" localhost group add_member \
     cluster-users \
     uid=slurm,ou=people,dc=example,dc=com
   ```

## 2. Deploy LDAP connector

1. Configure LDAP connector secret. Create a `-secret.yaml.local` file:

   ```yaml title="argo/ldap/secrets/ldap-connector-env-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: ldap-connector-env
     namespace: ldap
   type: Opaque
   stringData:
     AVAX_ENDPOINT_WS: wss://testnet.deepsquare.run/ws
     JOBMANAGER_SMART_CONTRACT: '0xCD563d4704e8B1Cd9b6F1BE398f4A0921aB2A3b2' # DeepSquare MetaScheduler smart contract

     LDAP_URL: ldaps://dirsrv-389ds.ldap.svc.cluster.local:3636 # Kubernetes LDAP service domain name, change it if needed
     LDAP_CA_PATH: /tls/ca.crt
     LDAP_BIND_DN: 'cn=Directory Manager'
     LDAP_BIND_PASSWORD: <389ds secret>
   ```

   Seal the secret and apply it:

   ```shell
   cfctl kubeseal
   kubectl apply -f argo/ldap/secrets/ldap-connector-env-sealed-secret.yaml
   ```

2. Configure the `others/ldap-connector/overlays/production` overlay:

   Edit the `configmap.yaml` accordingly (`peopleDN` and `groupDN`) and edit the `deployment.yaml` (`secretName` for the CA).

3. Commit and push:

   ```shell title="user@local:/ClusterFactory"
   git add .
   git commit -m "Added ldap-connector service"
   git push
   ```

   And deploy the Argo CD application:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/ldap/apps/ldap-connector-app.yaml
   ```

   Check the [ArgoCD dashboard](https://argocd.internal) to see if everything went well.

## 3. Deploy MariaDB

1. Deploy Namespace and AppProject

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/mariadb/
   ```

2. Configure the passwords in the secret. Create a `-secret.yaml.local` file:

   ```yaml title="argo/mariadb/secrets/mariadb-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: mariadb-secret
     namespace: mariadb
   stringData:
     mariadb-root-password: <..>
     mariadb-replication-password: <...>
     mariadb-password: <...>
   ```

   Seal and apply the secret:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/mariadb/secrets/mariadb-sealed-secret.yaml
   ```

3. Configure the `values.yaml` of the mariadb helm subchart:

   ```yaml title="helm-subcharts/mariadb/values.yaml"
   mariadb:
     global:
       storageClass: 'local-path'

     auth:
       existingSecret: 'mariadb-secret'

     primary:
       configuration: |-
         [mysqld]
         skip-name-resolve
         explicit_defaults_for_timestamp
         basedir=/opt/bitnami/mariadb
         plugin_dir=/opt/bitnami/mariadb/plugin
         port=3306
         socket=/opt/bitnami/mariadb/tmp/mysql.sock
         tmpdir=/opt/bitnami/mariadb/tmp
         max_allowed_packet=16M
         bind-address=*
         pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid
         log-error=/opt/bitnami/mariadb/logs/mysqld.log
         character-set-server=UTF8
         collation-server=utf8_general_ci
         slow_query_log=0
         slow_query_log_file=/opt/bitnami/mariadb/logs/mysqld.log
         long_query_time=10.0

         # Slurm requirements
         innodb_buffer_pool_size=4096M
         innodb_log_file_size=64M
         innodb_lock_wait_timeout=900
         innodb_default_row_format=dynamic

         [client]
         port=3306
         socket=/opt/bitnami/mariadb/tmp/mysql.sock
         default-character-set=UTF8
         plugin_dir=/opt/bitnami/mariadb/plugin

         [manager]
         port=3306
         socket=/opt/bitnami/mariadb/tmp/mysql.sock
         pid-file=/opt/bitnami/mariadb/tmp/mysqld.pid

       nodeSelector:
         kubernetes.io/hostname: <host with local-path>

       resources:
         limits:
           memory: 2048Mi
         requests:
           cpu: 250m
           memory: 2048Mi

     secondary:
       replicaCount: 0

     ## Init containers parameters:
     ## volumePermissions: Change the owner and group of the persistent volume mountpoint to runAsUser:fsGroup values from the securityContext section.
     ##
     volumePermissions:
       enabled: false

     metrics:
       enabled: false

       serviceMonitor:
         ## @param metrics.serviceMonitor.enabled Create ServiceMonitor Resource for scraping metrics using PrometheusOperator
         ##
         enabled: false
   ```

   **Change the `mariadb.global.storageClass` according to your need**! If you are using `local-path`, the pod need to be stuck on a node by using the `mariadb.nodeSelector` like so:

   ```yaml title="helm-subcharts/mariadb/values-production.yaml"
   mariadb:
     nodeSelector:
       kubernetes.io/hostname: mn1.example.com
   ```

4. Edit the ArgoCD Application to use our private fork:

   ```yaml title="argo/mariadb/apps/mariadb-app.yaml > spec > source"
   source:
     # You should have forked this repo. Change the URL to your fork.
     repoURL: git@github.com:<your account>/ClusterFactory.git
     # You should use your branch too.
     targetRevision: <your branch>
     path: helm-subcharts/mariadb
     helm:
       releaseName: mariadb

       # Create a values file inside your fork and change the values.
       valueFiles:
         - values-production.yaml
   ```

5. Commit and push:

   ```shell title="user@local:/ClusterFactory"
   git add .
   git commit -m "Added mariadb service"
   git push
   ```

   And deploy the Argo CD application:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/mariadb/apps/mariadb-app.yaml
   ```

   Check the [ArgoCD dashboard](https://argocd.internal) to see if everything went well.

6. Enter the shell of the container:

   ```shell title="user@local:/ClusterFactory"
   kubectl exec -i -t -n mariadb mariadb-0 -c mariadb -- sh -c "clear; (bash || ash || sh)"
   ```

   Initialize the SLURM database:

   ```shell title="pod: mariadb-0 (namespace: mariadb)
   mysql -u root -p -h localhost
   # Enter your root password
   ```

   ```shell title="pod: mariadb-0 (namespace: mariadb) (sql)
   create user 'slurmdb'@'%' identified by '<your password>';
   grant all on slurm_acct_db.* TO 'slurmdb'@'%';
   create database slurm_acct_db;
   ```

   Rememeber the **slurmdb** password to deploy SLURM.

## 4. Deploy SLURM

## 5. Deploy the Supervisor

## 6. Deploy CVMFS Stratum 1

## 7. Deploy Grendel

## What's next
