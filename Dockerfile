FROM --platform=$BUILDPLATFORM golang:1.14 as build-env

# xx wraps go to automatically configure $GOOS, $GOARCH, and $GOARM
# based on TARGETPLATFORM provided by Docker.
COPY --from=tonistiigi/xx:golang-1.0.0 / /

ARG APP_FOLDER

ADD . ${APP_FOLDER}
WORKDIR ${APP_FOLDER}

# Compile independent executable using go wrapper from xx:golang
ARG TARGETPLATFORM
RUN CGO_ENABLED=0 go build -a -o /bin/main .

FROM --platform=$TARGETPLATFORM alpine

RUN apk --update add git && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY --from=build-env /bin/main /app/main

EXPOSE 8080
ENTRYPOINT ["/app/main"]