ARG BINARY=/bin/linkfire-webapp

# Build stage
FROM golang:1.17-alpine as build-stage
ARG BINARY

ENV CGO_ENABLED=0

WORKDIR /go/src

COPY . ./
RUN go mod download
RUN go build -o "${BINARY}" .

FROM alpine:latest
ARG BINARY

COPY --from=build-stage "${BINARY}" "${BINARY}"

ENV BINARY=${BINARY}
ENTRYPOINT ["/bin/linkfire-webapp"]
