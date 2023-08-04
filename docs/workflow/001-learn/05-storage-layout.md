# Storage Layout and Policy

When running a job, multiple volumes is available for use with different lifecycles and isolation. The working directory is one of them.

## About Linux permissions

Linux permissions define access rights for files and directories in a Linux-based operating system. These permissions are assigned to three different categories of users: the **owner** of the file or directory, the **group** to which the file or directory belongs, and all other users (**world**). Each category can have three types of permissions: **read, write, and execute**.

The permissions notation can be seen by running `ls -la`: `rwxr--r--` which means the owner can read, write and execute, the group can read, and others can read. This notation can also be written is octal: `755`.

**In the DeepSquare context, the default group can also be considered as the "world". All DeepSquare users are part of this common group**.

## Layout

| Volumes                       | Description                                                                                                                                                                    | Lifecycle                       | Permissions                      | Example of usage                                             |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------- | -------------------------------- | ------------------------------------------------------------ |
| `STORAGE_PATH`                | This is the **working directory** also known as a **scratch** volume. This volume is **shared between nodes**. This volume is automatically **cleared at the end** of the job. | Job duration.                   | **Owner: read, write, execute.** | Input, output files, etc...                                  |
| `DEEPSQUARE_SHARED_TMP`       | This is a volume **shared between nodes**. This volume is automatically **cleared based on the provider policy**.                                                              | Infrastructure provider policy. | **Owner: read, write, execute.** | ML models, datasets ...                                      |
| `DEEPSQUARE_SHARED_WORLD_TMP` | This is a volume **shared between nodes and users**. This volume is automatically **cleared based on the provider policy**.                                                    | Infrastructure provider policy. | **World: read, write, execute.** | Shared ML models, datasets, ... Common cache between users.  |
| `DEEPSQUARE_DISK_TMP`         | This is a volume **per node**. This volume is automatically **cleared based on the provider policy**.                                                                          | Infrastructure provider policy. | **World: read, write, execute.** | Output files per process, fallback if shared is too slow ... |
| `DEEPSQUARE_DISK_WORLD_TMP`   | This is a volume **per node shared between users**. This volume is automatically **cleared based on the provider policy**.                                                     | Infrastructure provider policy. | **World: read, write, execute.** |                                                              |

## About provider policies

These volumes are considered as **caches** which are **free of charge**.

As a consequence, providers may have the autonomy to implement various policies, such as periodic volume deletions or other actions. This setup might lead to potential concerns regarding data availability and retention, as the user does not have control over how long their volumes will be retained by the provider.

Since the user does not incur charges for these volumes, there might be instances where data is unexpectedly deleted, leading to data loss if not adequately backed up by the user. Consequently, users must be aware of the provider's policies and practices and implement their data backup strategies to mitigate the risks associated with these limitations in volume management.

**However, we are planning to implement a storage solution to host your data close to the data center.**

## Precautions

By utilizing the DeepSquare Grid, you acknowledge and agree that we are not responsible for the management, handling, or security of your data once it is distributed across various infrastructure providers. Our platform operates as a decentralized network, allowing data to be stored on different infrastructure providers based on availability and demand.

We undertake the verification of infrastructure providers to ensure a certain level of reliability and security. However, we cannot guarantee the actions or practices of each individual provider in the network. Each infrastructure provider is responsible for adhering to their own data management policies and practices.

Therefore, by using our service, you expressly release us, our affiliates, and our personnel from any liability or responsibility for any loss, damage, or harm that may arise from data corruption, data leaks, or any malicious treatment of your data stored on the decentralized cloud infrastructure.

It is recommended that you take additional precautions, such as encryption, permissions and regular data backups, to protect your valuable information and minimize potential risks. Remember that your decision to use the decentralized cloud service is at your own risk, and you are solely responsible for evaluating and managing the associated risks.

**PLEASE FOLLOW THESE PRACTICES:**

- **Use a secure transport:** Always transfer data over secure channels such as SSH or SSL/TLS (HTTPS).
- **Use the file system permissions:** `chmod` and `chown` to manage the permissions of your data or force the permission by using `umask` (for example `umask 077` sets the default permissions for newly created files to be accessible only by the owner and denies all access to group members and others).
- **Use checksums:** Verify the data integrity before using it to avoid Man-in-the-middle attacks (for example `sha256sum`).
- **Prefer containers**: Use containers to make sure to avoid side-effects and reduce the attack surface. DeepSquare will clean these containers.
- **Delete the files at the end of the jobs explicitly if necessary.**

Other practices may exists to ensure the **confidentiality**, **integrity** and **availability** of your data. Again, you are responsible for managing the risks.
