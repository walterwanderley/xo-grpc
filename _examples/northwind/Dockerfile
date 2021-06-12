FROM golang:1.16-alpine

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go install github.com/cespare/reflex@latest
RUN go mod download -x

COPY configs/reflex.conf /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]