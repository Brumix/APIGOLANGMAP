FROM golang:1.17.8-alpine3.15

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc libc-dev


# RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/files
RUN go get -u github.com/swaggo/http-swagger


# HOT RELOAD
RUN go get -u github.com/githubnemo/CompileDaemon


# Set the Current Working Directory inside the container
WORKDIR /go/src/projetoapi

# Copy everything from the current directory to the Working Directory inside the container
COPY . .


RUN go mod tidy


# RUN Swagger
RUN swag init

# Build the Go app
RUN go build -o main .

# Expose port 8081 to the outside world
EXPOSE 8080

# Run the executable DEPLOYMENT
# CMD ["./main"]

# HOT RELOAD
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./main.go" -command="./main"



