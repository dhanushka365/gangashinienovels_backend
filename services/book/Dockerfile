FROM alpine:3.15
COPY ./.bin/book /book
EXPOSE 8080
ENV MICRO_REGISTRY=kubernetes
ENTRYPOINT ["./book"]