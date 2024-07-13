.PHONY: local

local:
	go build -o bin/micro-blog ./cmd/micro-blog && ./bin/micro-blog