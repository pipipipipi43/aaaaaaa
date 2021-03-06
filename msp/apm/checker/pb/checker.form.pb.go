// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: checker.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CreateCheckerRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateCheckerResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCheckerRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateCheckerResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteCheckerRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteCheckerResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListCheckersRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListCheckersResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckersRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckersResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckerRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DescribeCheckerResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Checker)(nil)

// CreateCheckerRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateCheckerRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "data":
				if m.Data == nil {
					m.Data = &Checker{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				m.Data.Type = vals[0]
			}
		}
	}
	return nil
}

// CreateCheckerResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateCheckerResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// UpdateCheckerRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCheckerRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "id":
				m.Id = vals[0]
			case "data":
				if m.Data == nil {
					m.Data = &Checker{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Checker{}
				}
				m.Data.Type = vals[0]
			}
		}
	}
	return nil
}

// UpdateCheckerResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateCheckerResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// DeleteCheckerRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteCheckerRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// DeleteCheckerResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteCheckerResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// ListCheckersRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListCheckersRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// ListCheckersResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListCheckersResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DescribeCheckersRequest implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckersRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			}
		}
	}
	return nil
}

// DescribeCheckersResponse implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckersResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DescribeCheckerRequest implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckerRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// DescribeCheckerResponse implement urlenc.URLValuesUnmarshaler.
func (m *DescribeCheckerResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Checker implement urlenc.URLValuesUnmarshaler.
func (m *Checker) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}
