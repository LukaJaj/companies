FROM golang:1.18 as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o app .

FROM gcr.io/distroless/static-debian11
COPY --from=build /usr/src/app/app /

CMD ["/app"]
