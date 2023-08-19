## small image
FROM alpine:latest

ARG BINARY_DEV=_main

WORKDIR /app

RUN apk add uuidgen

COPY src/services/main/configs/base-docker-dev.yaml /app/project/src/services/main/configs/base.yaml
COPY src/services/main/internal/repository/postgres/migrations /app/project/src/services/main/internal/repository/postgres/migrations
COPY out/$BINARY_DEV /app
COPY wait-for.sh .
COPY start.sh .

CMD ["/app/_main"]
ENTRYPOINT [ "/app/start.sh" ]
# ENTRYPOINT ["tail", "-f", "/dev/null"]

