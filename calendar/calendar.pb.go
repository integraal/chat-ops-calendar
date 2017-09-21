// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

/*
Package calendar is a generated protocol buffer package.

It is generated from these files:
	calendar.proto

It has these top-level messages:
	EventRequest
	Event
*/
package calendar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type EventRequest struct {
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EventRequest) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *EventRequest) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *EventRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// The response message containing the event
type Event struct {
	Uid         string            `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Id          string            `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Start       int64             `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	End         int64             `protobuf:"varint,4,opt,name=end" json:"end,omitempty"`
	Summary     string            `protobuf:"bytes,5,opt,name=summary" json:"summary,omitempty"`
	Description string            `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Attendees   []*Event_Attendee `protobuf:"bytes,7,rep,name=attendees" json:"attendees,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Event) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Event) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetAttendees() []*Event_Attendee {
	if m != nil {
		return m.Attendees
	}
	return nil
}

type Event_Attendee struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *Event_Attendee) Reset()                    { *m = Event_Attendee{} }
func (m *Event_Attendee) String() string            { return proto.CompactTextString(m) }
func (*Event_Attendee) ProtoMessage()               {}
func (*Event_Attendee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Event_Attendee) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRequest)(nil), "calendar.EventRequest")
	proto.RegisterType((*Event)(nil), "calendar.Event")
	proto.RegisterType((*Event_Attendee)(nil), "calendar.Event.Attendee")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calendar service

type CalendarClient interface {
	// Sends a greeting
	GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (Calendar_GetEventsClient, error)
}

type calendarClient struct {
	cc *grpc.ClientConn
}

func NewCalendarClient(cc *grpc.ClientConn) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (Calendar_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Calendar_serviceDesc.Streams[0], c.cc, "/calendar.Calendar/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &calendarGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calendar_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type calendarGetEventsClient struct {
	grpc.ClientStream
}

func (x *calendarGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Calendar service

type CalendarServer interface {
	// Sends a greeting
	GetEvents(*EventRequest, Calendar_GetEventsServer) error
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalendarServer).GetEvents(m, &calendarGetEventsServer{stream})
}

type Calendar_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type calendarGetEventsServer struct {
	grpc.ServerStream
}

func (x *calendarGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Calendar_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calendar.proto",
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xc4, 0xb4, 0xc9, 0x54, 0xaa, 0x2c, 0x22, 0x4b, 0x4f, 0x21, 0xa7, 0x9e, 0x82,
	0x54, 0x10, 0xaf, 0xa2, 0xe2, 0xc5, 0x53, 0xde, 0x60, 0xed, 0xce, 0x61, 0xa1, 0xd9, 0xd4, 0xdd,
	0x89, 0xe0, 0x13, 0xfb, 0x1a, 0xb2, 0xd3, 0xc4, 0x94, 0xdc, 0xe6, 0xff, 0xf3, 0xcf, 0x37, 0x99,
	0x59, 0x58, 0xef, 0xf5, 0x01, 0x9d, 0xd1, 0xbe, 0x3e, 0xfa, 0x8e, 0x3a, 0x99, 0x8f, 0xba, 0xfa,
	0x80, 0xab, 0xb7, 0x6f, 0x74, 0xd4, 0xe0, 0x57, 0x8f, 0x81, 0xe4, 0x2d, 0x64, 0x81, 0xb4, 0x27,
	0x25, 0x4a, 0xb1, 0x2d, 0x9a, 0x93, 0x90, 0x37, 0x90, 0xa2, 0x33, 0x2a, 0x61, 0x2f, 0x96, 0x31,
	0x87, 0xad, 0xb6, 0x07, 0x95, 0x9e, 0x72, 0x2c, 0xaa, 0x5f, 0x01, 0x19, 0xe3, 0x62, 0x47, 0x6f,
	0xcd, 0x40, 0x89, 0xa5, 0x5c, 0x43, 0x62, 0x47, 0x44, 0x62, 0xcd, 0x34, 0x29, 0x12, 0xd2, 0xd9,
	0xa4, 0x4b, 0xf6, 0x78, 0x92, 0x82, 0x65, 0xe8, 0xdb, 0x56, 0xfb, 0x1f, 0x95, 0x71, 0xf3, 0x28,
	0x65, 0x09, 0x2b, 0x83, 0x61, 0xef, 0xed, 0x91, 0x6c, 0xe7, 0xd4, 0x82, 0xbf, 0x9e, 0x5b, 0xf2,
	0x11, 0x0a, 0x4d, 0x84, 0xce, 0x20, 0x06, 0xb5, 0x2c, 0xd3, 0xed, 0x6a, 0xa7, 0xea, 0xff, 0x5b,
	0xf0, 0x9f, 0xd6, 0xcf, 0x43, 0xa0, 0x99, 0xa2, 0x9b, 0x12, 0xf2, 0xd1, 0x9e, 0x36, 0x15, 0x67,
	0x9b, 0xee, 0x5e, 0x21, 0x7f, 0x19, 0x38, 0xf2, 0x09, 0x8a, 0x77, 0x24, 0xa6, 0x05, 0x79, 0x37,
	0xe3, 0x0f, 0x87, 0xdd, 0x5c, 0xcf, 0xfc, 0xea, 0xe2, 0x5e, 0x7c, 0x2e, 0xf8, 0x39, 0x1e, 0xfe,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0x1e, 0x4e, 0x9f, 0xa0, 0x01, 0x00, 0x00,
}
