FROM golang:1.24 as builder

ENV LANG en_US.utf8
WORKDIR /o11ybench

COPY . .
RUN make

FROM ubuntu:22.04 as base

WORKDIR /o11ybench
COPY --from=builder /o11ybench/bin/o11ybench /usr/local/bin/
ENTRYPOINT ["o11ybench"]
