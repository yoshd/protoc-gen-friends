package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// parse stdin received from protoc.
func parse(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var req plugin.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

// process protobuf schema and return response.
func process(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	var res plugin.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		f := files[fname]
		for _, service := range f.GetService() {

			serviceName := service.GetName()
			methods := service.GetMethod()
			content := makeContent(serviceName, methods)
			outputFileName := "friends.txt"

			res.File = append(res.File, &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(outputFileName),
				Content: proto.String(content),
			})
		}
	}
	return &res
}

func makeContent(serviceName string, methods []*descriptor.MethodDescriptorProto) string {
	var content string
	for _, m := range methods {
		methodName := m.GetName()
		content += fmt.Sprintf("すごーい！%sは%sが得意なフレンズなんだね！\n", serviceName, methodName)
	}
	return content
}

// output response to stdout
func output(res *plugin.CodeGeneratorResponse) error {
	buf, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func main() {
	req, err := parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	res := process(req)

	err = output(res)
	if err != nil {
		panic(err)
	}
}
