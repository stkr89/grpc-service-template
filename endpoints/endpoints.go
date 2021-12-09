package endpoints

import (
	"context"
	"github.com/shopr-org/grpc-service-template/pb"

	"github.com/go-kit/kit/endpoint"
	"github.com/shopr-org/grpc-service-template/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s service.MathService) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.MathRequest)
		return s.Add(ctx, req)
	}
}
