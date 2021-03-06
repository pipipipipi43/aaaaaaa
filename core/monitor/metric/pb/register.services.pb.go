// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: meta.proto, metric.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterMetricMetaServiceImp meta.proto
func RegisterMetricMetaServiceImp(regester transport.Register, srv MetricMetaServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterMetricMetaServiceHandler(regester, MetricMetaServiceHandler(srv), _ops.HTTP...)
	RegisterMetricMetaServiceServer(regester, srv, _ops.GRPC...)
}

// RegisterMetricServiceImp metric.proto
func RegisterMetricServiceImp(regester transport.Register, srv MetricServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterMetricServiceHandler(regester, MetricServiceHandler(srv), _ops.HTTP...)
	RegisterMetricServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.monitor.metric.MetricMetaService",
		"erda.core.monitor.metric.MetricService",
	)
}

var (
	metricMetaServiceClientType  = reflect.TypeOf((*MetricMetaServiceClient)(nil)).Elem()
	metricMetaServiceServerType  = reflect.TypeOf((*MetricMetaServiceServer)(nil)).Elem()
	metricMetaServiceHandlerType = reflect.TypeOf((*MetricMetaServiceHandler)(nil)).Elem()

	metricServiceClientType  = reflect.TypeOf((*MetricServiceClient)(nil)).Elem()
	metricServiceServerType  = reflect.TypeOf((*MetricServiceServer)(nil)).Elem()
	metricServiceHandlerType = reflect.TypeOf((*MetricServiceHandler)(nil)).Elem()
)

// MetricMetaServiceClientType .
func MetricMetaServiceClientType() reflect.Type { return metricMetaServiceClientType }

// MetricMetaServiceServerType .
func MetricMetaServiceServerType() reflect.Type { return metricMetaServiceServerType }

// MetricMetaServiceHandlerType .
func MetricMetaServiceHandlerType() reflect.Type { return metricMetaServiceHandlerType }

// MetricServiceClientType .
func MetricServiceClientType() reflect.Type { return metricServiceClientType }

// MetricServiceServerType .
func MetricServiceServerType() reflect.Type { return metricServiceServerType }

// MetricServiceHandlerType .
func MetricServiceHandlerType() reflect.Type { return metricServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		metricMetaServiceClientType,
		metricServiceClientType,
		// server types
		metricMetaServiceServerType,
		metricServiceServerType,
		// handler types
		metricMetaServiceHandlerType,
		metricServiceHandlerType,
	}
}
