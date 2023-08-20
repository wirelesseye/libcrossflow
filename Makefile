.PHONY: all build clean run build_web rebuild_web build_exec rebuild_exec rebuild_all

all: web/node_modules build_web build_exec out/res

web/node_modules:
	@cd web && npm install

web/dist: web/node_modules
	cd web && npm run build

out/crossflow:
	go build -o out/crossflow

out/res: web/dist
	cp -r web/dist out/res/

clean:
	rm -rf out/*
	rm -rf web/dist

build_web: out/res
 
rebuild_web:
	rm -rf web/dist
	rm -rf out/res
	@$(MAKE) build_web

build_exec: out/crossflow

rebuild_exec:
	rm out/crossflow
	@$(MAKE) build_exec

rebuild_all:
	@$(MAKE) clean
	@$(MAKE) all

run: web/dist
	@RES_PATH=web/dist go run .