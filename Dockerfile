FROM golang:1.22.6
LABEL authors="KORIE-BRUH"

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk memanfaatkan cache layer
COPY go.mod go.sum ./
RUN go mod download

# Kemudian copy seluruh source code
COPY . .

# Install migrate tool
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["sh", "-c", "migrate -database 'mysql://root:korie123@tcp(mysql-db:3306)/krs_management' -path db/migrations up && ./main"]