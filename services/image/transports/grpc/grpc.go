package image_grpctransport



import (
	context "context"
        "fmt"

    oldcontext "golang.org/x/net/context"
	grpctransport "github.com/go-kit/kit/transport/grpc"

    pb "github.com/gofunct/services"
    endpoints "github.com/gofunct/services/endpoints"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.ImageServiceServer {
     	var options []grpctransport.ServerOption
        _ = options
	return &grpcServer{
		
			
				store: &server{
					e: endpoints.StoreEndpoint,
				},
			
                
			
				storesync: grpctransport.NewServer(
					endpoints.StoreSyncEndpoint,
					decodeRequest,
					encodeStoreSyncResponse,
					options...,
				),
			
                
			
				delete: grpctransport.NewServer(
					endpoints.DeleteEndpoint,
					decodeRequest,
					encodeDeleteResponse,
					options...,
				),
			
                
	}
}

type grpcServer struct {
	
		
			store streamHandler
		
	
		
			storesync grpctransport.Handler
		
	
		
			delete grpctransport.Handler
		
	
}


	
		func (s *grpcServer) Store(req *pb.StoreRequest, server pb.ImageService_StoreServer) error {
		        return s.store.Do(server, req)
		}
	

	
		func (s *grpcServer) StoreSync(ctx oldcontext.Context, req *pb.ImageStoreRequest) (*pb.ImageSyncResponse, error) {
		_, rep, err := s.storesync.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.ImageSyncResponse), nil
		}

		func encodeStoreSyncResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.ImageSyncResponse)
			return resp, nil
		}
	

	
		func (s *grpcServer) Delete(ctx oldcontext.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
		_, rep, err := s.delete.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.DeleteResponse), nil
		}

		func encodeDeleteResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.DeleteResponse)
			return resp, nil
		}
	


func decodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

type streamHandler interface{
	Do(server interface{}, req interface{}) (err error)
}

type server struct {
	e endpoints.StreamEndpoint
}

func (s server) Do(server interface{}, req interface{}) (err error) {
	if err := s.e(server, req); err != nil {
		return err
	}
	return nil
}
