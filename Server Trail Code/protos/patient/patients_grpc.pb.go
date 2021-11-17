// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package patients

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

// PatientServiceClient is the client API for PatientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientServiceClient interface {
	GetPatientById(ctx context.Context, in *GetPatientByIdRequest, opts ...grpc.CallOption) (*Patient, error)
	InsertPatient(ctx context.Context, in *Patient, opts ...grpc.CallOption) (*Patient, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
}

func (c *patientServiceClient) GetPatientById(ctx context.Context, in *GetPatientByIdRequest, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/PatientService/GetPatientById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) InsertPatient(ctx context.Context, in *Patient, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/PatientService/InsertPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	GetPatientById(context.Context, *GetPatientByIdRequest) (*Patient, error)
	InsertPatient(context.Context, *Patient) (*Patient, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
}

func (UnimplementedPatientServiceServer) GetPatientById(context.Context, *GetPatientByIdRequest) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientById not implemented")
}
func (UnimplementedPatientServiceServer) InsertPatient(context.Context, *Patient) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPatient not implemented")
}
func (UnimplementedPatientServiceServer) mustEmbedUnimplementedPatientServiceServer() {}

// UnsafePatientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServiceServer will
// result in compilation errors.
type UnsafePatientServiceServer interface {
	mustEmbedUnimplementedPatientServiceServer()
}

func RegisterPatientServiceServer(s grpc.ServiceRegistrar, srv PatientServiceServer) {
	s.RegisterService(&PatientService_ServiceDesc, srv)
}

func _PatientService_GetPatientById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetPatientById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PatientService/GetPatientById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetPatientById(ctx, req.(*GetPatientByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_InsertPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Patient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).InsertPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PatientService/InsertPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).InsertPatient(ctx, req.(*Patient))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientService_ServiceDesc is the grpc.ServiceDesc for PatientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PatientService",
	HandlerType: (*PatientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatientById",
			Handler:    _PatientService_GetPatientById_Handler,
		},
		{
			MethodName: "InsertPatient",
			Handler:    _PatientService_InsertPatient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/patients.proto",
}
