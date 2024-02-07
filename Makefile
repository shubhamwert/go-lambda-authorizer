build-goAuth:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap -tags lambda.norpc
	cp bootstrap "$(ARTIFACTS_DIR)/bootstrap"