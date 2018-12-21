package storage_grpctransport



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

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.StorageServiceServer {
     	var options []grpctransport.ServerOption
        _ = options
	return &grpcServer{
		
			
				store: grpctransport.NewServer(
					endpoints.StoreEndpoint,
					decodeRequest,
					encodeStoreResponse,
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
	
		
			store grpctransport.Handler
		
	
		
			delete grpctransport.Handler
		
	
}


	
		func (s *grpcServer) Store(ctx oldcontext.Context, req *pb.StoreRequest) (*pb.StorageObject, error) {
		_, rep, err := s.store.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.StorageObject), nil
		}

		func encodeStoreResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.StorageObject)
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
