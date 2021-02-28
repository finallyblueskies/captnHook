# Golang version - Docker Hub (hub.docker.com)
FROM golang:alpine3.12

# Environment variables
ENV APP_NAME bot
ENV PORT 8000

# Open system port
EXPOSE ${PORT}

# Working directory
WORKDIR /go/src/${APP_NAME}
COPY . /go/src/${APP_NAME}

# Install dependencies from mod file
RUN go mod download

# Hot reloading for development!
RUN go get github.com/githubnemo/CompileDaemon

# Build application
RUN go build -o ${APP_NAME} ./cmd/bot

# Run application   
CMD ./${APP_NAME}