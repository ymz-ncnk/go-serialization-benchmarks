names=( data_raw data_varint )
gobin=/home/yra/prog/gopath/bin/

for name in "${names[@]}"; do \
    protoc \
    --go_opt=paths=source_relative \
    --go_out=. --plugin protoc-gen-go="${gobin}/protoc-gen-go" \
    --go-vtproto_out=. --plugin protoc-gen-go-vtproto="${gobin}/protoc-gen-go-vtproto" \
    --go-vtproto_opt=features=marshal+unmarshal+size+unmarshal_unsafe \
    ./protobuf/${name}.proto; \
done