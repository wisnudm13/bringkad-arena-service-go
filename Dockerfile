FROM golang:1.24.0-alpine AS builder

# set work directory
WORKDIR /app

# copy the project
COPY . .

# copy go.mod and go.sum, then download dependencies
# COPY go.mod go.sum ./
# RUN go mod download


# build
RUN go build -o main main.go

# Run 
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

# expose port
EXPOSE 8080

# run the app
CMD ["/app/main"]