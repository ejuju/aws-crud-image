test:
	go test -timeout 2m -v -race ./...

build:
	podman build -t ejuju/crud-aws -f ./containerfile .

run:
	PORT=8080 podman run --rm -p 8080:8080 ejuju/crud-aws