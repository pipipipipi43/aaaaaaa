// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: autotest-plan.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/dop/autotest/testplan/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.dop.autotest.testplan",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType               = reflect.TypeOf((*Client)(nil)).Elem()
	testPlanServiceClientType = reflect.TypeOf((*pb.TestPlanServiceClient)(nil)).Elem()
	testPlanServiceServerType = reflect.TypeOf((*pb.TestPlanServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.dop.autotest.testplan-client":
		return p.client
	case "erda.core.dop.autotest.testplan.TestPlanService":
		return &testPlanServiceWrapper{client: p.client.TestPlanService(), opts: opts}
	case "erda.core.dop.autotest.testplan.TestPlanService.client":
		return p.client.TestPlanService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case testPlanServiceClientType:
		return p.client.TestPlanService()
	case testPlanServiceServerType:
		return &testPlanServiceWrapper{client: p.client.TestPlanService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.dop.autotest.testplan-client", &servicehub.Spec{
		Services: []string{
			"erda.core.dop.autotest.testplan.TestPlanService",
			"erda.core.dop.autotest.testplan-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			testPlanServiceClientType,
			// server types
			testPlanServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
