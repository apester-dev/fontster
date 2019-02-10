FROM golang AS builder
WORKDIR /src
COPY . /src
RUN go build

FROM alpine:latest
COPY --from=builder /src/fontster /fontster
USER nobody
CMD ["/fontster"]
