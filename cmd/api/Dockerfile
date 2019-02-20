FROM golang:alpine as base
RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc \
        wget \
        git \
    && apk add \
        bash

# Build the app
FROM base as builder
ADD . ./app
WORKDIR ./app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build ./cmd/api/

# Final image
FROM alpine:3.9
COPY --from=builder /go/app/api .
EXPOSE 1323/tcp
ENTRYPOINT ["api"]