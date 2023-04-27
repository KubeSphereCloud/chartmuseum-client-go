# chartmuseum-client-go
go client for chartmuseum

- Download swagger.json from https://github.com/KubeSphereCloud/chartmuseum/tree/main/cmd/chartmuseum/docs and saved it to ./openapi/swagger.json
- Generate go client:
```bash
rm -rf client models && swagger generate client -f ./openapi/swagger.json --skip-validation && go mod tidy
```
