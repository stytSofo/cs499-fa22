FROM golang:1.18.4

# set a build-time vriable 
ARG DB=memdb

# set a working directory for the app
WORKDIR /go/src/github.com/ucy-coast/hotel-app

# copy all the files to the container
COPY . .

RUN go install -tags ${DB} -ldflags="-s -w" ./cmd/...
