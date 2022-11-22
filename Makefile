test: 
	go test -v -vet=off ./...
    
init:
	terraform init

apply: 
	terraform apply

plugin:
	@mkdir -p ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64 

build: plugin
	@go build -o ~/.terraform.d/plugins/example.com/local/plugin/1.0.0/linux_amd64 
