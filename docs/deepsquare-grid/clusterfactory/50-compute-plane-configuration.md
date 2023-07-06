# Part 3: Deploying the compute plane

## About network booting, postscripts and BMC

**About network booting**

Grendel is a provisioning service for bare-metal server. It hosts a DHCP, TFTP, PXE and HTTP server.

This is the typical procedure of network booting a linux kernel and initrd:

1. On the compute node, the BIOS or UEFI firmware is configured for PXE. It initializes the network interface card (NIC) and performs a Power-On Self-Test (POST) on the PXE firmware.
2. The PXE firmware initializes a DHCP (Dynamic Host Configuration Protocol) request packet to the local network. The packet includes information such as the client's MAC address and a request of an IP address.
3. Grendel's DHCP server responds with an IP address based on the MAC address, and attach the IP address of the TFTP server.
4. Grendel's TFTP server send a iPXE bootloader to the compute node. The iPXE bootloader fetches the Linux kernel and initial RAM disk (initrd) from the Grendel TFTP server.
5. The kernel is initialized on the compute node and boot on the initrd. Based on the kernel command line parameters, the inird will boot on the OS.

   We use this set on kernel command line parameters:

   ```shell
   console=ttyS0 console=tty0 root=live:https://sos-ch-dk-2.exo.io/osimages/squareos-8.6/squareos-8.6.squashfs BOOTIF=01-{{ $.nic.MAC | toString | replace ":" "-" }} grendel.hostname={{ $.host.Name }} grendel.address=http://grendel.example.com rd.live.overlay.readonly=1 rd.live.overlay.overlayfs=1 rd.neednet=1
   ```

   - `console=ttyS0`: This parameter specifies that the system should output the console messages to the serial port ttyS0. It allows for console access and debugging through a serial connection.
   - `console=tty0`: This parameter indicates that the system should also output console messages to the first virtual terminal (tty0). It provides a text-based interface for interacting with the system.
   - `root=live:https://sos-ch-dk-2.exo.io/osimages/squareos-8.6/squareos-8.6.squashfs`: This parameter sets the root file system to be loaded from a live HTTP source. It specifies the URL `https://sos-ch-dk-2.exo.io/osimages/squareos-8.6/squareos-8.6.squashfs` as the location from which the root file system should be retrieved.
   - `BOOTIF=01-{{ $.nic.MAC | toString | replace ":" "-" }}`: This parameter sets the BOOTIF (Boot Interface) variable to a value generated based on the network interface MAC address. The value follows the Go templating format. Grendel will fill the field.
   - `grendel.hostname={{ $.host.Name }}`: This parameter sets the hostname of the system. Grendel will also fill the field.
   - `grendel.address=http://grendel.example.com`: This parameter specifies the address of Grendel. It is used to fetch the postscripts and private keys.
   - `rd.live.overlay.readonly=1`: This parameter indicates that the live system's overlay should be mounted as read-only. The overlay typically contains temporary or volatile data that overlays the read-only root file system.
   - `rd.live.overlay.overlayfs=1`: This parameter specifies that the overlay file system should be implemented using OverlayFS, which is a union file system used to overlay one file system onto another.
   - `rd.neednet=1`: This parameter indicates that the system requires network connectivity to boot successfully. It ensures that the necessary network-related services and configurations are set up during the boot process.

6. After booting the OS, SquareOS will fetch the postscripts from Grendel and change its hostname.

To summary:

```mermaid
sequenceDiagram
  participant http as HTTP server
  participant grendel as Grendel
  participant compute as Compute
  compute->>+grendel: PXE: DHCP Request with MAC address
  grendel->>-compute: DHCP Response with IP and TFTP server IP
  compute->>+grendel: PXE: TFTP Request for iPXE bootloader
  grendel->>-compute: TFTP Response with iPXE bootloader
  compute->>compute: PXE: Boot on iPXE
  compute->>+grendel: iPXE: Request kernel and inird
  grendel->>-compute: Response with kernel and inird
  compute->>compute: PXE: Boot with kernel on inird
  compute->>+http: inird: Request OS image
  http->>-compute: Response with OS image
  compute->>compute: inird: Setup OverlayFS (lowerdir: OS image, upperdir: tmpfs), boot on OverlayFS
  compute->>+grendel: OS: Request postscripts
  grendel->>-compute: Response with postscripts
  compute->>compute: Run postscripts
```

**About postscripts**

You can store a postscript on Grendel. However, we want to follow GitOps principles. Therefore, this postscript will be used to pull postscripts from Git and execute them.

**About BMC**

If your servers are equipped with a Baseboard Management Controller (BMC), Grendel can work with IPMI. You can start, stop, restart, soft reboot and get the status of your servers without having to access the server directly.

