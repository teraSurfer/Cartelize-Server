package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/terasurfer/Cartelize-Server/db/pkg/endpoint"
	pb "github.com/terasurfer/Cartelize-Server/db/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeConnectDBHandler creates the handler logic
func makeConnectDBHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ConnectDBEndpoint, decodeConnectDBRequest, encodeConnectDBResponse, options...)
}

// decodeConnectDBResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ConnectDB request.
// TODO implement the decoder
func decodeConnectDBRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Decoder is not impelemented")
}

// encodeConnectDBResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeConnectDBResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Encoder is not impelemented")
}
func (g *grpcServer) ConnectDB(ctx context1.Context, req *pb.ConnectDBRequest) (*pb.ConnectDBReply, error) {
	_, rep, err := g.connectDB.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ConnectDBReply), nil
}

// makeGetDBHandler creates the handler logic
func makeGetDBHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetDBEndpoint, decodeGetDBRequest, encodeGetDBResponse, options...)
}

// decodeGetDBResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetDB request.
// TODO implement the decoder
func decodeGetDBRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Decoder is not impelemented")
}

// encodeGetDBResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetDBResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Encoder is not impelemented")
}
func (g *grpcServer) GetDB(ctx context1.Context, req *pb.GetDBRequest) (*pb.GetDBReply, error) {
	_, rep, err := g.getDB.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetDBReply), nil
}

// makeInitHandler creates the handler logic
func makeInitHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.InitEndpoint, decodeInitRequest, encodeInitResponse, options...)
}

// decodeInitResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Init request.
// TODO implement the decoder
func decodeInitRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Decoder is not impelemented")
}

// encodeInitResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeInitResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Db' Encoder is not impelemented")
}
func (g *grpcServer) Init(ctx context1.Context, req *pb.InitRequest) (*pb.InitReply, error) {
	_, rep, err := g.init.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.InitReply), nil
}
