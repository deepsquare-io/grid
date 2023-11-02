# Joining as a Validator

We appreciate your interest in becoming a validator for the DeepSquare blockchain! At the moment, we've filled all available validator spots. However, you can still prepare your node to join when slots become available. If you wish to express your interest in becoming a validator, please complete [this form](https://share-eu1.hsforms.com/1tczPxyxMSGOS6a8ojGX-gwev6gi) and join us on our [Discord server](http://discord.gg/RyEARZnsMZ) for updates and community interaction. Below is a guide on how you can set up for future opportunities:

## Setting up for the Subnet

To set up your node for joining the subnet and validating the Mondrian Testnet in the future, you'll need the following information:

- Subnet: `8dRakstCMfHV8CXRhdtq9Wbxo75535Pfran1yDX5x4TYJq22A`
- Blockchain: `23q7DGje3AFbLKCgXFWcW6eo9zsB166mfknGHt5dySefGtJboZ`
- VM ID: `mDV28Yo1kHR1XAXo29LJsVh38vyKUdsvcdAZXYakdQd3LMwBY`

Here are the steps you need to follow to get your node set up:

### Set up the Chain Config

1. Create a config file for the blockchain (assuming the avalanchego data resides in `~/.avalanchego/configs`):

   ```bash
   AVAPATH_CFG=~/.avalanchego/configs
   mkdir -p ${AVAPATH_CFG}/chains/23q7DGje3AFbLKCgXFWcW6eo9zsB166mfknGHt5dySefGtJboZ
   touch ${AVAPATH_CFG}/chains/23q7DGje3AFbLKCgXFWcW6eo9zsB166mfknGHt5dySefGtJboZ/config.json
   ```

2. Add the following content to the file:

   ```json
   {
     "allow-missing-tries": true
   }
   ```

### Set up the VM

The current virtual machine used is `subnet-EVM`. To get the binary, visit our Github repository and follow these steps:

1.  Download the VM from [here](https://github.com/deepsquare-io/testnet-subnet-vm/blob/master/vms/v22/mDV28Yo1kHR1XAXo29LJsVh38vyKUdsvcdAZXYakdQd3LMwBY).

2.  Put the VM into your avalanchego plugins directory (e.g., `~/.avalanchego/plugins`):

3.  Make executable:

    ```shell
    chmod +x ~/.avalanchego/plugins/mDV28Yo1kHR1XAXo29LJsVh38vyKUdsvcdAZXYakdQd3LMwBY
    ```

Once you've done all these, your node will be ready to join the subnet and start validating when slots become available in the future.

## Start validating

You are now ready to launch the node and start validating the subnet.

Assuming you followed the step before and your node has been whitelisted you can start your avalanchego process with the following options:

```shell
--network-id=fuji
--track-subnets=8dRakstCMfHV8CXRhdtq9Wbxo75535Pfran1yDX5x4TYJq22A
```
