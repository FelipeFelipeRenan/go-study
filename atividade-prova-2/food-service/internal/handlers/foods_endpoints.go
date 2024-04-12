package handlers

import (
	"context"
	"encoding/json"
	"foods/internal/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetAllFoodsEndpoint endpoint.Endpoint	
}

func MakeEndpoints(s service.Service) Endpoints{
	return Endpoints{
		GetAllFoodsEndpoint: MakeGetAllFoodsEndpoint(s),
	}
	
}

func MakeGetAllFoodsEndpoint(s service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error){
		foods, err := s.GetAllFoods(ctx)
		if err != nil {
			return nil, err
		}
		return foods, nil
	}

}

func DecodeGetAllFoodsRequest(_ context.Context, r *http.Request)(interface{}, error){
	return nil, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, respose interface{}) error{
	return json.NewEncoder(w).Encode(respose)
}