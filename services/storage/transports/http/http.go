package storage_httptransport



import (
       "log"
	"net/http"
	"encoding/json"
	"context"

    pb "github.com/gofunct/services"
    gokit_endpoint "github.com/go-kit/kit/endpoint"
    httptransport "github.com/go-kit/kit/transport/http"
    endpoints "github.com/gofunct/services/endpoints"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient



	
		func MakeStoreHandler(svc pb.StorageServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeStoreRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeStoreRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.StoreRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeDeleteHandler(svc pb.StorageServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeDeleteRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.DeleteRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(svc pb.StorageServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
		
			log.Println("new HTTP endpoint: \"/Store\" (service=Storage)")
			mux.Handle("/Store", MakeStoreHandler(svc, endpoints.StoreEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Delete\" (service=Storage)")
			mux.Handle("/Delete", MakeDeleteHandler(svc, endpoints.DeleteEndpoint))
		
	
	return nil
}