## 1. Postscripts setup

1. Create a private Git repository on your favorite Git hosting service and add a [**read-only** deploy key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/managing-deploy-keys). You can name that repository "postscripts".

**We will create a Kubernetes Secret with that key.**

2. Copy the content of [deepsquare-io/postscripts](https://github.com/deepsquare-io/postscripts):

   ```shell title="user@local:/postscripts"
   curl -fsSL https://github.com/deepsquare-io/postscripts/archive/refs/heads/main.tar.gz -o main.tar.gz
   tar --strip-components=1 -xvf main.tar.gz
   rm -f main.tar.gz
   ```

3. Add the CA certificate in the `./files/certs` directory:

   ```shell title="user@local:/ClusterFactory"
   kubectl get secrets -n cert-manager ca-key-pair -o jsonpath='{.data.tls\.crt}' | base64 --decode > /postscripts/files/certs/private-ca.pem
   ```

4. Add the munge.key in the `./files/munge` directory:

   ```shell title="user@local:/ClusterFactory"
   kubectl get secrets -n slurm-cluster munge-secret -o jsonpath='{.data.munge\.key}' | base64 --decode > /postscripts/files/munge/munge.key
   ```

5. Add the sssd.conf in the `./files/sssd` directory. Don't forget to change the LDAP endpoint:

   ```properties title="files/sssd/sssd.conf"
   # https://sssd.io/docs/users/troubleshooting/how-to-troubleshoot-backend.html
   [sssd]
   services = nss,pam,sudo,ssh
   config_file_version = 2
   domains = deepsquare-ldap

   [sudo]

   [nss]
   filter_groups = root
   filter_users = root
   reconnection_retries = 3
   entry_cache_timeout = 300

   [pam]
   offline_credentials_expiration = 60
   reconnection_retries = 3
   offline_credentials_expiration = 2
   offline_failed_login_attempts = 3
   offline_failed_login_delay = 5

   [domain/deepsquare-ldap]
   debug_level=3
   cache_credentials = True

   override_homedir = /home/ldap-users/%u

   id_provider = ldap
   auth_provider = ldap
   access_provider = ldap
   chpass_provider = ldap

   ldap_schema = rfc2307bis

   ldap_uri = ldaps://ldap.example.com
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

6. In the `./files/ssh` directory, generate SSH host keys for your nodes:

   ```properties title="user@local:/postscripts/files/ssh"
   mkdir -p ./etc/ssh
   ssh-keygen -A -f $(pwd) -C compute
   mv ./etc/ssh/* .
   rmdir ./etc/ssh
   rmdir ./etc
   ```

7. In the `./files/slurm/slurmd_defaults` file, make sure the `CONTROLLER_HOST` is pointing the slurm controller address defined in the CoreDNS configuration.

   ```shell title="files/slurm/slurmd_defaults"
   CONTROLLER_HOST=slurm-cluster-example-1-controller-0.example.com
   PMIX_MCA_ptl=^usock
   PMIX_SYSTEM_TMPDIR=/var/empty
   PMIX_MCA_gds=hash
   ```

8. Lastly, **correct the `./postscripts/fs_mount` file.** Make sure the NFS (or other shared filesystem) mounts are correct. There is also a disk mount, **comment that line for now**.

   These are the required storage layout:

   | Path                                         | Description                                                                                                                                                         |
   | -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
   | /home/ldap-users                             | A shared volume hosting the home of the users.                                                                                                                      |
   | (optional but recommended) /opt/cache/enroot | A shared volume for Enroot, which is storing container layers.                                                                                                      |
   | /opt/cache/shared                            | A shared scratch volume for jobs. Must be a fast shared/distributed storage. Contains the jobs working directories.                                                 |
   | /opt/cache/persistent                        | A shared cache volume for jobs. It must be cleaned up after 7 days or more. Users can use this volume to temporarily host large data such as ML models or datasets. |
   | /mnt/persistent                              | A persistent volume on disk. It should be fast. It will be used by users to temporarily host large data such as ML models or datasets.                              |
   | /mnt/scratch                                 | A persistent volume on disk. It should be fast. It will be used by the container runtimes Apptainer and Enroot to create the runtime container filesystem.          |

   To help you get organized, we recommend creating directories inside `/mnt` like `/mnt/nfs` or `/mnt/disk`, and then making a symbolic link from the required path like `/opt/cache/shared` to your storage provisioner, like `/mnt/nfs/cache/shared`.

9. Commit and push:

   ```shell title="user@local:/postscripts"
   git add .
   git commit -m "Added postscripts"
   git push
   ```

## 2. BIOS setup

Configure your compute nodes to use PXE (network boot) and fetch the MAC addresses of your network interfaces.

Since this is vendor-specific, please check your motherboard manufacturer manual.

## 3. Deploying Grendel

1. Start by deploying the namespace and AppProject:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/provisioning/
   ```

2. Create a secret with the deploy key at `argo/provisioning/secrets/postscript-privatekey-secret.yaml.local` and fill it with:

   ```yaml title="argo/provisioning/secrets/postscript-privatekey-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: postscript-privatekey-secret
     namespace: provisioning
   type: Opaque
   stringData:
     key: |
       -----BEGIN OPENSSH PRIVATE KEY-----
       ...
       -----END OPENSSH PRIVATE KEY-----
   ```

   Seal the secret and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/provisioning/secrets/postscript-privatekey-sealed-secret.yaml
   ```

3. Create a secret for the Grendel configuration at `argo/provisioning/secrets/grendel-secret.yaml.local` and fill it with:

   ```yaml title="argo/provisioning/secrets/grendel-secret.yaml.local"
   apiVersion: v1
   kind: Secret
   metadata:
     name: grendel-secret
     namespace: provisioning
   type: Opaque
   stringData:
     grendel.toml: |
       dbpath = ":memory:"
       loggers = {cli="on", tftp="off", dhcp="on", dns="off", provision="on", api="on", pxe="on"}
       admin_ssh_pubkeys = [] # Unused

       [provision]
       listen = "0.0.0.0:80"
       # TTL in seconds for provision tokens. Defaults to 1 hour
       token_ttl = 3600
       root_password = "" # Unused
       default_image = "" # Unused
       repo_dir = "/var/lib/grendel"

       [dhcp]
       listen = "0.0.0.0:67"
       lease_time = "24h"
       dns_servers = []
       domain_search = []
       mtu = 1500
       proxy_only = false
       router_octet4 = 0

       # CHANGE THIS
       subnets = [
           {gateway = "192.168.0.1/24", dns="192.168.0.100", domainSearch="example.com", mtu="1500"}
       ]

       [dns]
       listen = "0.0.0.0:53"
       ttl = 86400

       [tftp]
       listen = "0.0.0.0:69"

       [pxe]
       listen = "0.0.0.0:4011"

       [api]
       socket_path = "/var/run/grendel/grendel-api.socket"

       [client]
       api_endpoint = "/var/run/grendel/grendel-api.socket"
       insecure = false

       [bmc]
       user = ""
       password = ""

       [discovery]
       user = ""
       password = ""
       domain = ""
   ```

   Seal the secret and apply it:

   ```shell title="user@local:/ClusterFactory"
   cfctl kubeseal
   kubectl apply -f argo/provisioning/secrets/grendel-sealed-secret.yaml
   ```

4. Fetch the MAC address of your node, create the values file at `helm/grendel/values-production.yaml` and fill it with:

   1. Fill the Grendel configuration:

      ```yaml title="helm/grendel/values-production.yaml"
      config:
        ## Secret containing grendel.toml
        secretName: grendel-secret
        secretKey: grendel.toml

        hosts:
          - name: cn1
            provision: true
            boot_image: squareos-8.6
            firmware: snponly-x86_64.efi
            interfaces:
              - ip: 192.168.0.51/24
                mac: aa:bb:cc:11:22:33
                bmc: false
              # Only if you have a BMC like IPMI
              - ip: 192.168.1.51/32
                bmc: true

        images:
          - name: squareos-8.6
            kernel: '/var/lib/grendel/vmlinuz-4.18.0-372.19.1.el8_6.x86_64'
            initrd:
              - '/var/lib/grendel/initramfs-4.18.0-372.19.1.el8_6.x86_64.img'
            cmdline: console=ttyS0 console=tty0 root=live:https://sos-ch-dk-2.exo.io/osimages/squareos-8.6/squareos-8.6.squashfs BOOTIF=01-{{ $.nic.MAC | toString | replace ":" "-" }} grendel.hostname={{ $.host.Name }} grendel.address=http://grendel.example.com rd.live.overlay.readonly=1 rd.live.overlay.overlayfs=1 rd.neednet=1

        postscript: |
          #!/bin/sh

          set -ex

          # Fetch encrypted deploy key
          curl --retry 5 -fsSL http://grendel.example.com/repo/key
          chmod 600 /key

          # Cloning git repo containing postscripts.
          mkdir -p /configs
          GIT_SSH_COMMAND='ssh -i /key -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null' git clone git@github.com:<your account>/postscripts.git /configs
          if [ -f /configs/post.sh ] && [ -x /configs/post.sh ]; then
            cd /configs || exit 1
            ./post.sh "$1"
          fi
          rm -f /key

          # Security
          chmod -R g-rwx,o-rwx .
      ```

   2. Add the persistence parameters. The persistence is used to host the kernel and inird:

      ```yaml title="helm/grendel/values-production.yaml"
      # ...
      persistence:
        storageClassName: 'dynamic-nfs'
        accessModes: ['ReadWriteMany']
        size: 20Gi
        selectorLabels: {}
      ```

   3. Add the IPVLAN parameters.

      ```yaml title="helm/grendel/values-production.yaml"
      # ...
      net:
        # Kubernetes host interface
        masterInterface: eth0
        mode: l2
        type: ipvlan

        # https://www.cni.dev/plugins/current/ipam/static/
        ipam:
          type: static
          addresses:
            - address: 192.168.0.3/24
              gateway: 192.168.0.1
          routes:
            - dst: 0.0.0.0/0

        # https://kubernetes.io/docs/concepts/services-networkingdns-pod-service/#pod-dns-config
        dns:
          nameservers:
            - 8.8.8.8
      ```

   4. Mount the private key for GitOps-based postscripts:

      ```yaml title="helm/grendel/values-production.yaml"
      # ...
      ## Extra volumes
      volumes:
        - name: postscript-privatekey
          secret:
            defaultMode: 384
            secretName: postscript-privatekey-secret

      ## Extra volume mounts
      volumeMounts:
        - name: postscript-privatekey
          subPath: key.enc
          mountPath: /var/lib/grendel/key.enc
      ```

   5. (Optional) Setup IPMI:

      ```yaml title="helm/grendel/values-production.yaml"
      # ...
      ipmi:
        ingress:
          enabled: true
          ingressClass: 'traefik'

          annotations:
            cert-manager.io/cluster-issuer: private-cluster-issuer
            traefik.ingress.kubernetes.io/router.entrypoints: websecure
            traefik.ingress.kubernetes.io/router.tls: 'true'

          hosts:
            - ipmi.example.com

          path: /

          tls:
            - secretName: ipmi.example.com-secret
              hosts:
                - ipmi.example.com
      ```

5. Edit the ArgoCD Application to use our private fork:

   ```yaml title="argo/provisioning/apps/grendel-app.yaml > spec > source"
   source:
     # You should have forked this repo. Change the URL to your fork.
     repoURL: git@github.com:<your account>/ClusterFactory.git
     # You should use your branch too.
     targetRevision: <your branch>
     path: helm/grendel
     helm:
       releaseName: grendel

       # Create a values file inside your fork and change the values.
       valueFiles:
         - values-production.yaml

   destination:
     server: 'https://kubernetes.default.svc'
     namespace: provisioning
   ```

6. Commit and push:

   ```shell title="user@local:/ClusterFactory"
   git add .
   git commit -m "Added grendel service"
   git push
   ```

   And deploy the Argo CD application:

   ```shell title="user@local:/ClusterFactory"
   kubectl apply -f argo/provisioning/apps/grendel-app.yaml
   ```

   Check the [ArgoCD dashboard](https://argocd.internal) to see if everything went well.

7. Reboot the compute nodes and check the serial console to see if everything went well.

## Compute nodes configuration

The DeepSquare Grid needs a local disk/partition mounted at `/mnt/disk`. Format the disk/partition as ext4, xfs, zfs or btrfs4 and fetch its UUID. Assuming you want XFS, these are the command:

```shell "root@cn1:/root"
# List the disks
fdisk -l
# WARNING: ALL DATA WILL BE DELETED
# Delete the file system
wipefs -af /dev/nvme0n1
fdisk /dev/nvme0n1
# Fdisk command g: Create a new empty GPT partition table
g
# Fdisk command w: Write table to disk and exit
w
# Create a primary partition
parted -a optimal --script /dev/nvme0n1 mkpart primary 0% 100%
# Create a xfs filesystem on the primary partition (NOTE the nvme0n1p1, not nvme0n1!)
mkfs.xfs -f /dev/nvme0n1p1
# Change the UUID of the disk, you can use the one in the postscript
xfs_admin -U <uuid> /dev/nvme0n1p1
```

You can the **re-enable the disk mount in the fs_mount postscript.** Don't forget to commit and push!

## Congratulations!

You have officially setup your cluster! You may apply to join the DeepSquare Grid.

You can also launch some commands on the login pod to check if everything went well:

```shell title="pod: slurm-cluster-example-1-login-<hash> (namespace: ldap)"
sbatch --wrap "sleep 5 && hostname"
# Submitted batch job 1
squeue
#              JOBID PARTITION     NAME     USER ST       TIME  NODES NODELIST(REASON)
#                  1      main     wrap     root  R       0:01      1 cn1
sinfo
# PARTITION   AVAIL  TIMELIMIT  NODES  STATE NODELIST
# main*          up   infinite      1    mix cn1
```
