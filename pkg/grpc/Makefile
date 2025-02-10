# Генерация proto файлов осуществляется с помощью этого скрипта, вызвать можно командой make url_gen
URL_SERVICE_PATH="proto/url_shortener/url_shortener.proto"
URL_SERVICE_OUT_PATH="./gen/"

.PHONY: all
all: help

.PHONY: help
help:
	@echo "Available commands:"
	@echo "url_gen  -- url shortener go files generation by proto"

.PHONY: url_gen
url_gen:
	@echo "Run url shortener go files generation..."
	protoc -I proto $(URL_SERVICE_PATH) --go_out=$(URL_SERVICE_OUT_PATH) --go_opt=paths=source_relative --go-grpc_out=$(URL_SERVICE_OUT_PATH) --go-grpc_opt=paths=source_relative
	@echo "End url shortener go files generation"
