# Multistage build - phase 1 - build application code
FROM golang:1.22 as builder
WORKDIR /hello-k8s
COPY go.mod .
RUN go mod download
ADD . .
RUN make build 

# Multistage build - phase 2 - create final container
FROM scratch as release
COPY --from=builder /hello-k8s/bin/hello /go/bin/hello
COPY --from=builder /hello-k8s/configuration.json /go/bin/configuration.json