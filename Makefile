# These will be provided to the target
VERSION := 0.0.1
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-tags release -ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
ARCH ?= `go env GOHOSTARCH`
GOOS ?= `go env GOOS`

.godeps:
	dep ensure
	touch .godeps

all: agent master rendezvous

package: clean all .godeps
	cd bin && tar -zcvf reverse-shell-$(VERSION)-$(GOOS)-$(ARCH).tar.gz agent master rendezvous

agent: build_dir .godeps
	cd agents/go && go build $(LDFLAGS) -o ../../bin/agent

master: build_dir .godeps
	cd master && go build $(LDFLAGS) -o ../bin/master

rendezvous: build_dir .godeps
	cd rendezvous && go build $(LDFLAGS) -o ../bin/rendezvous

build_dir:
	mkdir -p bin

clean:
	rm -rf bin/*
