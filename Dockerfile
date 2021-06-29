FROM golang as build

COPY . /go/src
WORKDIR /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o semver .


FROM node:alpine
RUN apk --no-cache add git ca-certificates
WORKDIR /root/
COPY --from=build /go/src/semver /usr/bin/semver
CMD ["semver"]
