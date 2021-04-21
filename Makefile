.PHONY: help test cover

help:
	@echo
	@echo "  help："
	@echo "  - make help"
	@echo "  - make test"
	@echo "  - make cover"
	@echo

test:
	@go test -v -covermode=count -coverprofile=coverage.txt $(shell go list ./...);

cover: test
	go tool cover -html=coverage.txt
	@rm coverage.txt

