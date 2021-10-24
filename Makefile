BINARY_NAME = gco

build:
	go build -v -o gco ./cmd

run: build
	./$(BINARY_NAME) 

install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/

test:
	go test -cover -parallel 5 -failfast  ./...

tidy:
	go mod tidy

lint:
	revive -formatter friendly -config revive.toml ./...

spell:
	misspell -error ./**

staticcheck:
	staticcheck ./...

gosec:
	gosec -tests ./... 

inspect: lint spell gosec staticcheck


# (build but with a smaller binary)
dist:
	go build -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
pack: dist
	upx ./$(BINARY_NAME)
