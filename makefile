build:
	ftmpl jsonconv
	ftmpl conv
	go run gojson2models/main.go -package github.com/tkrajina/typescriptify-golang-structs/example -language java -target=example/example.java example/example.go
	go run gojson2models/main.go -package github.com/tkrajina/typescriptify-golang-structs/example -language typescript -target=example/example.ts example/example.go
test:
	go test ./...
	go run example/example.go
	tsc browser_test/example_output.ts
