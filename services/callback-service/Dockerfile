FROM golang:1.20-alpine AS mapper-service

RUN  apk add make git gcc musl-dev
# Set the Current Working Directory inside the container

WORKDIR /app/mapper-service

COPY . .
RUN go mod download
RUN make deps
RUN GOFLAGS=" -tags=musl" make all
EXPOSE 8765
CMD ["/app/mapper-service/mapper-service", "--port", "8766", "--host", "0.0.0.0"]
