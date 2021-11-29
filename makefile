compile:
	@docker run --rm -v $(CURDIR):/usr/src/myapp -w /usr/src/myapp golang:1.17 go build -v

run:
	@./tuimusement