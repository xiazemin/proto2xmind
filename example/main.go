package main

import (
	gen "github.com/xiazemin/proto2xmind/proto"
)

func main() {
	gen.GenFiles("./example/example.xmind", "./example/sub.proto", "./example/request.proto")
}
