package storage_endpoints



import (
	context "context"
	"fmt"
	oldcontext "golang.org/x/net/context"
    pb "github.com/gofunct/services"
	"github.com/go-kit/kit/endpoint"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	
		
			StoreEndpoint endpoint.Endpoint
		
	
		
			DeleteEndpoint endpoint.Endpoint
		
	
}


	
		
			func (e *Endpoints)Store(ctx oldcontext.Context, in *pb.StoreRequest) (*pb.StorageObject, error) {
				out, err := e.StoreEndpoint(ctx, in)
				if err != nil {
					return &pb.StorageObject{ErrMsg: err.Error()}, err
				}
				return out.(*pb.StorageObject), err
			}
		
	

	
		
			func (e *Endpoints)Delete(ctx oldcontext.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
				out, err := e.DeleteEndpoint(ctx, in)
				if err != nil {
					return &pb.DeleteResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.DeleteResponse), err
			}
		
	



	
		func MakeStoreEndpoint(svc pb.StorageServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.StoreRequest)
				rep, err := svc.Store(ctx, req)
				if err != nil {
					return &pb.StorageObject{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeDeleteEndpoint(svc pb.StorageServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.DeleteRequest)
				rep, err := svc.Delete(ctx, req)
				if err != nil {
					return &pb.DeleteResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	


func MakeEndpoints(svc pb.StorageServiceServer) Endpoints {
	return Endpoints{
		
			StoreEndpoint: MakeStoreEndpoint(svc),
		
			DeleteEndpoint: MakeDeleteEndpoint(svc),
		
	}
}
