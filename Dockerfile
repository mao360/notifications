FROM golang:latest
ENV GOPATH=/
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o notifications ./cmd/main.go
CMD ["./notifications"]