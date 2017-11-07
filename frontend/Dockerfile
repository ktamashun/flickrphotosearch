FROM scratch

ARG BUILD_DIR
ARG BINARY_NAME

COPY $BUILD_DIR/$BINARY_NAME /app
COPY ./templates/index.html /templates/index.html

EXPOSE 8080 8080

CMD ["/app"]