#!/usr/bin/env bash

# Define log files
stdout_log="stdout.log"
stderr_log="stderr.log"

# Redirect stdout and stderr
exec > >(tee -a "$stdout_log") 2> >(tee -a "$stderr_log" >&2)

# Your script starts here
echo "Logging stdout to $stdout_log"
echo "Logging stderr to $stderr_log"

docker run -v "$(pwd)":/grpc-gateway -v "$(pwd)"/certs:/etc/ssl/certs -w /grpc-gateway --rm ghcr.io/grpc-ecosystem/grpc-gateway/build-env:1.24 \
    /bin/bash -c 'make install && \
        make clean && \
        make generate'
docker run -itv "$(pwd)":/grpc-gateway -v "$(pwd)"/certs:/etc/ssl/certs -w /grpc-gateway --entrypoint /bin/bash --rm \
    ghcr.io/grpc-ecosystem/grpc-gateway/build-env:1.24 -c '\
        bazel run :gazelle -- update-repos -from_file=go.mod -to_macro=repositories.bzl%go_repositories && \
        bazel run :gazelle && \
        bazel run :buildifier &&
        bazel test //...'

beep
