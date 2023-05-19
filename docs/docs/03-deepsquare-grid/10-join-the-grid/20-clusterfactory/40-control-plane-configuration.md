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
   dsconf localhost plugin attr-uniq add "mail attribute uniqueness" --attr-name mail --subtree "ou=people,dc=example,dc=com"
   # Unique uid
   dsconf localhost plugin attr-uniq add "uid attribute uniqueness" --attr-name uid --subtree "ou=people,dc=example,dc=com"
   # Unique uid number
   dsconf localhost plugin attr-uniq add "uidNumber attribute uniqueness" --attr-name uidNumber --subtree "ou=people,dc=example,dc=com"
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

   ```yaml title="others/ldap-connector/overlays/production/configmap.yaml"
   peopleDN: ou=people,dc=example,dc=com
   groupDN: cn=cluster-users,ou=groups,dc=example,dc=com
   addUserToGroup:
     memberAttributes:
       - member
   createUser:
     rdnAttribute: uid
     objectClasses:
       - nsPerson
       - nsOrgPerson
       - nsAccount
       - posixAccount
     userNameAttributes:
       - uid
       - displayName
       - cn
     defaultAttributes:
       homeDirectory:
         - /dev/shm
       gidNumber:
         - '1600'
       uidNumber:
         - '-1'
   ```

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

   You can see the default values [here](https://artifacthub.io/packages/helm/bitnami/mariadb).

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

### a. Secrets

1. Deploy the Namespace and AppProject:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/slurm-cluster/
   ```

2. Create the sssd configuration secret at `argo/slurm-cluster/secrets/sssd-secret.yaml.local`:

   ```yaml title="argo/slurm-cluster/secrets/sssd-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: sssd-secret
     namespace: slurm-cluster
   type: Opaque
   stringData:
     sssd.conf: |
       # https://sssd.io/docs/users/troubleshooting/how-to-troubleshoot-backend.html
       [sssd]
       services = nss,pam,sudo,ssh
       config_file_version = 2
       domains = example-ldap

       [sudo]

       [nss]

       [pam]
       offline_credentials_expiration = 60

       [domain/example-ldap]
       debug_level=3
       cache_credentials = True
       dns_resolver_timeout = 15

       override_homedir = /home/ldap-users/%u

       id_provider = ldap
       auth_provider = ldap
       access_provider = ldap
       chpass_provider = ldap

       ldap_schema = rfc2307bis

       ldap_uri = ldaps://dirsrv-389ds.ldap.svc.cluster.local:3636
       ldap_default_bind_dn = cn=Directory Manager
       ldap_default_authtok = <password>
       ldap_search_timeout = 50
       ldap_network_timeout = 60
       ldap_user_member_of = memberof
       ldap_user_gecos = cn
       ldap_user_uuid = nsUniqueId
       ldap_group_uuid = nsUniqueId

       ldap_search_base = ou=people,dc=example,dc=com
       ldap_group_search_base = ou=groups,dc=example,dc=com
       ldap_sudo_search_base = ou=sudoers,dc=example,dc=com
       ldap_user_ssh_public_key = nsSshPublicKey

       ldap_account_expire_policy = rhds
       ldap_access_order = filter, expire
       ldap_access_filter = (objectClass=posixAccount)

       ldap_tls_cipher_suite = HIGH
       # On Ubuntu, the LDAP client is linked to GnuTLS instead of OpenSSL => cipher suite names are different
       # In fact, it's not even a cipher suite name that goes here, but a so called "priority list" => see $> gnutls-cli --priority-list
       # See https://backreference.org/2009/11/18/openssl-vs-gnutls-cipher-names/ , gnutls-cli is part of package gnutls-bin
   ```

   We need the SSSD configuration to authenticate the users from DeepSquare. There is a SSS Daemon running on each SLURM container.

   The `ldap_uri` points to the 389ds Kubernetes service. The internal addresses of the Kubernetes cluster follow the format `<service>.<namespace>.svc.cluster.local`. When 389ds was deployed, a Kubernetes service was also deployed.

   Replace `<password>`, `ldap_search_base`, `ldap_group_search_base` and `ldap_sudo_search_base` based on your configuration.

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/sssd-sealed-secret.yaml
   ```

3. Generate a JWT Key for the slurm controller:

   ```shell title="user@local:/ClusterFactory"
   ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key
   ```

   This key is used for the SLURM Rest API which we do not use, but is necessary to start the container.
   Put the key in the `argo/slurm-cluster/secrets/slurm-secret.yaml.local` file:

   ```yaml title="argo/slurm-cluster/secrets/slurm-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: slurm-secret
     namespace: slurm-cluster
   type: Opaque
   stringData:
     jwt_hs256.key: |
       -----BEGIN RSA PRIVATE KEY-----
       ...
       -----END RSA PRIVATE KEY-----
   ```

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/slurm-sealed-secret.yaml
   ```

   You can delete the generated `jwt_hs256.key`.

4. Generate a MUNGE Key for **all the nodes (login, controller, db, compute)**.

   ```shell title="user@local:/ClusterFactory"
   dd if=/dev/urandom bs=1 count=1024 > munge.key
   ```

   **The MUNGE key is the key used for symmetric encryption between the SLURM services. In the configuration of the compute plane, we will need this generated key. Don't lose it!**

   Encode the key in base64:

   ```shell title="user@local:/ClusterFactory"
   cat munge.key | base64
   ```

   Store the value in the `argo/slurm-cluster/secrets/munge-secret.yaml.local`:

   ```yaml title="argo/slurm-cluster/secrets/munge-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: munge-secret
     namespace: slurm-cluster
   type: Opaque
   data:
     munge.key: |
       9hIkThiXtf2Q7qJQ0coDkZGvvuTh9vddFv8guzIrz1Um29Gtxg0CTU4DmrLMBixtetTy/GMs/P7k
       BUwipOoQgM7ALok3+uMWK9zvqmMPruhtW5khw0meV9jsPdc2paZdQFsLIB3nFGhobMl16SOQajy7
       bfWaS6olAD9FksFRcumiyzRMuLYQ+yj1l7yuPSYAtQ29AU9aBPtgCH6pK7R+LQddz5dRJyMM9Doq
       CC+7cEITDJfW+Us+zvXzv9iCdYWuwovvjune6VCuxniRofKJjXB9QQFgKQ27jZuptKNZ5bMIdd46
       +vC0G/MQy/7gckTM0KcYvhoEvijJLSm3LMCKVr+U/XOwtKzeqDPtocQZBTTySZRO7i0RZTvAeuu0
       fwPFY2tUdA35/Ij0Yb8m+ZxwNvCKuPKYNbwaJxGHmluEhUu/H5iXSsD0ekhq7B4OFwieHnGxKxQV
       8I3QU++EF7uub44G3F0huLaFTYjAO09GMUoh5TZ7W5a9z8LUWsaz/yPrV0KOYgm+B3M/X5aNiTXz
       fs8h6+0Oun+j4i+pWlnWdkI37tWU0lzCgJG2M17l9RZfsou+Bcogu0vn1L+XnRRhQoUA4KqfPGcv
       6ybVPHaXCQNryTHxiqQLL7p7hL9RRUHkr/J54ucv9kkKibJBuR3A/7SHCHF745VlrgoqATZzpRD6
       AKtJkeKXWMfrzjp2iAUa7Veb5C/TDAnlWCKRFh08ayaJ3eBzFMEzb8WVR/+zFMeJWWbJR/d1CXSR
       UwM5+ajBWiYc8Mbv99zFGHiwHKMjG0DiQVTqyKovR68v5+VBaq8gzXyRO4Zp4Ij2wUm8ybsFhTnw
       8ijZRLNI4Bh1CKyCn5Z+YCvqvnPZw2Oix0sS1CDh5B1PH88essu1rUmbKuI4VcRaoKUUPngV/Yb4
       GsRcMRjApg91KoyCQ/TyvyaKUcsv+5FfsCR5QchxdpWy6JoNo/iGA+Uy/q7qQAWlx2Y9NNooSsTA
       yCv4bdZ5kffWtgpB1qx3oR2oBbD9FtDV3hyVtH4VvwpajLfi41eXhCFkeZblTrz12IVhsLsjoua4
       SioOmEoAdmw+w3Q+eg6VyHq4CLJLN2KLzxMxbbGyhA1GblkCAYA5YwUYEhGH+P36Fck8QYIxt0JG
       ycLveLexxQY74YaXYZGzoKeW09GXcfO+Dm0Ufr1nkuV6NSt+2z8RrKtHhg5WDANf8yn1TQF3gOMM
       H/ui38Iq+gU+reEg/yKsmA45Tuo4lcznmS9V6kohQUX8T9gK02vQNc/7z0C/rRvvJzZLE10NwvL9
       UdpQ/stRI618zNCF9SXYGagOo8Ks1IMmCVBVvE4E2/AiNe9d97rLGsKlC6s9TJuESl1XxAlJLA==
   ```

   **Make sure it is `.data` not `.stringData`!**

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/munge-sealed-secret.yaml
   ```

5. Configure the SLURM DB by creating a secret file at `argo/slurm-cluster/secrets/slurmdbd-conf-example-1-secret.yaml.local`:

   ```yaml title="argo/slurm-cluster/secrets/slurmdbd-conf-example-1-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
   name: slurmdbd-conf-example-1-secret
   namespace: slurm-cluster
   type: Opaque
   stringData:
   slurmdbd.conf: |
     # See https://slurm.schedmd.com/slurmdbd.conf.html
     ### Main
     DbdHost=slurm-cluster-example-1-db-0
     SlurmUser=slurm # user IDs have to be identical accross the cluster, otherwise munge auth fails

     ### Logging
     DebugLevel=debug5 # optional, defaults to 'info'. Possible values: fatal, error, info, verbose, debug, debug[2-5]
     LogFile=/var/log/slurm/slurmdbd.log
     PidFile=/var/run/slurmdbd.pid
     LogTimeFormat=thread_id

     AuthAltTypes=auth/jwt
     AuthAltParameters=jwt_key=/var/spool/slurm/jwt_hs256.key

     ### Database server configuration
     StorageType=accounting_storage/mysql
     StorageHost=mariadb.mariadb.svc.cluster.local
     StorageUser=slurmdb
     StoragePass=<The password of the slurmdb user of mariadb>
     StorageLoc=slurm_acct_db
   ```

   Make sure that the `Storage*` parameters are correctly set.

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/slurmdbd-conf-example-1-sealed-secret.yaml
   ```

6. We need to add the certificate of the CA to the SLURM containers, so that SSSD is able to verify the certificate of the LDAP server when using LDAPS. We've created a `ca-key-pair` when deploying the core of ClusterFactory. Fetch the secret with `kubectl get secrets -n cert-manager ca-key-pair -o yaml` and copy the content of `data."tls.crt"` to a new secret at `argo/slurm-cluster/secrets/local-ca-secret.yaml.local`:

   ```yaml title="argo/slurm-cluster/secrets/local-ca-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: local-ca-secret
     namespace: slurm-cluster
   type: Opaque
   data:
     example.com.pem: <base64 encoded CA certificate>
   ```

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/local-ca-sealed-secret.yaml
   ```

7. The SLURM login node is running an SSH server. Therefore, we need to add SSH Host Keys and a SSH configuration. Generate the keys with:

   ```shell title="user@local:/ClusterFactory"
   mkdir -p ./etc/ssh
   ssh-keygen -A -f $(pwd) -C login
   mv ./etc/ssh/* .
   rmdir ./etc/ssh
   rmdir /etc
   ```

   And put the content of the keys inside a new secret at:

   ```yaml title="argo/slurm-cluster/secrets/login-sshd-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: login-sshd-secret
     namespace: slurm-cluster
   type: Opaque
   stringData:
     ssh_host_ecdsa_key: |
       -----BEGIN OPENSSH PRIVATE KEY-----
       ...
       -----END OPENSSH PRIVATE KEY-----
     ssh_host_ecdsa_key.pub: |
       ecdsa-sha2-nistp256 ... login
     ssh_host_ed25519_key: |
       -----BEGIN OPENSSH PRIVATE KEY-----
       ...
       -----END OPENSSH PRIVATE KEY-----
     ssh_host_ed25519_key.pub: |
       ssh-ed25519 ... login
     ssh_host_rsa_key: |
       -----BEGIN OPENSSH PRIVATE KEY-----
       ...
       -----END OPENSSH PRIVATE KEY-----
     ssh_host_rsa_key.pub: |
       ssh-rsa ... login
     sshd_config: |
       Port 22
       AddressFamily any
       ListenAddress 0.0.0.0
       ListenAddress ::

       HostKey /etc/ssh/ssh_host_rsa_key
       HostKey /etc/ssh/ssh_host_ecdsa_key
       HostKey /etc/ssh/ssh_host_ed25519_key

       PermitRootLogin prohibit-password
       PasswordAuthentication yes
       PubkeyAuthentication yes

       # Change to yes to enable challenge-response passwords (beware issues with
       # some PAM modules and threads)
       ChallengeResponseAuthentication no
       GSSAPIAuthentication no

       UsePAM yes

       X11Forwarding yes
       PrintMotd no
       AcceptEnv LANG LC_*

       # override default of no subsystems
       Subsystem sftp	/usr/lib/openssh/sftp-server

       AuthorizedKeysCommand /authorizedkeys/custom-authorizedkeys "%u"
       AuthorizedKeysFile .ssh/authorized_keys
       AuthorizedKeysCommandUser nobody

       GSSAPIAuthentication no
       GSSAPICleanupCredentials no

       PrintMotd no

       AcceptEnv LANG LC_CTYPE LC_NUMERIC LC_TIME LC_COLLATE LC_MONETARY LC_MESSAGES
       AcceptEnv LC_PAPER LC_NAME LC_ADDRESS LC_TELEPHONE LC_MEASUREMENT
       AcceptEnv LC_IDENTIFICATION LC_ALL LANGUAGE
       AcceptEnv XMODIFIERS

       X11Forwarding yes
   ```

   Seal and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/slurm-cluster/secrets/login-sshd-sealed-secret.yaml
   ```

### b. Configuring the SLURM `values.yaml` file

1. "example-1" will be the name of our SLURM cluster. Create a `helm/slurm-cluster/values-example-1.yaml`. You can use `helm/slurm-cluster/values.yaml` to see the default values.

2. Let's start by plugging our secrets:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   sssd:
     # secret containing sssd.conf
     # Will be mounted in /secrets/sssd
     secretName: sssd-secret

   munge:
     # secret containing munge.key
     # Will be mounted in /secrets/munge
     secretName: munge-secret

   # secret containing jwt_hs256.key
   # Will be mounted in /secrets/slurm and copied to /var/spool/slurm/jwt_hs256.key
   jwt:
     secretName: slurm-secret
   ```

3. Then edit the slurm configuration. Name the cluster:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   # ...
   slurmConfig:
     clusterName: example-1
   ```

   Add accounting settings:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   # ...
   slurmConfig:
     # ...
     accounting: |
       AccountingStorageType=accounting_storage/slurmdbd
       AccountingStorageHost=slurm-cluster-example-1-db.slurm-cluster.svc.cluster.local
       AccountingStoragePort=6819
       AccountingStorageTRES=gres/gpu
       AccountingStoreFlags=job_comment,job_env,job_script
   ```

   The `AccountingStorageHost` matches the URL of the SLURM DB Service. The name follows this pattern: "slurm-cluster-&lt;cluster-name&gt;-db.&lt;namespace&gt;.svc.cluster.local".

   We allow to store the job comments, environment variables and job scripts inside the database.

   Add the default resources allocation:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   # ...
   slurmConfig:
     # ...
     defaultResourcesAllocation: |
       DefCpuPerGPU=8
       DefMemPerCpu=1000
   ```

   We recommend dividing the resources by GPU. For example, if your system consists of 16 CPUs, 2 GPUs, 16000 MiB of RAM, then the parameters are the same as described above.

   Add the nodes, partitions and gres (generic resources):

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   # ...
   slurmConfig:
     # ...
     nodes: |
       NodeName=cn1 CPUs=16 Boards=1 SocketsPerBoard=1 CoresPerSocket=8 ThreadsPerCore=2 RealMemory=15000 Gres=gpu:2
       # You can also group:
       #NodeName=cn[2-8] CPUs=64 Boards=1 SocketsPerBoard=1 CoresPerSocket=32 ThreadsPerCore=2 RealMemory=125000 Gres=gpu:3

     partitions: |
       PartitionName=main Nodes=cn1 Default=YES MaxTime=INFINITE State=UP OverSubscribe=EXCLUSIVE

     gres: |
       NodeName=cn1 File=/dev/nvidia[0-1] AutoDetect=nvml
   ```

   :::caution

   **You would normally use `slurmd -C` on the compute nodes to get the real configuration. We recommend subtracting `RealMemory` by 1G.**

   If the node is registered as INVALID by SLURM, this means that the configuration of the node does not match what you described in `slurmConfig.nodes`.

   Normally you would define the following:

   - CPUs (Central processing units): Number of logical processors on the node, which is equal to Boards × SocketsPerBoard × CoresPerSocket × ThreadsPerCore.
   - Boards: Number of baseboards in nodes with a baseboard controller.
   - SocketsPerBoard: Number of physical processor sockets/chips on a base board.
   - CoresPerSocket: Number of processor cores in a single physical processor socket.
   - ThreadsPerCore: Number of physical processor sockets/chips on the node. This is often 2 when enabling simultaneous hyperthreading/multithreading (SMT).

   If in doubt, you can leave a dummy value and recheck the actual value after provisioning the compute nodes.

   :::

   Add the following extra configuration:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   # ...
   slurmConfig:
     # ...
     extra: |
       LaunchParameters=enable_nss_slurm
       DebugFlags=Script,Gang,SelectType
       TCPTimeout=5

       # MPI stacks running over Infiniband or OmniPath require the ability to allocate more
       # locked memory than the default limit. Unfortunately, user processes on login nodes
       # may have a small memory limit (check it by ulimit -a) which by default are propagated
       # into Slurm jobs and hence cause fabric errors for MPI.
       PropagateResourceLimitsExcept=MEMLOCK

       ProctrackType=proctrack/cgroup
       TaskPlugin=task/cgroup
       SwitchType=switch/none
       MpiDefault=pmix_v4
       ReturnToService=2
       GresTypes=gpu
       PreemptType=preempt/qos
       PreemptMode=REQUEUE
       PreemptExemptTime=-1
       Prolog=/etc/slurm/prolog.d/*
       Epilog=/etc/slurm/epilog.d/*
       RebootProgram="/usr/sbin/reboot"

       # Federation
       FederationParameters=fed_display

       JobCompType=jobcomp/provider
       JobAcctGatherType=jobacct_gather/cgroup
   ```

   The most important line is `JobCompType=jobcomp/provider`, this enables the DeepSquare Provider Completion Plugin and allows proper job state reporting to DeepSquare.

   For the other parameters, you can read about them in the [official slurm.conf documentation](https://slurm.schedmd.com/slurm.conf.html).

4. Let's configure the deployment of the SLURM DB:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   db:
     enabled: true

     config:
       secretName: slurmdbd-conf-example-1-secret

     # Load new CA
     command: ['sh', '-c', 'update-ca-trust && /init']

     volumeMounts:
       - name: ca-cert
         mountPath: /etc/pki/ca-trust/source/anchors/

     # Extra volumes
     volumes:
       - name: ca-cert
         secret:
           secretName: local-ca-secret
   ```

   We mounted the CA certificate in the container's CA store and trusted the CA by running `update-ca-trust`. The configuration of the SLURM DB is already done via through the secret.

5. Let's configure the deployment of the SLURM Controller. **For stability reasons, we will expose the SLURM controller by using LoadBalancer instead of IPVLAN or HostPort. This doesn't limit the functionnality of the SLURM controller and is the recommended way of exposing a Kubernetes Service.**

   Add the following in the values file:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   controller:
     enabled: true
     replicas: 1

     command: ['sh', '-c', 'update-ca-trust && /init']

     useHostPort: false
     useNetworkAttachment: false

     persistence:
       storageClassName: 'dynamic-nfs'
       accessModes: ['ReadWriteOnce']
       size: 50Gi
       selectorLabels: {}

     # Extra volume mounts
     volumeMounts:
       - name: ca-cert
         mountPath: /etc/pki/ca-trust/source/anchors/

     # Extra volumes
     volumes:
       - name: ca-cert
         secret:
           secretName: local-ca-secret

     servicePerReplica:
       enabled: true

       # Use LoadBalancer to expose via MetalLB
       type: LoadBalancer

       annotations:
         metallb.universe.tf/address-pool: slurm-controller-example-1-pool
   ```

6. Let's configure the login node. **For stability reasons, we will choose to use a LoadBalancer service from Kubernetes to expose the SLURM login to the local network. This disables the `srun` command (run a command interactively) in the login node but does not block the `sbatch` command. If you want to launch jobs quickly on the login node, you can use `sbatch --wrap "command args"`.**

   Add the following in the values file:

   ```yaml title="helm/slurm-cluster/values-example-1.yaml"
   login:
     enabled: true
     replicas: 1

     command: ['sh', '-c', 'update-ca-trust && /init']

     sshd:
       secretName: login-sshd-secret

     tmp:
       medium: ''
       size: 50Gi

     initContainers:
       - name: download-ssh-authorized-keys
         imagePullPolicy: Always
         command: ['sh', '-c']
         image: ghcr.io/deepsquare-io/provider-ssh-authorized-keys:latest
         args:
           [
             'cp /app/provider-ssh-authorized-keys /out/provider-ssh-authorized-keys',
           ]
         volumeMounts:
           - name: custom-authorizedkeys
             mountPath: /out

       - name: prepare-custom-authorizedkeys
         imagePullPolicy: Always
         image: docker.io/library/busybox:latest
         command: ['sh', '-c']
         args:
           - |
             cat << 'EOF' > /out/custom-authorizedkeys
             #!/bin/sh
             # SSSD
             /usr/bin/sss_ssh_authorizedkeys "$1" || true

             # Our authorized keys
             /authorizedkeys/provider-ssh-authorized-keys --supervisor.tls --supervisor.tls.insecure --supervisor.endpoint supervisor.example.com:443 || true
             echo
             EOF
             chown root:root /out/custom-authorizedkeys
             chown root:root /out/provider-ssh-authorized-keys
             chmod 755 /out/custom-authorizedkeys
             chmod 755 /out/provider-ssh-authorized-keys
             chmod 755 /out
         volumeMounts:
           - name: custom-authorizedkeys
             mountPath: /out

     # Extra volume mounts
     volumeMounts:
       - name: ca-cert
         mountPath: /etc/pki/ca-trust/source/anchors/
       - name: ldap-users-example-1-pvc
         mountPath: /home/ldap-users
       - name: custom-authorizedkeys
         mountPath: /authorizedkeys
         readOnly: true

     # Extra volumes
     volumes:
       - name: ca-cert
         secret:
           secretName: local-ca-secret
       - name: ldap-users-example-1-pvc
         persistentVolumeClaim:
           claimName: ldap-users-example-1-pvc
       - name: custom-authorizedkeys
         emptyDir: {}

     service:
       enabled: true
       # Use LoadBalancer to expose via MetalLB
       type: LoadBalancer

       annotations:
         metallb.universe.tf/address-pool: slurm-example-1-pool

     # Expose via IPVLAN, can be unstable.
     # Using IPVLAN permits srun commands.
     net:
       enabled: false

     # Slurm REST API
     rest:
       enabled: false

     metrics:
       enabled: false
   ```

## 5. Deploy the Supervisor

## 6. Deploy CVMFS Stratum 1

## 7. Deploy Grendel

## What's next
