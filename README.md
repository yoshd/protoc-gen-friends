# protoc-gen-friends

protoc-gen-friends is a garbage-like protoc plugin that just generates `friends.txt` like this:

```txt:friends.txt
すごーい！<GRPC_SERVICE_NAME>は<GRPC_METHOD_NAME1>が得意なフレンズなんだね！
すごーい！<GRPC_SERVICE_NAME>は<GRPC_METHOD_NAME2>が得意なフレンズなんだね！
```

## Installation

```shell
go get -u github.com/yoshd/protoc-gen-friends
```

## Invoking the Plugin

```shell
protoc -I. --plugin=path/to/protoc-gen-friends --friends_out=. your.proto
```
