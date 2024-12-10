# Stage 1: Frontend build (Vue3 + TypeScript)
FROM node:20 AS frontend-builder
WORKDIR /app

# Copy only the necessary files for installing dependencies
COPY ./app/package*.json ./

# Install dependencies (this step will be cached if dependencies don't change)
RUN npm install

# Copy the rest of the frontend code
COPY ./app ./

# Build the Vue app using npx to avoid global installation of vite
RUN npx vite build

# Stage 2: Backend build (Go)
FROM golang:1.20 AS backend-builder
WORKDIR /go/work

# Set the Go proxy for module downloading
ENV GOPROXY=https://goproxy.cn,direct

# Copy only the go.mod and go.sum files to cache dependencies
COPY ./go.mod ./go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the backend code
COPY . ./

# Compile the Go backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

# Stage 3: Final production image
FROM nginx:alpine AS prod

# Copy the compiled Go binary from the backend build
COPY --from=backend-builder /go/work/main /usr/local/bin/main

# Copy the configuration file
COPY --from=backend-builder /go/work/config.yaml /usr/local/bin/config.yaml

# Copy the log directory
COPY --from=backend-builder /go/work/logs /usr/local/bin/logs

# Copy the static assets directory
COPY --from=backend-builder /go/work/src/web/static /usr/local/bin/static

# Copy the frontend dist files to the static directory in Nginx
COPY --from=frontend-builder /app/dist /usr/share/nginx/html

# Copy the custom Nginx configuration
COPY ./nginx.conf /etc/nginx/nginx.conf

# Expose necessary ports (e.g., 80 for Nginx, 8001 for the backend)
EXPOSE 80 8001

# Start Nginx and the Go backend
CMD ["sh", "-c", "nginx -g 'daemon off;' & /usr/local/bin/main"]
