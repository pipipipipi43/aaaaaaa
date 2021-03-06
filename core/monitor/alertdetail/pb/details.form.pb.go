// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: details.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*QuerySystemPodMetricsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QuerySystemPodMetricsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PodInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PodInfoSummary)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PodInfoInstanse)(nil)

// QuerySystemPodMetricsRequest implement urlenc.URLValuesUnmarshaler.
func (m *QuerySystemPodMetricsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "timestamp":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Timestamp = val
			}
		}
	}
	return nil
}

// QuerySystemPodMetricsResponse implement urlenc.URLValuesUnmarshaler.
func (m *QuerySystemPodMetricsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
			case "data.summary":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
			case "data.summary.clusterName":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.ClusterName = vals[0]
			case "data.summary.nodeName":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.NodeName = vals[0]
			case "data.summary.hostIP":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.HostIP = vals[0]
			case "data.summary.namespace":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.Namespace = vals[0]
			case "data.summary.podName":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.PodName = vals[0]
			case "data.summary.restartTotal":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Summary.RestartTotal = val
			case "data.summary.stateCode":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Summary.StateCode = val
			case "data.summary.terminatedReason":
				if m.Data == nil {
					m.Data = &PodInfo{}
				}
				if m.Data.Summary == nil {
					m.Data.Summary = &PodInfoSummary{}
				}
				m.Data.Summary.TerminatedReason = vals[0]
			}
		}
	}
	return nil
}

// PodInfo implement urlenc.URLValuesUnmarshaler.
func (m *PodInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "summary":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
			case "summary.clusterName":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.ClusterName = vals[0]
			case "summary.nodeName":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.NodeName = vals[0]
			case "summary.hostIP":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.HostIP = vals[0]
			case "summary.namespace":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.Namespace = vals[0]
			case "summary.podName":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.PodName = vals[0]
			case "summary.restartTotal":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Summary.RestartTotal = val
			case "summary.stateCode":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Summary.StateCode = val
			case "summary.terminatedReason":
				if m.Summary == nil {
					m.Summary = &PodInfoSummary{}
				}
				m.Summary.TerminatedReason = vals[0]
			}
		}
	}
	return nil
}

// PodInfoSummary implement urlenc.URLValuesUnmarshaler.
func (m *PodInfoSummary) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clusterName":
				m.ClusterName = vals[0]
			case "nodeName":
				m.NodeName = vals[0]
			case "hostIP":
				m.HostIP = vals[0]
			case "namespace":
				m.Namespace = vals[0]
			case "podName":
				m.PodName = vals[0]
			case "restartTotal":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.RestartTotal = val
			case "stateCode":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StateCode = val
			case "terminatedReason":
				m.TerminatedReason = vals[0]
			}
		}
	}
	return nil
}

// PodInfoInstanse implement urlenc.URLValuesUnmarshaler.
func (m *PodInfoInstanse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "containerId":
				m.ContainerId = vals[0]
			case "hostIP":
				m.HostIP = vals[0]
			case "startedAt":
				m.StartedAt = vals[0]
			case "finishedAt":
				m.FinishedAt = vals[0]
			case "exitCode":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ExitCode = val
			case "oomKilled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.OomKilled = val
			}
		}
	}
	return nil
}
