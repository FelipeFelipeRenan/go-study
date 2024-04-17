package grpc

import (
	"context"
	"errors"
	"movieexample/gen"
	"movieexample/metadata/internal/controller/metadata"
	"movieexample/metadata/pkg/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie metadata gRPC handler
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	crtl *metadata.Controller
}

// New creates a new movie metadata gRPC handler
func New(ctrl *metadata.Controller) *Handler {go mo
	return &Handler{crtl: ctrl}
}

// GetMetadataByID returns a movie metadata by its id
func (h *Handler) GetMetadataByID(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil request or empty id")
	}
	m, err := h.crtl.Get(ctx, req.MovieId)
	if err == nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil
}
