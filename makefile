build:
	ftmpl jsonconv
	ftmpl conv
	go run gojson2models.go -package github.com/tkrajina/gojson2models/example -language java -target=example/example.java example/example.go
	go run gojson2models.go -package github.com/tkrajina/gojson2models/example -language typescript -target=example/example.ts example/example.go
install-tools:
	go get -u github.com/tkrajina/ftmpl
fmt:
	gofmt -w .
	goimports -w .
test:
	go test ./...
