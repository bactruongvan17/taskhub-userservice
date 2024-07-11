#!/bin/bash

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

go mod tidy

check_program_installed() {
  if [ -x "$(command -v $1)" ]; then
    echo "installed"
    return 0
  else
    echo "no"
    return 1
  fi
}

if check_program_installed "protoc-gen-go"; then
  echo "{protoc-gen-go} is installed."
else
  echo "{protoc-gen-go} is not installed. Downloading..."

  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
fi

if check_program_installed "air"; then
  echo "air is installed."
else
  echo "air is not installed. Downloading..."

  go install github.com/air-verse/air@latest
fi

if check_program_installed "air"; then
  echo "air is installed."
else
  echo "air is not installed. Downloading..."

  go install github.com/air-verse/air@latest
fi

if check_program_installed "protoc"; then
  echo "protoc is installed."
else
  apt update && apt install -y protobuf-compiler
fi

air
