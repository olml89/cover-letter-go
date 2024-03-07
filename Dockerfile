# Compilation stage
ARG GOLANG_VERSION
FROM golang:${GOLANG_VERSION}-alpine AS build-env

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./src /go/src
COPY .env /go/.env
WORKDIR /go/src

# Compile the application with the optimizations turned off
# This is important for the debugger to correctly work with the binary
RUN go build -gcflags "all=-N -l" -o /go/bin/app

# Final stage
FROM alpine

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /go/bin/app /
COPY --from=build-env /go/.env /

RUN chmod +x /dlv
RUN chmod +x /app
