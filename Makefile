SRCS := $(shell find . -name \*.go)

all: vu

vu: $(SRCS)
	go build

test:
	go test ./...

clean:
	rm -f vu

.PHONY: clean test