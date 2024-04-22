run:
	go run main.go
runalltests:
	go test ./...
runcoverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out