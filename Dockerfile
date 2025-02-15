FROM golang:1.23 as build

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /go/bin/tagsky .



FROM scratch

COPY --from=build /go/bin/tagsky /tagsky

CMD ["/tagsky"]
