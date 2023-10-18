get-oapi:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

gen-oapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	oapi-codegen --config=oapi/cfg-server.yaml -include-tags="server" oapi/integration.yaml
	oapi-codegen --config=oapi/cfg-client.yaml -include-tags="client" --exclude-schemas=PlayMode oapi/integration.yaml

gen-ogen:
	go install -v github.com/ogen-go/ogen/cmd/ogen@latest
	go generate ./...
