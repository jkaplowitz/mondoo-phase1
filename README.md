# Purpose

Web server to demonstrate basic competency at coding in Go. It will serve "Hello from Mondoo Engineer!" in response to any request.

# Requirements

Nothing beyond the standard Go toolchain and standard library.

# Instructions

## Running locally

Run main.go. It will listen on all network interfaces on the TCP port specified by the `PORT` environment variable, or on TCP port 8080 if that environment variable is not specified.

## Running in a Container

Publish a release. This will build and publish a binary release asset via GitHub Actions and then trigger the [jkaplowitz/mondoo-phase2](https://github.com/jkaplowitz/mondoo-phase2) repository to build and publish a Docker image. This image will be available to Docker (and Kubernetes) as `ghcr.io/jkaplowitz/mondoo-phase2:main`, configured to listen within the container on port 8080 unless overridden by a PORT environment variable setting, and with port 8080 configured in the Dockerfilie as an exposed port.