
package calc

import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2 
type AddRequest struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2" json:"n2,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *AddRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type AddReply struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}


func init() {
	proto.RegisterType((*AddRequest)(nil), "calc.AddRequest")
	proto.RegisterType((*AddReply)(nil), "calc.AddReply")
}

var _ context.Context
var _ grpc.ClientConn

const _ = grpc.SupportPackageIsVersion4

type CalculatorClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/calc.Calculator/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


// Server API for Calculator service

type CalculatorServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

