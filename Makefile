wintermute = ./src/

all: run


run:
	go run -race ${wintermute}


# build for specific OS target
build-%:
	GOOS=$* GOARCH=amd64 go build -o wintermute-$* ${wintermute}


build:
	go build -o wintermute ${wintermute}


# clean any generated files
clean:
	rm -rvf wintermute wintermute-*
