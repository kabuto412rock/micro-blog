.PHONY: local
wire:
	wire gen ./...

local:
	make wire && go build -o bin/micro-blog ./cmd/micro-blog && ./bin/micro-blog