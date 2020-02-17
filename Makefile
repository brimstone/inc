.PHONY: build
build:
	goreleaser release --skip-publish --rm-dist --skip-validate --skip-sign

.PHONY: clean
clean:
	rm -rf dist
