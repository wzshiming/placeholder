PACKAGES := .

.PHONY: run generate

run:
	go run ./cmd/placeholder

generate:
	mkdir -p route && gen route --openapi -o ./route/route_gen.go ${PACKAGES}