package image_endpoints



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
	
		
			StoreEndpoint StreamEndpoint
		
	
		
			StoreSyncEndpoint endpoint.Endpoint
		
	
		
			DeleteEndpoint endpoint.Endpoint
		
	
}


	
		
			func (e *Endpoints)Store(in *pb.StoreRequest, server pb.ImageService_StoreServer) error {
				return fmt.Errorf("not implemented")
			}
		
	

	
		
			func (e *Endpoints)StoreSync(ctx oldcontext.Context, in *pb.ImageStoreRequest) (*pb.ImageSyncResponse, error) {
				out, err := e.StoreSyncEndpoint(ctx, in)
				if err != nil {
					return &pb.ImageSyncResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.ImageSyncResponse), err
			}
		
	

	
		
			func (e *Endpoints)Delete(ctx oldcontext.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
				out, err := e.DeleteEndpoint(ctx, in)
				if err != nil {
					return &pb.DeleteResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.DeleteResponse), err
			}
		
	



	
		func MakeStoreEndpoint(svc pb.ImageServiceServer) StreamEndpoint {
			return func(server interface{}, request interface{}) error {
				
				return svc.Store(request.(*pb.StoreRequest), server.(pb.ImageService_StoreServer))
				
			}
		}
	

	
		func MakeStoreSyncEndpoint(svc pb.ImageServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.ImageStoreRequest)
				rep, err := svc.StoreSync(ctx, req)
				if err != nil {
					return &pb.ImageSyncResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeDeleteEndpoint(svc pb.ImageServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.DeleteRequest)
				rep, err := svc.Delete(ctx, req)
				if err != nil {
					return &pb.DeleteResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	


func MakeEndpoints(svc pb.ImageServiceServer) Endpoints {
	return Endpoints{
		
			StoreEndpoint: MakeStoreEndpoint(svc),
		
			StoreSyncEndpoint: MakeStoreSyncEndpoint(svc),
		
			DeleteEndpoint: MakeDeleteEndpoint(svc),
		
	}
}
