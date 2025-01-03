// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: drive/v1/drive.proto

package drivev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/taubyte/tau/pkg/spore-drive/proto/gen/drive/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// DriveServiceName is the fully-qualified name of the DriveService service.
	DriveServiceName = "drive.v1.DriveService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DriveServiceNewProcedure is the fully-qualified name of the DriveService's New RPC.
	DriveServiceNewProcedure = "/drive.v1.DriveService/New"
	// DriveServicePlotProcedure is the fully-qualified name of the DriveService's Plot RPC.
	DriveServicePlotProcedure = "/drive.v1.DriveService/Plot"
	// DriveServiceDisplaceProcedure is the fully-qualified name of the DriveService's Displace RPC.
	DriveServiceDisplaceProcedure = "/drive.v1.DriveService/Displace"
	// DriveServiceProgressProcedure is the fully-qualified name of the DriveService's Progress RPC.
	DriveServiceProgressProcedure = "/drive.v1.DriveService/Progress"
	// DriveServiceAbortProcedure is the fully-qualified name of the DriveService's Abort RPC.
	DriveServiceAbortProcedure = "/drive.v1.DriveService/Abort"
	// DriveServiceFreeProcedure is the fully-qualified name of the DriveService's Free RPC.
	DriveServiceFreeProcedure = "/drive.v1.DriveService/Free"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	driveServiceServiceDescriptor        = v1.File_drive_v1_drive_proto.Services().ByName("DriveService")
	driveServiceNewMethodDescriptor      = driveServiceServiceDescriptor.Methods().ByName("New")
	driveServicePlotMethodDescriptor     = driveServiceServiceDescriptor.Methods().ByName("Plot")
	driveServiceDisplaceMethodDescriptor = driveServiceServiceDescriptor.Methods().ByName("Displace")
	driveServiceProgressMethodDescriptor = driveServiceServiceDescriptor.Methods().ByName("Progress")
	driveServiceAbortMethodDescriptor    = driveServiceServiceDescriptor.Methods().ByName("Abort")
	driveServiceFreeMethodDescriptor     = driveServiceServiceDescriptor.Methods().ByName("Free")
)

// DriveServiceClient is a client for the drive.v1.DriveService service.
type DriveServiceClient interface {
	New(context.Context, *connect.Request[v1.DriveRequest]) (*connect.Response[v1.Drive], error)
	Plot(context.Context, *connect.Request[v1.PlotRequest]) (*connect.Response[v1.Course], error)
	Displace(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error)
	Progress(context.Context, *connect.Request[v1.Course]) (*connect.ServerStreamForClient[v1.DisplacementProgress], error)
	Abort(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error)
	Free(context.Context, *connect.Request[v1.Drive]) (*connect.Response[v1.Empty], error)
}

