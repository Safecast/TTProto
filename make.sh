## protoc executable is placed on the standard path after installing standard PB support
##old was "protoc --go_out=./golang tt.proto"
protoc --go_out=./golang --go_opt=Mtt.proto=. tt.proto

## Download and unzip nanopb sdk into a folder, and point to it with the env var NANOPBSDK
## https://koti.kapsi.fi/jpa/nanopb/
## Then,
export NANOPBSDK=~/dev/nordic/external/nano-pb
${NANOPBSDK}/generator-bin/protoc -I./ -I${NANOPBSDK}/generator/proto/ --nanopb_out=./clang tt.proto
