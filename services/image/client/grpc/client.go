package image_clientgrpc

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



func New(conn *grpc.ClientConn, logger log.Logger) pb.ImageServiceServer {
        
		
        
		
			var storesyncEndpoint endpoint.Endpoint
			{
				storesyncEndpoint = grpctransport.NewClient(
					conn,
					"image.ImageService",
					"StoreSync",
					EncodeStoreSyncRequest,
					DecodeStoreSyncResponse,
					pb.StoreSyncResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var deleteEndpoint endpoint.Endpoint
			{
				deleteEndpoint = grpctransport.NewClient(
					conn,
					"image.ImageService",
					"Delete",
					EncodeDeleteRequest,
					DecodeDeleteResponse,
					pb.DeleteResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        

        return &endpoints.Endpoints {
                
			
                
			
				StoreSyncEndpoint: storesyncEndpoint,
			
                
			
				DeleteEndpoint: deleteEndpoint,
			
                
        }
}


	

	
		func EncodeStoreSyncRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.StoreSyncRequest)
			return req, nil
		}

		func DecodeStoreSyncResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.StoreSyncResponse)
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
	

