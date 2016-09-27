## protoc executable is placed on the standard path after installing standard PB support
protoc --go_out=./golang teletype.proto

## Download and unzip nanopb sdk into a folder, and point to it with the env var NANOPBSDK
## https://koti.kapsi.fi/jpa/nanopb/
## Then,
export NANOPBSDK=~/dev/nano/nanopb
${NANOPBSDK}/generator-bin/protoc -I./ -I${NANOPBSDK}/generator/proto/ --nanopb_out=./clang teletype.proto
