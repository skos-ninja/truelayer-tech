FROM golang:latest as builder

LABEL maintainer="Jake Oliver <docker@skos.ninja>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# We copy in the complete source after downloading dependencies
# to allow us to cache image layers to reduce build time.
COPY . .

# Set some go build options
ENV CGO_ENABLED=0
ENV GOOS=linux

# Run a build against the command directory.
# For the purposes of using sqlite we require CGO to be enabled
RUN go build -a -installsuffix cgo -o main cmd/cmd.go

# Create new image and import just the binary
FROM alpine:latest

# Alpine doesn't include timezones
RUN apk --no-cache add tzdata

# Alpine doesn't include cert auth certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

# Set our github.com/gin-gonic/gin library to release mode
ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["./main"]