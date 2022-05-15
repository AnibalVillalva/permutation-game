FROM golang:1.14.7 as build-env

RUN mkdir /go
RUN mkdir /go/permutation-game
WORKDIR /go/permutation-game
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build the binary
WORKDIR /go/permutation-game/cmd/api
RUN go build -o /go/bin/permutation-game
FROM scratch
COPY --from=build-env /go/bin/permutation-game /go/bin/permutation-game
ENTRYPOINT ["/go/bin/permutation-game"]
