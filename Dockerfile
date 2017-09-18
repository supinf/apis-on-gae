FROM golang:1.9-alpine3.6 AS build-env
RUN apk add --no-cache git
RUN go get github.com/supinf/apis-on-gae/api/cmd/demo-apis-server

FROM alpine:3.6
RUN apk add --no-cache ca-certificates

COPY --from=build-env /go/bin/demo-apis-server /usr/local/bin/api-server
ENV APP_VERSION=1.0.0

EXPOSE 8080
ENTRYPOINT ["api-server", "--host", "0.0.0.0", "--port", "8080"]
