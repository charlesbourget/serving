.PHONY: build

build:
	rm -f bin/serving
	rm -rf static
	mkdir static
	cd web && yarn build && cp -r prod/ ../static
	go build -o bin/serving ./...
