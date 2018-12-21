package account_endpoints



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
	
		
			ListEndpoint endpoint.Endpoint
		
	
		
			GetByIdEndpoint endpoint.Endpoint
		
	
		
			GetByEmailEndpoint endpoint.Endpoint
		
	
		
			AuthenticateByEmailEndpoint endpoint.Endpoint
		
	
		
			GeneratePasswordTokenEndpoint endpoint.Endpoint
		
	
		
			ResetPasswordEndpoint endpoint.Endpoint
		
	
		
			ConfirmAccountEndpoint endpoint.Endpoint
		
	
		
			CreateEndpoint endpoint.Endpoint
		
	
		
			UpdateEndpoint endpoint.Endpoint
		
	
		
			DeleteEndpoint endpoint.Endpoint
		
	
}


	
		
			func (e *Endpoints)List(ctx oldcontext.Context, in *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
				out, err := e.ListEndpoint(ctx, in)
				if err != nil {
					return &pb.ListAccountsResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.ListAccountsResponse), err
			}
		
	

	
		
			func (e *Endpoints)GetById(ctx oldcontext.Context, in *pb.GetByIdRequest) (*pb.Account, error) {
				out, err := e.GetByIdEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)GetByEmail(ctx oldcontext.Context, in *pb.GetByEmailRequest) (*pb.Account, error) {
				out, err := e.GetByEmailEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)AuthenticateByEmail(ctx oldcontext.Context, in *pb.AuthenticateByEmailRequest) (*pb.Account, error) {
				out, err := e.AuthenticateByEmailEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)GeneratePasswordToken(ctx oldcontext.Context, in *pb.GeneratePasswordTokenRequest) (*pb.GeneratePasswordTokenResponse, error) {
				out, err := e.GeneratePasswordTokenEndpoint(ctx, in)
				if err != nil {
					return &pb.GeneratePasswordTokenResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.GeneratePasswordTokenResponse), err
			}
		
	

	
		
			func (e *Endpoints)ResetPassword(ctx oldcontext.Context, in *pb.ResetPasswordRequest) (*pb.Account, error) {
				out, err := e.ResetPasswordEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)ConfirmAccount(ctx oldcontext.Context, in *pb.ConfirmAccountRequest) (*pb.Account, error) {
				out, err := e.ConfirmAccountEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)Create(ctx oldcontext.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
				out, err := e.CreateEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)Update(ctx oldcontext.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
				out, err := e.UpdateEndpoint(ctx, in)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Account), err
			}
		
	

	
		
			func (e *Endpoints)Delete(ctx oldcontext.Context, in *pb.DeleteAccountRequest) (*pb.Empty, error) {
				out, err := e.DeleteEndpoint(ctx, in)
				if err != nil {
					return &pb.Empty{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Empty), err
			}
		
	



	
		func MakeListEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.ListAccountsRequest)
				rep, err := svc.List(ctx, req)
				if err != nil {
					return &pb.ListAccountsResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeGetByIdEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.GetByIdRequest)
				rep, err := svc.GetById(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeGetByEmailEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.GetByEmailRequest)
				rep, err := svc.GetByEmail(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeAuthenticateByEmailEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.AuthenticateByEmailRequest)
				rep, err := svc.AuthenticateByEmail(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeGeneratePasswordTokenEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.GeneratePasswordTokenRequest)
				rep, err := svc.GeneratePasswordToken(ctx, req)
				if err != nil {
					return &pb.GeneratePasswordTokenResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeResetPasswordEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.ResetPasswordRequest)
				rep, err := svc.ResetPassword(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeConfirmAccountEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.ConfirmAccountRequest)
				rep, err := svc.ConfirmAccount(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeCreateEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.CreateAccountRequest)
				rep, err := svc.Create(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeUpdateEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.UpdateAccountRequest)
				rep, err := svc.Update(ctx, req)
				if err != nil {
					return &pb.Account{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeDeleteEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.DeleteAccountRequest)
				rep, err := svc.Delete(ctx, req)
				if err != nil {
					return &pb.Empty{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	


func MakeEndpoints(svc pb.AccountServiceServer) Endpoints {
	return Endpoints{
		
			ListEndpoint: MakeListEndpoint(svc),
		
			GetByIdEndpoint: MakeGetByIdEndpoint(svc),
		
			GetByEmailEndpoint: MakeGetByEmailEndpoint(svc),
		
			AuthenticateByEmailEndpoint: MakeAuthenticateByEmailEndpoint(svc),
		
			GeneratePasswordTokenEndpoint: MakeGeneratePasswordTokenEndpoint(svc),
		
			ResetPasswordEndpoint: MakeResetPasswordEndpoint(svc),
		
			ConfirmAccountEndpoint: MakeConfirmAccountEndpoint(svc),
		
			CreateEndpoint: MakeCreateEndpoint(svc),
		
			UpdateEndpoint: MakeUpdateEndpoint(svc),
		
			DeleteEndpoint: MakeDeleteEndpoint(svc),
		
	}
}
package account_endpoints



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
	
		
			AuthenticateEndpoint endpoint.Endpoint
		
	
		
			ValidateEndpoint endpoint.Endpoint
		
	
}


	
		
			func (e *Endpoints)Authenticate(ctx oldcontext.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
				out, err := e.AuthenticateEndpoint(ctx, in)
				if err != nil {
					return &pb.AuthResponse{ErrMsg: err.Error()}, err
				}
				return out.(*pb.AuthResponse), err
			}
		
	

	
		
			func (e *Endpoints)Validate(ctx oldcontext.Context, in *pb.ValidateRequest) (*pb.Empty, error) {
				out, err := e.ValidateEndpoint(ctx, in)
				if err != nil {
					return &pb.Empty{ErrMsg: err.Error()}, err
				}
				return out.(*pb.Empty), err
			}
		
	



	
		func MakeAuthenticateEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.AuthRequest)
				rep, err := svc.Authenticate(ctx, req)
				if err != nil {
					return &pb.AuthResponse{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	

	
		func MakeValidateEndpoint(svc pb.AccountServiceServer) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*pb.ValidateRequest)
				rep, err := svc.Validate(ctx, req)
				if err != nil {
					return &pb.Empty{ErrMsg: err.Error()}, err
				}
				return rep, nil
			}
		}
	


func MakeEndpoints(svc pb.AccountServiceServer) Endpoints {
	return Endpoints{
		
			AuthenticateEndpoint: MakeAuthenticateEndpoint(svc),
		
			ValidateEndpoint: MakeValidateEndpoint(svc),
		
	}
}
