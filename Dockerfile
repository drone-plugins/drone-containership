# Docker image for the Drone ContainerShip plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-containership
#     make deps build docker

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-containership /bin/
ENTRYPOINT ["/bin/drone-containership"]
