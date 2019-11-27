# we will use a multi stage build to:
# first download the go module dependencies,
# then build the go binary,
# and finally copy it into the production image

# golang 1.13 as base image for build stage
FROM golang:1.13 AS downloader

# git is required for fetching the dependencies.
RUN apt-get update && apt-get install -y git && apt-get install -y ssh

# set app dir
WORKDIR /piggy

# copy go module files that specify our dependencies
COPY go.mod go.sum ./

# download go module dependencies
RUN go mod download

# new stage
FROM golang:1.13 AS builder
COPY --from=downloader /go /go
# now that we have downloaded all dependencies,
# copy the remaining contents over
# by downloading module dependencies prior to copying,
# the dependencies get cached so that they aren't redownloaded each time.
WORKDIR /piggy
COPY . .

# change working directory to the build location
WORKDIR /piggy/cmd/scan
# build the binary, outputing to /go/bin/piggy
RUN go build -o /go/bin/piggy-scan

# this is the actual base image for our production image
FROM gcr.io/google-appengine/golang
# copy our statically compiled binary in to the production image
COPY --from=builder /go/bin/piggy-scan /go/bin/piggy-scan
# declare our entry point for the container as the built binary
# CMD ["/go/bin/piggy-scan"]