test: 
	go test -v -vet=off ./...

docs: 
	go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	
init:
	terraform init

apply: 
	terraform apply

plugin: clean
	@mkdir -p ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64 

build: plugin
	@go build -o ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64 

clean: 
	rm ./.terraform -rf
	rm -f ./.terraform.lock.hcl
	rm -f terraform.tfstate
	rm -f terraform.tfstate.backup
