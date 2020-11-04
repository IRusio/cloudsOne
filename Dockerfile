FROM golang:1.12.0-alpine3.9 as build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git
RUN apk add --no-cache tzdata
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=build /app/main .
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ=Europe/Warsaw
ENV ZONEINFO=/zoneinfo.zip
CMD ["./main"]