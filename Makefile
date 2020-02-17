.PHONY: build
build:
	goreleaser release --skip-publish --rm-dist --skip-validate --skip-sign

.PHONY: clean
clean:
	rm -rf dist
	rm -rf inc incd inc-client

.PHONY: test
test:
	go build -v .../cmd/inc-client
	./inc-client -h
	go build -v .../cmd/incd
	./incd -h
	go build -v .../cmd/inc
	./inc -h
