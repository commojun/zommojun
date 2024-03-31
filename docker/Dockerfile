FROM golang:1.22 AS builder

WORKDIR /
COPY . .
RUN go mod download
RUN go build -o zommojun ./main.go

FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y locales ca-certificates && \
    rm -rf /var/lib/apt/lists/*  && \
    localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.UTF-8

COPY --from=builder /zommojun /usr/local/bin/zommojun

RUN useradd -r app
USER app
WORKDIR /home/app/

CMD ["zommojun"]
