.PHONY: start-server
start-server:
	go run -v ./cmd/server/

.PHONY: start-client
start-client:
	go run -v ./cmd/client/