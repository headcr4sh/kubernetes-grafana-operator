.POSIX:

SHELL = /bin/sh
CWD := `pwd -LP`
DIST_DIR := "$(CWD)/dist"

.PHONY: all
all: build

.PHONY: prepare
prepare:
	@echo "Creating output directory ($(DIST_DIR))..."
	@mkdir -p $(DIST_DIR)

.PHONY: build
build: prepare
	@echo "Creating executable for linux-amd64 systems..."
	@GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/kubernetes-grafana-operator-linux-amd64 ./main.go
