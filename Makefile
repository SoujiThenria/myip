BINNAME = myip

TIME = $(shell date)
USER = $(shell id -u -n)@$(shell hostname)
VERSIONPATH = github.com/SoujiThenria/myip/cmd.Version
VERSION = v1.1.2

all: vet fmt
	go build -ldflags="-s -w -X '${VERSIONPATH}=${VERSION} by ${USER} at ${TIME}'"

clean:
	rm -f ${BINNAME}

fmt:
	go fmt

vet:
	go vet

.PHONY: all clean fmt vet
