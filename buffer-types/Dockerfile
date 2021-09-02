FROM registry.hub.docker.com/library/golang:alpine AS build
RUN apk --update add ca-certificates git

WORKDIR /src/
COPY . /src/
ARG BUILD_COMMIT
RUN go get -d ./...
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
