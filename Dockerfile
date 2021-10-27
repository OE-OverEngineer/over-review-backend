FROM --platform=linux/amd64 golang:1.17 as base

# Container for develop
FROM base as dev
ENV GIN_MODE=debug

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]

# Build for Production
FROM base as builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./build/server ./cmd/main.go


# Container for production
FROM alpine:3.14.2 as production
ENV GIN_MODE=release

COPY --from=builder /src/build/server .

ENTRYPOINT [ "./server" ]