.PHONY: prepare

prepare:
	@echo "Installing golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Install Husky"
	@go install github.com/go-courier/husky/cmd/husky@latest && husky init
