BINNAME = myip

all: vet fmt
	go build -ldflags="-s -w"

clean:
	rm -f ${BINNAME}

fmt:
	go fmt

vet:
	go vet

.PHONY: all clean fmt vet
