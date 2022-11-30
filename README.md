# Terraform provider plugin

This folder encapsulates the Terraform Provider that issues API calls to the [pkid](https://github.com/rawdaGastan/pkid).

## Running the example

To run the Terraform Provider locally there are a few steps to complete:

Step 1: Make sure your pkid server is [running](https://github.com/rawdaGastan/pkid#how-to-run-locally)

Step 2: Build the source code locally and move into the local terraform plugin folder:

```bash
mkdir -p ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64
go build -o ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64 
```

Or:

```bash
make build
```

> Note: The plugin folder may need to be created.

Step 3: Initialize Terraform:

```bash
terraform init
```

### Create

- Run an apply via Terraform:

```bash
terraform apply
```

The output generated should look similar to the following:

```bash
Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

Outputs:

vm_1 = {
  "encrypt" = true
  "id" = "pkid_key"
  "key" = "key"
  "project" = "pkid"
  "timeouts" = null /* object */
  "value" = "value"
}

vm_2 = {
  "encrypt" = true
  "id" = "pkid_key"
  "key" = "key"
  "project" = "pkid"
  "timeouts" = null /* object */
  "value" = "value"
}
```

### Update

- make a change in the resource
- Run an apply via Terraform:

```bash
terraform apply
```

The output generated should look similar to the following:

```bash
Apply complete! Resources: 0 added, 2 changed, 0 destroyed.

Outputs:

vm_1 = {
  "encrypt" = false
  "id" = "pkid_key"
  "key" = "key"
  "project" = "pkid"
  "timeouts" = null /* object */
  "value" = "value"
}

vm_2 = {
  "encrypt" = false
  "id" = "pkid_key"
  "key" = "key"
  "project" = "pkid"
  "timeouts" = null /* object */
  "value" = "value"
}
```

### Delete

- delete via terraform

```bash
terraform destroy -target=plugin_pkid_key_query.vm_1
```

The output generated should look similar to the following:

```bash
Destroy complete! Resources: 1 destroyed.
```

## Test

```bash
make test
```
