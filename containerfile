FROM golang:1.18.4 as build-server
WORKDIR /usr/local/go/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main

FROM scratch
WORKDIR /
COPY --from=build-server /usr/local/go/app/main ./
EXPOSE 8080
ENTRYPOINT [ "./main" ]