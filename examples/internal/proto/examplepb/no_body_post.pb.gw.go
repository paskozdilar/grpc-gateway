// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: examples/internal/proto/examplepb/no_body_post.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_NoBodyPostService_RpcEmptyRpc_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.RpcEmptyRpc(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_NoBodyPostService_RpcEmptyRpc_0(ctx context.Context, marshaler runtime.Marshaler, server NoBodyPostServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.RpcEmptyRpc(ctx, &protoReq)
	return msg, metadata, err
}

func request_NoBodyPostService_RpcEmptyStream_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (NoBodyPostService_RpcEmptyStreamClient, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	stream, err := client.RpcEmptyStream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

func request_NoBodyPostService_RpcEmptyRpcWithResponse_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.RpcEmptyRpcWithResponse(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_NoBodyPostService_RpcEmptyRpcWithResponse_0(ctx context.Context, marshaler runtime.Marshaler, server NoBodyPostServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.RpcEmptyRpcWithResponse(ctx, &protoReq)
	return msg, metadata, err
}

func request_NoBodyPostService_RpcEmptyStreamWithResponse_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (NoBodyPostService_RpcEmptyStreamWithResponseClient, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	stream, err := client.RpcEmptyStreamWithResponse(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

func request_NoBodyPostService_RpcEmptyRpcWithBody_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RpcEmptyRpcWithBody(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_NoBodyPostService_RpcEmptyRpcWithBody_0(ctx context.Context, marshaler runtime.Marshaler, server NoBodyPostServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.RpcEmptyRpcWithBody(ctx, &protoReq)
	return msg, metadata, err
}

func request_NoBodyPostService_RpcEmptyStreamWithBody_0(ctx context.Context, marshaler runtime.Marshaler, client NoBodyPostServiceClient, req *http.Request, pathParams map[string]string) (NoBodyPostService_RpcEmptyStreamWithBodyClient, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	stream, err := client.RpcEmptyStreamWithBody(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterNoBodyPostServiceHandlerServer registers the http handlers for service NoBodyPostService to "mux".
// UnaryRPC     :call NoBodyPostServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNoBodyPostServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterNoBodyPostServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server NoBodyPostServiceServer) error {
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpc", runtime.WithHTTPPathPattern("/rpc/no-body/rpc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NoBodyPostService_RpcEmptyRpc_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpcWithResponse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpcWithResponse", runtime.WithHTTPPathPattern("/rpc/no-body/rpc-with-response"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NoBodyPostService_RpcEmptyRpcWithResponse_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpcWithResponse_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStreamWithResponse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpcWithBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpcWithBody", runtime.WithHTTPPathPattern("/rpc/no-body/rpc-with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NoBodyPostService_RpcEmptyRpcWithBody_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpcWithBody_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStreamWithBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterNoBodyPostServiceHandlerFromEndpoint is same as RegisterNoBodyPostServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNoBodyPostServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterNoBodyPostServiceHandler(ctx, mux, conn)
}

// RegisterNoBodyPostServiceHandler registers the http handlers for service NoBodyPostService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNoBodyPostServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNoBodyPostServiceHandlerClient(ctx, mux, NewNoBodyPostServiceClient(conn))
}

// RegisterNoBodyPostServiceHandlerClient registers the http handlers for service NoBodyPostService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NoBodyPostServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NoBodyPostServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NoBodyPostServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterNoBodyPostServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client NoBodyPostServiceClient) error {
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpc", runtime.WithHTTPPathPattern("/rpc/no-body/rpc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyRpc_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyStream", runtime.WithHTTPPathPattern("/rpc/no-body/stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyStream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyStream_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpcWithResponse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpcWithResponse", runtime.WithHTTPPathPattern("/rpc/no-body/rpc-with-response"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyRpcWithResponse_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpcWithResponse_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStreamWithResponse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyStreamWithResponse", runtime.WithHTTPPathPattern("/rpc/no-body/stream-with-response"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyStreamWithResponse_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyStreamWithResponse_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyRpcWithBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpcWithBody", runtime.WithHTTPPathPattern("/rpc/no-body/rpc-with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyRpcWithBody_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyRpcWithBody_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_NoBodyPostService_RpcEmptyStreamWithBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyStreamWithBody", runtime.WithHTTPPathPattern("/rpc/no-body/stream-with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoBodyPostService_RpcEmptyStreamWithBody_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_NoBodyPostService_RpcEmptyStreamWithBody_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_NoBodyPostService_RpcEmptyRpc_0                = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 0}, []string{"rpc", "no-body"}, ""))
	pattern_NoBodyPostService_RpcEmptyStream_0             = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "no-body", "stream"}, ""))
	pattern_NoBodyPostService_RpcEmptyRpcWithResponse_0    = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "no-body", "rpc-with-response"}, ""))
	pattern_NoBodyPostService_RpcEmptyStreamWithResponse_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "no-body", "stream-with-response"}, ""))
	pattern_NoBodyPostService_RpcEmptyRpcWithBody_0        = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "no-body", "rpc-with-body"}, ""))
	pattern_NoBodyPostService_RpcEmptyStreamWithBody_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "no-body", "stream-with-body"}, ""))
)

var (
	forward_NoBodyPostService_RpcEmptyRpc_0                = runtime.ForwardResponseMessage
	forward_NoBodyPostService_RpcEmptyStream_0             = runtime.ForwardResponseStream
	forward_NoBodyPostService_RpcEmptyRpcWithResponse_0    = runtime.ForwardResponseMessage
	forward_NoBodyPostService_RpcEmptyStreamWithResponse_0 = runtime.ForwardResponseStream
	forward_NoBodyPostService_RpcEmptyRpcWithBody_0        = runtime.ForwardResponseMessage
	forward_NoBodyPostService_RpcEmptyStreamWithBody_0     = runtime.ForwardResponseStream
)
