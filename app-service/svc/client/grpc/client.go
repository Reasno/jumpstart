// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 819426ab18
// Version Date: 2020-10-16T09:15:44Z

// Package grpc provides a gRPC client for the App service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/Reasno/jumpstart"
	"github.com/Reasno/jumpstart/app-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.AppServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(
			conn,
			"jumpstart.App",
			"Create",
			EncodeGRPCCreateRequest,
			DecodeGRPCCreateResponse,
			pb.GenericReply{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		CreateEndpoint: createEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCCreateResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC create reply to a user-domain create response. Primarily useful in a client.
func DecodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GenericReply)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCCreateRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain create request to a gRPC create request. Primarily useful in a client.
func EncodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}