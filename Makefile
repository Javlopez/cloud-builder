
test:
	go test ./...

run:
	go run main.go


test-with-coverage:
	go test ./... -v -coverprofile cover.out

coverage: 
	make test-with-coverage
	go tool cover -html=cover.out