FROM golang:1.9
RUN mkdir -p /go/src/github.com/aflashyrhetoric/payup-api
ADD ./build /go/src/github.com/aflashyrhetoric/payup-api/
WORKDIR /go/src/github.com/aflashyrhetoric/payup-api
CMD ["/go/src/github.com/aflashyrhetoric/payup-api/main"]

EXPOSE 8114