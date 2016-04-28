build:
	ftmpl jsonconv
	ftmpl conv
	go run gojson2models.go -package github.com/tkrajina/gojson2models/example -language java -target=example/example.java example/example.go
	go run gojson2models.go -package github.com/tkrajina/gojson2models/example -language typescript -target=example/example.ts example/example.go
fmt:
	gofmt -w .
	goimports -w .
test:
	go test ./...
	go run example/example.go
	tsc browser_test/example_output.ts
