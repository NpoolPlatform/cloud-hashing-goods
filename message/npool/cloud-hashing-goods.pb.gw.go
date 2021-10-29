// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: npool/cloud-hashing-goods.proto

/*
Package npool is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package npool

import (
	"context"
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
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_CloudHashingGoods_Version_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.Version(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_Version_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.Version(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_CreateVendorLocation_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateVendorLocationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateVendorLocation(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_CreateVendorLocation_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateVendorLocationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateVendorLocation(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_UpdateVendorLocation_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateVendorLocationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateVendorLocation(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_UpdateVendorLocation_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateVendorLocationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateVendorLocation(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_GetVendorLocations_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetVendorLocations(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_GetVendorLocations_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetVendorLocations(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_CreateTargetArea_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateTargetAreaRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateTargetArea(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_CreateTargetArea_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateTargetAreaRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateTargetArea(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_UpdateTargetArea_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateTargetAreaRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateTargetArea(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_UpdateTargetArea_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateTargetAreaRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateTargetArea(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudHashingGoods_GetTargetAreas_0(ctx context.Context, marshaler runtime.Marshaler, client CloudHashingGoodsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetTargetAreas(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudHashingGoods_GetTargetAreas_0(ctx context.Context, marshaler runtime.Marshaler, server CloudHashingGoodsServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetTargetAreas(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCloudHashingGoodsHandlerServer registers the http handlers for service CloudHashingGoods to "mux".
// UnaryRPC     :call CloudHashingGoodsServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCloudHashingGoodsHandlerFromEndpoint instead.
func RegisterCloudHashingGoodsHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CloudHashingGoodsServer) error {

	mux.Handle("GET", pattern_CloudHashingGoods_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/Version", runtime.WithHTTPPathPattern("/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_Version_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_Version_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudHashingGoods_CreateVendorLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateVendorLocation", runtime.WithHTTPPathPattern("/v1/vendor-location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_CreateVendorLocation_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_CreateVendorLocation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_CloudHashingGoods_UpdateVendorLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateVendorLocation", runtime.WithHTTPPathPattern("/v1/vendor-location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_UpdateVendorLocation_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_UpdateVendorLocation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CloudHashingGoods_GetVendorLocations_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/GetVendorLocations", runtime.WithHTTPPathPattern("/v1/vendor-locations"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_GetVendorLocations_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_GetVendorLocations_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudHashingGoods_CreateTargetArea_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateTargetArea", runtime.WithHTTPPathPattern("/v1/target-area"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_CreateTargetArea_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_CreateTargetArea_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_CloudHashingGoods_UpdateTargetArea_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateTargetArea", runtime.WithHTTPPathPattern("/v1/target-area"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_UpdateTargetArea_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_UpdateTargetArea_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CloudHashingGoods_GetTargetAreas_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/GetTargetAreas", runtime.WithHTTPPathPattern("/v1/target-areas"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudHashingGoods_GetTargetAreas_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_GetTargetAreas_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCloudHashingGoodsHandlerFromEndpoint is same as RegisterCloudHashingGoodsHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCloudHashingGoodsHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCloudHashingGoodsHandler(ctx, mux, conn)
}

// RegisterCloudHashingGoodsHandler registers the http handlers for service CloudHashingGoods to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCloudHashingGoodsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCloudHashingGoodsHandlerClient(ctx, mux, NewCloudHashingGoodsClient(conn))
}

// RegisterCloudHashingGoodsHandlerClient registers the http handlers for service CloudHashingGoods
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CloudHashingGoodsClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CloudHashingGoodsClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CloudHashingGoodsClient" to call the correct interceptors.
func RegisterCloudHashingGoodsHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CloudHashingGoodsClient) error {

	mux.Handle("GET", pattern_CloudHashingGoods_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/Version", runtime.WithHTTPPathPattern("/version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_Version_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_Version_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudHashingGoods_CreateVendorLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateVendorLocation", runtime.WithHTTPPathPattern("/v1/vendor-location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_CreateVendorLocation_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_CreateVendorLocation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_CloudHashingGoods_UpdateVendorLocation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateVendorLocation", runtime.WithHTTPPathPattern("/v1/vendor-location"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_UpdateVendorLocation_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_UpdateVendorLocation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CloudHashingGoods_GetVendorLocations_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/GetVendorLocations", runtime.WithHTTPPathPattern("/v1/vendor-locations"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_GetVendorLocations_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_GetVendorLocations_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudHashingGoods_CreateTargetArea_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateTargetArea", runtime.WithHTTPPathPattern("/v1/target-area"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_CreateTargetArea_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_CreateTargetArea_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PATCH", pattern_CloudHashingGoods_UpdateTargetArea_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateTargetArea", runtime.WithHTTPPathPattern("/v1/target-area"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_UpdateTargetArea_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_UpdateTargetArea_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CloudHashingGoods_GetTargetAreas_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/cloud.hashing.goods.v1.CloudHashingGoods/GetTargetAreas", runtime.WithHTTPPathPattern("/v1/target-areas"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudHashingGoods_GetTargetAreas_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudHashingGoods_GetTargetAreas_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CloudHashingGoods_Version_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"version"}, ""))

	pattern_CloudHashingGoods_CreateVendorLocation_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "vendor-location"}, ""))

	pattern_CloudHashingGoods_UpdateVendorLocation_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "vendor-location"}, ""))

	pattern_CloudHashingGoods_GetVendorLocations_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "vendor-locations"}, ""))

	pattern_CloudHashingGoods_CreateTargetArea_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "target-area"}, ""))

	pattern_CloudHashingGoods_UpdateTargetArea_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "target-area"}, ""))

	pattern_CloudHashingGoods_GetTargetAreas_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "target-areas"}, ""))
)

var (
	forward_CloudHashingGoods_Version_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_CreateVendorLocation_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_UpdateVendorLocation_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_GetVendorLocations_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_CreateTargetArea_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_UpdateTargetArea_0 = runtime.ForwardResponseMessage

	forward_CloudHashingGoods_GetTargetAreas_0 = runtime.ForwardResponseMessage
)
