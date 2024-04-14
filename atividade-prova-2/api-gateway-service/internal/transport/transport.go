package transport

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeGetAllFoodsHandler( e endpoint.Endpoint) http.Handler{
	return httptransport.NewServer(
		e,
		DecodeGetAllFoodsRequest,
		EncodeResponse,
	)
}