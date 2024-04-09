# Check for any possible errors
# Detect any possible shadowed variables
vet:
	go fmt ./...
	go vet ./...
	shadow ./...


# Add missing and remove unsed modules from go.mod
tidy:
	go fmt ./...
	go mod tidy

# Run the program
run:
	go run main.go
