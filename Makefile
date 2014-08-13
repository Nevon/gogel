all: gogel

gogel: dom.go html.go main/main.go
	go build .

clean:
	rm -f gogel

run: gogel
	./gogel --html=example/html/test.html

fmt:
	go fmt .

install: gogel
	cp -f gogel /usr/local/bin/gogel

test: gogel
	go test .

coverage: gogel
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

install_deps:
	go get -u github.com/stretchr/testify/assert
	go get -u github.com/codegangsta/cli
