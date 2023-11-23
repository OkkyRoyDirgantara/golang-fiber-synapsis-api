# Gunakan golang image sebagai base image
FROM golang:latest AS build

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum ke dalam container
COPY go.mod .
COPY go.sum .

# Install dependencies menggunakan go get
RUN go mod download

# Copy source code aplikasi ke dalam container
COPY . .



# Build aplikasi
RUN go build -o main .

# Stage kedua, gunakan alpine image yang lebih ringan
FROM alpine:latest

# Set working directory di dalam container
WORKDIR /app

# Copy template engine file aplikasi
COPY src/views/ ./src/views/
COPY ./.env.example ./.env 

# Copy executable dari stage pertama ke dalam container
COPY --from=build /app/main .

# Install dependensi untuk MySQL
RUN apk add --no-cache libc6-compat

# Expose port yang digunakan oleh aplikasi
EXPOSE 3000

# Environment variables untuk koneksi MySQL
ENV DB_USERNAME=admin
ENV DB_PASSWORD=admin
ENV DB_HOST=db
ENV DB_PORT=3306
ENV DB_NAME=synapsis


# Command untuk menjalankan aplikasi
CMD ["./main"]
