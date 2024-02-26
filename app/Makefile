.PHONY: build
build:
	go mod download
ifeq ($(OS),Windows_NT)
	go build -buildvcs=false -o bin/hello.exe .
else
	CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o bin/hello .
endif

.PHONY: run
run: build
ifeq ($(OS),Windows_NT)
	bin/hello.exe --config configuration.json
else
	bin/hello --config configuration.json
endif

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	rd /S /Q bin
else
	rm -f bin
endif
