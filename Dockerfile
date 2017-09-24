FROM golang:1.9-alpine3.6 AS build-env
ADD . /go/src/github.com/supinf/apis-on-gae/
RUN cd /go/src/github.com/supinf/apis-on-gae \
    && go build api/cmd/demo-apis-server/main.go

FROM alpine:3.6
RUN apk add --no-cache ca-certificates

COPY --from=build-env /go/src/github.com/supinf/apis-on-gae/main /usr/local/bin/api-server
ENV APP_VERSION=1.0.0

EXPOSE 8080
ENTRYPOINT ["api-server", "--host", "0.0.0.0", "--port", "8080"]
