.PHONY: clean
.DEFAULT_GOAL:=help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

install: ## 安装依赖
	go mod tidy 

start: ## 启动 air 热更新
	air -c .air.conf

lint: ## 格式化
	find . -name "*.go"|gawk -n 1|xargs goimports -w

clean: ## 执行清理
	rm -rf ./output/*

build: lint clean ## 打包 先格式化
	GIN_MODE=release go build -o ./output/main .

docs: lint  ## swag 文档生成
	swag init
