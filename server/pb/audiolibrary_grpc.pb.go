// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: audiolibrary.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AudiolibraryClient is the client API for Audiolibrary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudiolibraryClient interface {
	GetAudiofilePath(ctx context.Context, in *AudiolibraryRequest, opts ...grpc.CallOption) (*AudiolibraryResponse, error)
}

type audiolibraryClient struct {
	cc grpc.ClientConnInterface
}

func NewAudiolibraryClient(cc grpc.ClientConnInterface) AudiolibraryClient {
	return &audiolibraryClient{cc}
}

func (c *audiolibraryClient) GetAudiofilePath(ctx context.Context, in *AudiolibraryRequest, opts ...grpc.CallOption) (*AudiolibraryResponse, error) {
	out := new(AudiolibraryResponse)
	err := c.cc.Invoke(ctx, "/Audiolibrary/GetAudiofilePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudiolibraryServer is the server API for Audiolibrary service.
// All implementations must embed UnimplementedAudiolibraryServer
// for forward compatibility
type AudiolibraryServer interface {
	GetAudiofilePath(context.Context, *AudiolibraryRequest) (*AudiolibraryResponse, error)
	mustEmbedUnimplementedAudiolibraryServer()
}

// UnimplementedAudiolibraryServer must be embedded to have forward compatible implementations.
type UnimplementedAudiolibraryServer struct {
}

func (UnimplementedAudiolibraryServer) GetAudiofilePath(context.Context, *AudiolibraryRequest) (*AudiolibraryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAudiofilePath not implemented")
}
func (UnimplementedAudiolibraryServer) mustEmbedUnimplementedAudiolibraryServer() {}

// UnsafeAudiolibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudiolibraryServer will
// result in compilation errors.
type UnsafeAudiolibraryServer interface {
	mustEmbedUnimplementedAudiolibraryServer()
}

func RegisterAudiolibraryServer(s grpc.ServiceRegistrar, srv AudiolibraryServer) {
	s.RegisterService(&Audiolibrary_ServiceDesc, srv)
}

func _Audiolibrary_GetAudiofilePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AudiolibraryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudiolibraryServer).GetAudiofilePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Audiolibrary/GetAudiofilePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudiolibraryServer).GetAudiofilePath(ctx, req.(*AudiolibraryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Audiolibrary_ServiceDesc is the grpc.ServiceDesc for Audiolibrary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Audiolibrary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Audiolibrary",
	HandlerType: (*AudiolibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAudiofilePath",
			Handler:    _Audiolibrary_GetAudiofilePath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audiolibrary.proto",
}
