#!/bin/bash

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
cd ./go/cmd/seed/; echo y | make go-db-init; cd -

# echo "================== [ test ] =================="
# while true; do
#     echo "Hello ulala"
#     sleep 10
# done
