# build stage
FROM golang:1.9 AS build
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN go-wrapper install

# final container
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
ENTRYPOINT ["/app"]
EXPOSE 3000
