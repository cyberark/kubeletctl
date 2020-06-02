FROM golang:1.13.1-alpine3.10 as builder

RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go get github.com/mitchellh/gox
RUN gox -ldflags "-s -w" -osarch linux/386 -output "kubeletctl"

FROM alpine:latest
COPY --from=builder /src/kubeletctl /app/
WORKDIR /app

# Create a non-root user and group
# -H Don't create home directory
# -D Don't assign a password
# -g GECOS field
# -G Group
RUN set -ex \
  && addgroup kubeletctl \
  && adduser \
    -H \
    -D \
    -g ',,,,' \
    -G kubeletctl \
    kubeletctl \
  && > /var/log/faillog \
  && > /var/log/lastlog

USER kubeletctl

ENTRYPOINT ["/app/kubeletctl"]
