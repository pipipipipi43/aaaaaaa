// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: autotest-plan.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*TestPlanUpdateByHookRequest)(nil)
var _ json.Unmarshaler = (*TestPlanUpdateByHookRequest)(nil)
var _ json.Marshaler = (*Content)(nil)
var _ json.Unmarshaler = (*Content)(nil)
var _ json.Marshaler = (*TestPlanUpdateByHookResponse)(nil)
var _ json.Unmarshaler = (*TestPlanUpdateByHookResponse)(nil)

// TestPlanUpdateByHookRequest implement json.Marshaler.
func (m *TestPlanUpdateByHookRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TestPlanUpdateByHookRequest implement json.Marshaler.
func (m *TestPlanUpdateByHookRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Content implement json.Marshaler.
func (m *Content) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Content implement json.Marshaler.
func (m *Content) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TestPlanUpdateByHookResponse implement json.Marshaler.
func (m *TestPlanUpdateByHookResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TestPlanUpdateByHookResponse implement json.Marshaler.
func (m *TestPlanUpdateByHookResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
