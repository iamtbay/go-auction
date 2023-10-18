run: build 
	./go-auction
build:
	@go build -o go-auction ./cmd/*.go