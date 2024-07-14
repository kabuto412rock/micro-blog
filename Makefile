# 定義變數
EXECUTABLE := micro-blog
SOURCES ?= $(shell find . -name "*.go" -type f)
GO ?= go
BIN_DIR := bin
WIRE := wire
SWAG := swag

# 標籤和鏈接標誌變數
TAGS ?= 
EXTLDFLAGS ?= 
LDFLAGS ?= -s -w

# 生成執行檔
build: 
	$(GO) build -o $(BIN_DIR)/$(EXECUTABLE) ./cmd/$(EXECUTABLE)

install:
	${GO} install github.com/google/wire/cmd/wire@latest
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

# wire 目標
wire: 
	$(WIRE) ./...

# swag 目標
swag:
	$(SWAG) init

# 執行前的目標 (只運行 wire 和 swag)
prebuild: wire swag

# 運行目標，編譯後直接運行，不生成執行檔
run: prebuild
	$(GO) run -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS) $(LDFLAGS)' .

# install air command
.PHONY: air
air:
	@hash air > /dev/null 2>&1; \
	if [ $$? -ne 0 ]; then \
		$(GO) install github.com/air-verse/air@latest; \
	fi

# run air
.PHONY: dev
dev: air
	air --build.cmd "make build" --build.bin "${BIN_DIR}/${EXECUTABLE}"

# air --build.cmd "make build" --build.bin "bin/api-service" -build.exclude_dir=docs -build.exclude_file=wire_gen.go

# 清理執行檔
clean:
	rm -rf $(BIN_DIR)/$(EXECUTABLE)

# .PHONY 目標
.PHONY: build wire swag run clean 



# .PHONY: local
# wire:
# 	wire gen ./...

# local:
# 	make wire && go build -o bin/micro-blog ./cmd/micro-blog && ./bin/micro-blog