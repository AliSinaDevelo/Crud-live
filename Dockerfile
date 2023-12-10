# official golang image
FROM golang:1.18.0-alpine3.14
# set working directory
WORKDIR /app
# copy the src code
COPY . .
# dependencies
RUN go get -d -v ./...
# build the binary
RUN go build -o api .
# expose port 8080
EXPOSE 8080
# run the binary
CMD ["./api"]
