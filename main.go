package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/PbPipes/Astounding/"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"

	"github.com/golang/glog"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var (
	globalPkg = &ProtoPackage{
		name:     "",
		parent:   nil,
		children: make(map[string]*ProtoPackage),
		types:    make(map[string]*descriptor.DescriptorProto),
	}
)

// ProtoPackage describes a package of Protobuf, which is an container of message types.
type ProtoPackage struct {
	name     string
	parent   *ProtoPackage
	children map[string]*ProtoPackage
	types    map[string]*descriptor.DescriptorProto
}

func getPipelineMessageOptions(msg *descriptor.DescriptorProto) (*protos.PipelineMessageOptions, error) {
	options := msg.GetOptions()
	if options == nil {
		return nil, nil
	}

	if !proto.HasExtension(options, PipelineOpts) {
		return nil, nil
	}

	return proto.GetExtension(options, PipelineOpts).(*protos.PipelineOptions), nil
}

func convertFile(file *descriptor.FileDescriptorProto) ([]*plugin.CodeGeneratorResponse_File, error) {
	name := path.Base(file.GetName())
	/*
		pkg, ok := globalPkg.relativelyLookupPackage(file.GetPackage())
		if !ok {
			return nil, fmt.Errorf("no such package found: %s", file.GetPackage())
		}
	*/

	response := []*plugin.CodeGeneratorResponse_File{}
	for msgIndex, msg := range file.GetMessageType() {
		path := fmt.Sprintf("%d.%d", messagePath, msgIndex)

		opts, err := getPipelineMessageOptions(msg)
		if err != nil {
			return nil, err
		}
		if opts == nil {
			continue
		}

		tableName := opts.GetPubSubTopicName()
		if len(pubSubTopicName) == 0 {
			continue
		}

		/*
			glog.V(2).Info("Generating schema for a message type ", msg.GetName())
			schema, err := convertMessageType(pkg, msg, opts, make(map[*descriptor.DescriptorProto]bool), comments, path)
			if err != nil {
				glog.Errorf("Failed to convert %s: %v", name, err)
				return nil, err
			}

			jsonSchema, err := json.MarshalIndent(schema, "", " ")
			if err != nil {
				glog.Error("Failed to encode schema", err)
				return nil, err
			}
		*/

		resFile := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(fmt.Sprintf("%s/%s.schema", strings.Replace(file.GetPackage(), ".", "/", -1), tableName)),
			Content: proto.String(string(jsonSchema)),
		}
		response = append(response, resFile)
	}

	return response, nil
}

func convert(req *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	generateTargets := make(map[string]bool)
	for _, file := range req.GetFileToGenerate() {
		generateTargets[file] = true
	}

	res := &plugin.CodeGeneratorResponse{}
	for _, file := range req.GetProtoFile() {
		for msgIndex, msg := range file.GetMessageType() {
			glog.V(1).Infof("Loading a message type %s from package %s", msg.GetName(), file.GetPackage())
			registerType(file.Package, msg, ParseComments(file), fmt.Sprintf("%d.%d", messagePath, msgIndex))
		}
	}
	for _, file := range req.GetProtoFile() {
		if _, ok := generateTargets[file.GetName()]; ok {
			glog.V(1).Info("Converting ", file.GetName())
			converted, err := convertFile(file)
			if err != nil {
				res.Error = proto.String(fmt.Sprintf("Failed to convert %s: %v", file.GetName(), err))
				return res, err
			}
			res.File = append(res.File, converted...)
		}
	}
	return res, nil
}

func convertFrom(rd io.Reader) (*plugin.CodeGeneratorResponse, error) {
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		fmt.Error("Failed to read request:", err)
		return nil, err
	}
	req := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(input, req)
	if err != nil {
		glog.Error("Can't unmarshal input:", err)
		return nil, err
	}

	glog.V(1).Info("Converting input")
	return convert(req)
}

func main() {
	flag.Parse()
	ok := true
	glog.Info("Processing code generator request")
	res, err := convertFrom(os.Stdin)
	if err != nil {
		ok = false
		if res == nil {
			message := fmt.Sprintf("Failed to read input: %v", err)
			res = &plugin.CodeGeneratorResponse{
				Error: &message,
			}
		}
	}

	glog.Info("Serializing code generator response")
	data, err := proto.Marshal(res)
	if err != nil {
		glog.Fatal("Cannot marshal response", err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		glog.Fatal("Failed to write response", err)
	}

	if ok {
		glog.Info("Succeeded to process code generator request")
	} else {
		glog.Info("Failed to process code generator but successfully sent the error to protoc")
		os.Exit(1)
	}
}
