// Code generated by protoc-gen-go.
// source: github.com/golang/protobuf/ptypes/timestamp/timestamp.proto
// DO NOT EDIT!

/*
Package timestamp is a generated protocol buffer package.

It is generated from these files:
	github.com/golang/protobuf/ptypes/timestamp/timestamp.proto

It has these top-level messages:
	Timestamp
*/
package timestamp

import (
	proto "github.com/golang/protobuf/proto"
)



// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal



// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}