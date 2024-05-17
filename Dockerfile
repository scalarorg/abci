FROM alpine:3.19

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file

VOLUME /scalaris
WORKDIR /scalaris
EXPOSE 26656 26657 26660
ENTRYPOINT ["/usr/bin/start.sh"]
CMD ["node"]
STOPSIGNAL SIGTERM

COPY start.sh /usr/bin/start.sh

