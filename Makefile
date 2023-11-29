.PHONY: all build clean run build_web rebuild_web build_bin rebuild_bin rebuild_all

all: build_web build_bin out/res

web/node_modules:
	@cd web && npm install

web/dist: web/node_modules
	cd web && npm run build

out/crossflow:
	go build -o out/crossflow

out/res: web/dist
	mkdir -p out/res
	cp -r web/dist out/res/

clean:
	rm -rf out/*
	rm -rf web/dist

build_web: out/res
 
rebuild_web:
	rm -rf web/dist
	rm -rf out/res
	@$(MAKE) build_web

build_bin: out/crossflow

rebuild_bin:
	rm out/crossflow
	@$(MAKE) build_exec

rebuild_all:
	@$(MAKE) clean
	@$(MAKE) all

run: web/dist
	@APP_ENV=dev go run .