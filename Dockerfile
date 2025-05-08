# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app

# Copy all modules
COPY ./src/chatchk chatchk
COPY ./src/ingest ingest
COPY ./src/knowledge knowledge
COPY ./src/prompts prompts
COPY ./src/admin admin
COPY ./src/open_webui open_webui
COPY ./src/ollama ollama
COPY ./src/utils utils

# Build the binary
RUN cd chatchk && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/chatchk

# Stage 2: Create minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/chatchk .

# Copy other files
COPY ./knowledge_files /app/knowledge_files/

# Expose port (adjust if your app uses a specific port)
EXPOSE 8080

# Environment variables
ENV OLLAMA_KNOWLEDGE_FILE=./knowledge_files/customer_support_log.txt
# Follwoing is temporary: mv to k8s manifest/helm when ready
ENV OLLAMA_IP=<IP Address>
ENV OLLAMA_PORT=30111
ENV OLLAMA_API_KEY=<API Key>

# Run the binary
CMD ["./chatchk"]
