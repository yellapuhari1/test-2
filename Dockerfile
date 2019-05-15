FROM alpine:3.9.4 AS build

RUN apk add --no-cache \
    && apk add --update git go libc-dev

COPY main.go /home/go/main.go

WORKDIR /home/go

RUN go get -d ./... && go build -a -o test-2-go-app && ls -l

FROM alpine:3.9.4

EXPOSE 8080

RUN addgroup -g 1000 go \
    && adduser -u 1000 -G go -s /bin/sh -D go \
    && apk add --no-cache \
    && apk add --update git

WORKDIR /home/go

COPY --from=build /home/go/test-2-go-app test-2-go-app

RUN chown -R go:go /home/go

USER go

CMD ["/bin/sh", "-c", "/home/go/test-2-go-app"]



