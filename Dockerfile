# stage 1
FROM golang:1.21.3-alpine3.18 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

# stage 2 - use the built binary from stage 1 as a dependency
FROM alpine:3.18 AS prod

# set the working directory
WORKDIR /app

# Copy only the necessary files from the build image

COPY --from=build /api /api 


# Set entry point
CMD ["/api"]
