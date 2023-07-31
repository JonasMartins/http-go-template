## small image
FROM alpine:latest

ARG BINARY_DEV=_main

WORKDIR /app

COPY src/services/main/configs/base.yaml /app/project/src/services/main/configs/base.yaml
COPY out/$BINARY_DEV /app

# CMD ["/app/_main"]
ENTRYPOINT ["tail", "-f", "/dev/null"]


