FROM golang:1.18-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN GOPROXY="https://goproxy.cn,direct" go mod download
COPY . .
RUN go build -o ./bin/ ./...

FROM alpine
WORKDIR /app
COPY --from=build /src/bin /app
COPY configs configs
COPY designer_data designer_data
CMD ["./server", "-conf", "./configs"]