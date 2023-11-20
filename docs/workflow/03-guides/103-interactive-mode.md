# Interactive mode

Interactive mode in an HPC environment allows users to directly interact with the system in real-time, offering a dynamic and responsive computing experience. This mode is particularly useful for debugging, code development, and iterative testing, as it enables users to receive instant feedback on their work

However, due to the decentralized aspect of DeepSquare, it is almost impossible to open an interactive session on the node. Each job runs without a terminal (TTY), which complicates user interactions.

There are two still ways to open an interactive session, similar to `docker exec -it` or `kubectl exec -it`, which are the commands used to open an interactive pseudo-TTY on Docker and Kubernetes respectively.

## tty2web (recommended)

[tty2web](https://github.com/kost/tty2web) is a command line tool that are able to open TTY sessions inside the container, written in Go. Thanks to Go, the tty2web is compiled statically, which means it can run anywhere.

This is similar to running an SSH server, but the front-end is a web client. Because it is a server, it opens a port which needs to be forwarded by using `bore`:

```yaml title="Simple Workflow"
steps:
  - name: 'interactive'
    run:
      container:
        image: curlimages/curl:latest
      env:
        # Set the terminal type. xterm is supported by tty2web
        - key: TERM
          value: xterm
      command: |
        curl -fsSL -o "$STORAGE_PATH/tty2web" https://github.com/kost/tty2web/releases/download/v3.0.3/tty2web_linux_amd64
        chmod +x "$STORAGE_PATH/tty2web"
        "$STORAGE_PATH/tty2web" --permit-write --port 8080 --credential admin:password sh
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 8080
```

The logs of the job will output a "bore.deepsquare.run" URL which contains the URL of the web TTY server:

```shell title="Example of output"
Generated HTTP URL:  http://ac21705f.bore.deepsquare.run
Generated HTTPS URL: https://ac21705f.bore.deepsquare.run
Direct TCP:          tcp://bore.deepsquare.run:60577
```

Use the HTTPS link for a secure connection.

Each time a user connects to the server, a new TTY session is created and will execute `sh`. If you want to allow only one session and exit after the first session use `--once`:

```yaml title="Workflow with once"
steps:
  - name: 'interactive'
    run:
      container:
        image: curlimages/curl:latest
      env:
        # Set the terminal type. xterm is supported by tty2web
        - key: TERM
          value: xterm
      command: |
        curl -fsSL -o "$STORAGE_PATH/tty2web" https://github.com/kost/tty2web/releases/download/v3.0.3/tty2web_linux_amd64
        chmod +x "$STORAGE_PATH/tty2web"
        "$STORAGE_PATH/tty2web" --permit-write --once --port 8080 --credential admin:password sh
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 8080
```

If your container is not able to run `curl` or `wget`, you can do it in two-steps:

```yaml title="Two steps workflow"
steps:
  - name: 'setup-tty2web'
    run:
      container:
        image: curlimages/curl:latest
      command: |
        curl -fsSL -o "$STORAGE_PATH/tty2web" https://github.com/kost/tty2web/releases/download/v3.0.3/tty2web_linux_amd64
        chmod +x "$STORAGE_PATH/tty2web"
  - name: 'interactive'
    run:
      container:
        image: busybox:latest
      env:
        # Set the terminal type. xterm is supported by tty2web
        - key: TERM
          value: xterm
      command: |
        "$STORAGE_PATH/tty2web" --permit-write --port 8080 --credential admin:password sh
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 8080
```

If you want to allow multiple web session, but want to avoid losing a terminal session, it is possible to use `screen`:

```yaml title="Workflow with screen"
steps:
  - name: 'setup-tty2web'
    run:
      container:
        image: curlimages/curl:latest
      command: |
        curl -fsSL -o "$STORAGE_PATH/tty2web" https://github.com/kost/tty2web/releases/download/v3.0.3/tty2web_linux_amd64
        chmod +x "$STORAGE_PATH/tty2web"
  - name: 'interactive'
    run:
      container:
        image: alpine:latest
      env:
        # Set the terminal type. xterm is supported by tty2web
        - key: TERM
          value: xterm
      command: |
        apk add screen
        screen -dmS my-session-name sh
        "$STORAGE_PATH/tty2web" --permit-write --port 8080 --credential admin:password screen -x my-session-name
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 8080
```

You can also use the module:

```yaml title="Workflow with screen"
steps:
  - use:
      source: github.com/deepsquare-io/workflow-modules/tty2web
  - name: 'interactive'
    run:
      container:
        image: alpine:latest
      env:
        # Set the terminal type. xterm is supported by tty2web
        - key: TERM
          value: xterm
      command: |
        apk add screen
        screen -dmS my-session-name sh
        "$STORAGE_PATH/tty2web" --permit-write --port 8080 --credential admin:password screen -x my-session-name
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 8080
```

## SSH server

It is possible to run an SSH server and forward its port using `bore`.

```yaml title="Workflow"
steps:
  - name: 'interactive'
    run:
      container:
        image: alpine:latest
      command: |
        apk add openssh

        # Pass public key
        mkdir -p "$HOME/.ssh"
        echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPd+X08wpIGwKZ0FsJu1nkR3o1CzlXF3OkgQd/WYB2fX" > "$HOME/.ssh/authorized_keys"
        chmod 600 "$HOME/.ssh/authorized_keys"
        chmod 700 "$HOME/.ssh"

        # Generate new host keys
        ssh-keygen -A

        # Print username, so we know how to login.
        echo $USER

        # Start server in foreground (use port 2200 since we are not privileged)
        /usr/sbin/sshd -D -p 2200
      network: slirp4netns
      customNetworkInterfaces:
        - bore:
            address: bore.deepsquare.run
            port: 2200
            targetPort: 2200
```

The logs of the job will output a "bore.deepsquare.run" URL which contains the URL of the SSH server:

```shell title="Example of output"
0x75761b17c3088ce5cd8e02575c6daa438ffa6e12
Generated HTTP URL:  http://ac21705f.bore.deepsquare.run
Generated HTTPS URL: https://ac21705f.bore.deepsquare.run
Direct TCP:          tcp://bore.deepsquare.run:60577
```

We need to use `bore.deepsquare.run:60577`:

```shell title="Client"
ssh -p 60577 0x75761b17c3088ce5cd8e02575c6daa438ffa6e12@bore.deepsquare.run
```
