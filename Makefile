GO111MODULE=on
COVER_PROFILE=cover.out
GOLANGCI_LINT_VERSION=v1.15.0

up:
	docker-compose up -d

down:
	docker-compose down

qainstall:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)

test:
	go test -race -cover ./...

coverage:
	go test ./... -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

lint:
	golangci-lint run

bench:
	GOMAXPROCS=4 go test -bench=. -benchtime=10s -benchmem ./...
