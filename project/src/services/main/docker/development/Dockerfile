## small image
FROM alpine:latest

ARG BINARY_DEV=_main

WORKDIR /app

COPY src/services/main/configs/base-docker-dev.yaml /app/project/src/services/main/configs/base.yaml
COPY src/services/main/internal/repository/postgres/migrations /app/project/src/services/main/internal/repository/postgres/migrations
COPY out/$BINARY_DEV /app

CMD ["/app/_main"]
# ENTRYPOINT ["tail", "-f", "/dev/null"]

