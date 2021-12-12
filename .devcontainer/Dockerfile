FROM golang:1.17-bullseye
RUN go install golang.org/dl/gotip@latest && /go/bin/gotip download
ENV PATH=/root/sdk/gotip/bin:$PATH
