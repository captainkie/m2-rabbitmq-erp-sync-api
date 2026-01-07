.PHONY: changelog test test-coverage test-coverage-html

# Install git-chglog if not exists
install-chglog:
	@which git-chglog >/dev/null || (echo "Installing git-chglog..." && go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest)

# Generate changelog
changelog: install-chglog
	@echo "Generating changelog..."
	@git-chglog --output CHANGELOG.md

# Generate changelog for next version
changelog-next: install-chglog
	@echo "Generating changelog for next version..."
	@git-chglog --next-tag $(shell git-chglog --next-tag) --output CHANGELOG.md

# Generate changelog for specific version
changelog-version: install-chglog
	@echo "Generating changelog for version $(version)..."
	@git-chglog --output CHANGELOG.md $(version)

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p coverage
	@go test -v -coverprofile=coverage/coverage.out -covermode=atomic ./...
	@go tool cover -func=coverage/coverage.out

# Generate HTML coverage report
test-coverage-html: test-coverage
	@echo "Generating HTML coverage report..."
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report generated at coverage/coverage.html"

# Help command
help:
	@echo "Available commands:"
	@echo "  make changelog        - Generate full changelog"
	@echo "  make changelog-next   - Generate changelog for next version"
	@echo "  make changelog-version version=x.x.x - Generate changelog for specific version"
	@echo "  make test            - Run all tests"
	@echo "  make test-coverage   - Run tests with coverage report"
	@echo "  make test-coverage-html - Generate HTML coverage report"
	@echo "  make help            - Show this help message" 