// NewDriveServiceClient constructs a client for the drive.v1.DriveService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDriveServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DriveServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &driveServiceClient{
		new: connect.NewClient[v1.DriveRequest, v1.Drive](
			httpClient,
			baseURL+DriveServiceNewProcedure,
			connect.WithSchema(driveServiceNewMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		plot: connect.NewClient[v1.PlotRequest, v1.Course](
			httpClient,
			baseURL+DriveServicePlotProcedure,
			connect.WithSchema(driveServicePlotMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		displace: connect.NewClient[v1.Course, v1.Empty](
			httpClient,
			baseURL+DriveServiceDisplaceProcedure,
			connect.WithSchema(driveServiceDisplaceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		progress: connect.NewClient[v1.Course, v1.DisplacementProgress](
			httpClient,
			baseURL+DriveServiceProgressProcedure,
			connect.WithSchema(driveServiceProgressMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		abort: connect.NewClient[v1.Course, v1.Empty](
			httpClient,
			baseURL+DriveServiceAbortProcedure,
			connect.WithSchema(driveServiceAbortMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		free: connect.NewClient[v1.Drive, v1.Empty](
			httpClient,
			baseURL+DriveServiceFreeProcedure,
			connect.WithSchema(driveServiceFreeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// driveServiceClient implements DriveServiceClient.
type driveServiceClient struct {
	new      *connect.Client[v1.DriveRequest, v1.Drive]
	plot     *connect.Client[v1.PlotRequest, v1.Course]
	displace *connect.Client[v1.Course, v1.Empty]
	progress *connect.Client[v1.Course, v1.DisplacementProgress]
	abort    *connect.Client[v1.Course, v1.Empty]
	free     *connect.Client[v1.Drive, v1.Empty]
}

// New calls drive.v1.DriveService.New.
func (c *driveServiceClient) New(ctx context.Context, req *connect.Request[v1.DriveRequest]) (*connect.Response[v1.Drive], error) {
	return c.new.CallUnary(ctx, req)
}

// Plot calls drive.v1.DriveService.Plot.
func (c *driveServiceClient) Plot(ctx context.Context, req *connect.Request[v1.PlotRequest]) (*connect.Response[v1.Course], error) {
	return c.plot.CallUnary(ctx, req)
}

// Displace calls drive.v1.DriveService.Displace.
func (c *driveServiceClient) Displace(ctx context.Context, req *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error) {
	return c.displace.CallUnary(ctx, req)
}

// Progress calls drive.v1.DriveService.Progress.
func (c *driveServiceClient) Progress(ctx context.Context, req *connect.Request[v1.Course]) (*connect.ServerStreamForClient[v1.DisplacementProgress], error) {
	return c.progress.CallServerStream(ctx, req)
}

// Abort calls drive.v1.DriveService.Abort.
func (c *driveServiceClient) Abort(ctx context.Context, req *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error) {
	return c.abort.CallUnary(ctx, req)
}

// Free calls drive.v1.DriveService.Free.
func (c *driveServiceClient) Free(ctx context.Context, req *connect.Request[v1.Drive]) (*connect.Response[v1.Empty], error) {
	return c.free.CallUnary(ctx, req)
}

// DriveServiceHandler is an implementation of the drive.v1.DriveService service.
type DriveServiceHandler interface {
	New(context.Context, *connect.Request[v1.DriveRequest]) (*connect.Response[v1.Drive], error)
	Plot(context.Context, *connect.Request[v1.PlotRequest]) (*connect.Response[v1.Course], error)
	Displace(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error)
	Progress(context.Context, *connect.Request[v1.Course], *connect.ServerStream[v1.DisplacementProgress]) error
	Abort(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error)
	Free(context.Context, *connect.Request[v1.Drive]) (*connect.Response[v1.Empty], error)
}

// NewDriveServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDriveServiceHandler(svc DriveServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	driveServiceNewHandler := connect.NewUnaryHandler(
		DriveServiceNewProcedure,
		svc.New,
		connect.WithSchema(driveServiceNewMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	driveServicePlotHandler := connect.NewUnaryHandler(
		DriveServicePlotProcedure,
		svc.Plot,
		connect.WithSchema(driveServicePlotMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	driveServiceDisplaceHandler := connect.NewUnaryHandler(
		DriveServiceDisplaceProcedure,
		svc.Displace,
		connect.WithSchema(driveServiceDisplaceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	driveServiceProgressHandler := connect.NewServerStreamHandler(
		DriveServiceProgressProcedure,
		svc.Progress,
		connect.WithSchema(driveServiceProgressMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	driveServiceAbortHandler := connect.NewUnaryHandler(
		DriveServiceAbortProcedure,
		svc.Abort,
		connect.WithSchema(driveServiceAbortMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	driveServiceFreeHandler := connect.NewUnaryHandler(
		DriveServiceFreeProcedure,
		svc.Free,
		connect.WithSchema(driveServiceFreeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/drive.v1.DriveService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DriveServiceNewProcedure:
			driveServiceNewHandler.ServeHTTP(w, r)
		case DriveServicePlotProcedure:
			driveServicePlotHandler.ServeHTTP(w, r)
		case DriveServiceDisplaceProcedure:
			driveServiceDisplaceHandler.ServeHTTP(w, r)
		case DriveServiceProgressProcedure:
			driveServiceProgressHandler.ServeHTTP(w, r)
		case DriveServiceAbortProcedure:
			driveServiceAbortHandler.ServeHTTP(w, r)
		case DriveServiceFreeProcedure:
			driveServiceFreeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDriveServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDriveServiceHandler struct{}

func (UnimplementedDriveServiceHandler) New(context.Context, *connect.Request[v1.DriveRequest]) (*connect.Response[v1.Drive], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.New is not implemented"))
}

func (UnimplementedDriveServiceHandler) Plot(context.Context, *connect.Request[v1.PlotRequest]) (*connect.Response[v1.Course], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.Plot is not implemented"))
}

func (UnimplementedDriveServiceHandler) Displace(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.Displace is not implemented"))
}

func (UnimplementedDriveServiceHandler) Progress(context.Context, *connect.Request[v1.Course], *connect.ServerStream[v1.DisplacementProgress]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.Progress is not implemented"))
}

func (UnimplementedDriveServiceHandler) Abort(context.Context, *connect.Request[v1.Course]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.Abort is not implemented"))
}

func (UnimplementedDriveServiceHandler) Free(context.Context, *connect.Request[v1.Drive]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drive.v1.DriveService.Free is not implemented"))
}
