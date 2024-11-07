.PHONY: bin/matcher
bin/matcher:
	go build -o $@ ./cmd/matcher

.PHONY: lint
lint: bin/staticcheck
	go vet ./...
	bin/staticcheck ./...

.PHONY: test test-unit test-behave
test: test-unit test-behave

test-unit:
	go test ./...

test-behave: .venv/bin/activate
	source .venv/bin/activate && behave

.venv/bin/activate: features/requirements.txt
	python3 -m venv .venv
	source $@ && pip install --upgrade pip && pip install -r $<

bin/staticcheck:
	GOBIN=$(shell pwd)/bin go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: docs
docs: api/v1/spec.html
api/v1/spec.html: api/v1/openapi.yaml
  # npm i -g @redocly/cli@latest
	redocly build-docs $< -o $@
