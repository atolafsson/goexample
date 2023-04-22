## We specify the base image we need for our
## go application
FROM golang:1.19.3-alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
## we run go build to compile the binary
## executable of our Go program
RUN go build -o goexample .
RUN chmod +x /app/goexample
## app uses the 8081 port
EXPOSE 8081
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/goexample"]