.PHONY: r
r:
	go run main.go --config ./config.yml

.PHONY: b
b:
	go build

.PHONY: c
c:
	go clean -modcache

.PHONY: l
l:
	golangci-lint run

.PHONY: t
t:
	go mod tidy

.PHONY: tag
tag:
	@if [ -z "$(TAG)" ]; then \
		echo "Error: TAG variable is not set. Use 'make tag TAG=v1.0.0'"; \
		exit 1; \
	fi
	git tag $(TAG)
	git push origin $(TAG)