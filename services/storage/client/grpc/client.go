package storage_clientgrpc

import (
	context "context"

        "github.com/go-kit/kit/log"
        "google.golang.org/grpc"
        grpctransport "github.com/go-kit/kit/transport/grpc"
        "github.com/go-kit/kit/endpoint"
        jwt "github.com/go-kit/kit/auth/jwt"

        pb "github.com/gofunct/services"
        endpoints "github.com/gofunct/services/endpoints"
)



func New(conn *grpc.ClientConn, logger log.Logger) pb.StorageServiceServer {
        
		
			var storeEndpoint endpoint.Endpoint
			{
				storeEndpoint = grpctransport.NewClient(
					conn,
					"storage.StorageService",
					"Store",
					EncodeStoreRequest,
					DecodeStoreResponse,
					pb.StoreResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var deleteEndpoint endpoint.Endpoint
			{
				deleteEndpoint = grpctransport.NewClient(
					conn,
					"storage.StorageService",
					"Delete",
					EncodeDeleteRequest,
					DecodeDeleteResponse,
					pb.DeleteResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        

        return &endpoints.Endpoints {
                
			
				StoreEndpoint: storeEndpoint,
			
                
			
				DeleteEndpoint: deleteEndpoint,
			
                
        }
}


	
		func EncodeStoreRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.StoreRequest)
			return req, nil
		}

		func DecodeStoreResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.StoreResponse)
			return response, nil
		}
	

	
		func EncodeDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.DeleteRequest)
			return req, nil
		}

		func DecodeDeleteResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.DeleteResponse)
			return response, nil
		}
	

