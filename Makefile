GOFILES = $(shell find . -type f -name '*.go' -not -path "./.git/*")

.PHONY: build

build:
	CGO_ENABLED=0 go build -o build/autoscaling-demo github.com/laszlocph/autoscaling-demo/cmd
