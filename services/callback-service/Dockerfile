FROM golang:1.20-alpine AS callback-service

RUN  apk add make git gcc musl-dev
# Set the Current Working Directory inside the container

WORKDIR /app/callback-service

COPY . .
RUN go mod download
RUN make deps
RUN GOFLAGS=" -tags=musl" make all
EXPOSE 8765
CMD ["/app/callback-service/callback-service", "--port", "8765"]
