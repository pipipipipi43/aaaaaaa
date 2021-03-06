// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: commit.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CommitDetail)(nil)

// CommitDetail implement urlenc.URLValuesUnmarshaler.
func (m *CommitDetail) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "commitID":
				m.CommitID = vals[0]
			case "repo":
				m.Repo = vals[0]
			case "repoAddr":
				m.RepoAddr = vals[0]
			case "author":
				m.Author = vals[0]
			case "email":
				m.Email = vals[0]
			case "time":
				if m.Time == nil {
					m.Time = &timestamppb.Timestamp{}
				}
			case "time.seconds":
				if m.Time == nil {
					m.Time = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Time.Seconds = val
			case "time.nanos":
				if m.Time == nil {
					m.Time = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Time.Nanos = int32(val)
			case "comment":
				m.Comment = vals[0]
			}
		}
	}
	return nil
}
