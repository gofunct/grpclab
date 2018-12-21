package image_httptransport



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



	

	
		func MakeStoreSyncHandler(svc pb.ImageServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeStoreSyncRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeStoreSyncRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.ImageStoreRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeDeleteHandler(svc pb.ImageServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
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

func RegisterHandlers(svc pb.ImageServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
		
	
		
			log.Println("new HTTP endpoint: \"/StoreSync\" (service=Image)")
			mux.Handle("/StoreSync", MakeStoreSyncHandler(svc, endpoints.StoreSyncEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Delete\" (service=Image)")
			mux.Handle("/Delete", MakeDeleteHandler(svc, endpoints.DeleteEndpoint))
		
	
	return nil
}
