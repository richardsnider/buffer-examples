FROM registry.hub.docker.com/library/golang:alpine AS build-artifact-stage

WORKDIR /src/
COPY . /src/
ARG BUILD_COMMIT
RUN go get -d ./...
RUN export CURRENT_DATE_VERSION=$(date --utc +'%Y.%m.%d.%H.%M.%S') && \
CGO_ENABLED=0 \
go build \
-ldflags "-X github.com/richardsnider/nft-processor/main.BuildCommitLinkerFlag=$BUILD_COMMIT -X github.com/richardsnider/nft-processor/main.BuildDateVersionLinkerFlag=$CURRENT_DATE_VERSION" \
-o /bin/app

FROM scratch AS copy-artifact-stage
COPY --from=build-artifact-stage /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
