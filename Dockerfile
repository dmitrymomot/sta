FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
WORKDIR /
COPY ./sta /
ENV APP_PORT=8080
EXPOSE $APP_PORT
ENTRYPOINT ["/sta"]
