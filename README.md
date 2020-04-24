# Truelayer Tech Test

This tech test has been implemented in golang and tested on macOS 10.15.4 with go1.14.1.

Futher to the intial requirements of the tech test this has been implemented with an LRU cache for each external service. This is due to the fact that the translation
API has strict rate limits so allowed for more conclusive testing.

## Running

You can run this project by using the go run command as `go run cmd/cmd.go pokemon`.

It can also be run using docker with `docker build -t truelayer-tech:latest .` to build the image. You can then run `docker run -p 8080:8080 truelayer-tech:latest pokemon`
to run the program as part of a docker container. This is confirmed to be working with docker 19.03.8.

## Testing

Running go tests is as simple as using the command `go test ./...` whilst in the top level directory of this project.

## Project Layout

This project is laid out as a monorepo for all services within the organization. It also includes a design that accommodates for a monobinary of which I have
personally found has helped to improve release cycles of backend services development and improve the reliability of services without the added need of extensive
integration testing between services.

The layout of this project is as such that each service lives within its own folder within the `svc` directory. This includes registering a service command within
the `cmd/cmd.go` file to allow launching from the monobinary.

Subsequently any libraries that are internally developed can be stored within the `lib` folder to allow for easy use across multiple services and contained testing
of each library.

A larger example of a monorepo can be provided if further details is needed on request.
