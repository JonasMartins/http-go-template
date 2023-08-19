## small image
FROM alpine:latest

ARG BINARY_PROD=_main

WORKDIR /app

RUN apk add uuidgen

COPY src/services/main/configs/base-docker-prod.yaml /app/project/src/services/main/configs/base.yaml
COPY src/services/main/internal/repository/postgres/migrations /app/project/src/services/main/internal/repository/postgres/migrations
COPY out/$BINARY_PROD /app

CMD ["/app/_main"]
# ENTRYPOINT ["tail", "-f", "/dev/null"]


