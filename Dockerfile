FROM golang:1.13-buster AS builder

WORKDIR /build
COPY . .
RUN apt update && apt install -y build-essential
RUN make build

FROM ubuntu:18.04
WORKDIR /app

RUN apt update && apt install -y openjdk-11-jdk libjna-java
COPY --from=builder /build/bin/* /app/
COPY --from=builder /build/srv/filebot.deb .
RUN dpkg -i ./filebot.deb

CMD ["./goautoplex"]
