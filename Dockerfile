FROM golang:1.9

# Create folder
RUN mkdir -p /go/src/github.com/aflashyrhetoric/payup-api

# Copy over current folder
ADD . /go/src/github.com/aflashyrhetoric/payup-api

# Set it as the directory in which CMD will be run
WORKDIR /go/src/github.com/aflashyrhetoric/payup-api

# Run the main.go file
CMD ["/go/src/github.com/aflashyrhetoric/payup-api/main"]

EXPOSE 8114