.PHONY: r
r:
	# cp config-example.yml config.yml
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