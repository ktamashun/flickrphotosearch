FROM scratch

ARG BUILD_DIR
ARG BINARY_NAME

COPY $BUILD_DIR/$BINARY_NAME /app

EXPOSE 18000 18000

CMD ["/app"]