# Stage 1: Frontend build (Vue3 + TypeScript)
FROM node:20 AS frontend-builder
WORKDIR /app

# Copy the frontend code into the container
COPY ./app /app

# Install dependencies and build the Vue app
RUN npm install && vite build

# Stage 2: Backend build (Go)
FROM golang:1.20 AS backend-builder
WORKDIR /go/src

# Set the Go proxy for module downloading
ENV GOPROXY https://goproxy.cn

# Copy the backend code into the container
COPY . /go

# Compile the Go backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo main.go

# Stage 3: Final production image
FROM nginx:alpine AS prod

# Copy the compiled Go binary from the backend build
COPY --from=backend-builder /go/main /usr/local/bin/main
# Copy configuration from the backend build
COPY --from=backend-builder /go/config.yaml /usr/local/bin/config.yaml
# Copy the log.txt file from the backend build
COPY --from=backend-builder /go/logs /usr/local/bin/logs
# Copy the static dir from the backend build
COPY --from=backend-builder /go/src/web/static /usr/local/bin/static

# Copy the frontend dist files to the static directory in Nginx
COPY --from=frontend-builder /app/dist /usr/share/nginx/html

# Copy the Nginx config (you can mount this config later)
COPY ./nginx.conf /etc/nginx/nginx.conf

# Expose necessary ports (e.g., 80 for Nginx, 8001 for the backend)
EXPOSE 80 8001

# Start Nginx and the Go backend
CMD ["sh", "-c", "nginx -g 'daemon off;' & /usr/local/bin/main"]
