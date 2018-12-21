package account_grpctransport



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

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.AccountServiceServer {
     	var options []grpctransport.ServerOption
        _ = options
	return &grpcServer{
		
			
				list: grpctransport.NewServer(
					endpoints.ListEndpoint,
					decodeRequest,
					encodeListResponse,
					options...,
				),
			
                
			
				getbyid: grpctransport.NewServer(
					endpoints.GetByIdEndpoint,
					decodeRequest,
					encodeGetByIdResponse,
					options...,
				),
			
                
			
				getbyemail: grpctransport.NewServer(
					endpoints.GetByEmailEndpoint,
					decodeRequest,
					encodeGetByEmailResponse,
					options...,
				),
			
                
			
				authenticatebyemail: grpctransport.NewServer(
					endpoints.AuthenticateByEmailEndpoint,
					decodeRequest,
					encodeAuthenticateByEmailResponse,
					options...,
				),
			
                
			
				generatepasswordtoken: grpctransport.NewServer(
					endpoints.GeneratePasswordTokenEndpoint,
					decodeRequest,
					encodeGeneratePasswordTokenResponse,
					options...,
				),
			
                
			
				resetpassword: grpctransport.NewServer(
					endpoints.ResetPasswordEndpoint,
					decodeRequest,
					encodeResetPasswordResponse,
					options...,
				),
			
                
			
				confirmaccount: grpctransport.NewServer(
					endpoints.ConfirmAccountEndpoint,
					decodeRequest,
					encodeConfirmAccountResponse,
					options...,
				),
			
                
			
				create: grpctransport.NewServer(
					endpoints.CreateEndpoint,
					decodeRequest,
					encodeCreateResponse,
					options...,
				),
			
                
			
				update: grpctransport.NewServer(
					endpoints.UpdateEndpoint,
					decodeRequest,
					encodeUpdateResponse,
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
	
		
			list grpctransport.Handler
		
	
		
			getbyid grpctransport.Handler
		
	
		
			getbyemail grpctransport.Handler
		
	
		
			authenticatebyemail grpctransport.Handler
		
	
		
			generatepasswordtoken grpctransport.Handler
		
	
		
			resetpassword grpctransport.Handler
		
	
		
			confirmaccount grpctransport.Handler
		
	
		
			create grpctransport.Handler
		
	
		
			update grpctransport.Handler
		
	
		
			delete grpctransport.Handler
		
	
}


	
		func (s *grpcServer) List(ctx oldcontext.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
		_, rep, err := s.list.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.ListAccountsResponse), nil
		}

		func encodeListResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.ListAccountsResponse)
			return resp, nil
		}
	

	
		func (s *grpcServer) GetById(ctx oldcontext.Context, req *pb.GetByIdRequest) (*pb.Account, error) {
		_, rep, err := s.getbyid.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeGetByIdResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) GetByEmail(ctx oldcontext.Context, req *pb.GetByEmailRequest) (*pb.Account, error) {
		_, rep, err := s.getbyemail.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeGetByEmailResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) AuthenticateByEmail(ctx oldcontext.Context, req *pb.AuthenticateByEmailRequest) (*pb.Account, error) {
		_, rep, err := s.authenticatebyemail.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeAuthenticateByEmailResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) GeneratePasswordToken(ctx oldcontext.Context, req *pb.GeneratePasswordTokenRequest) (*pb.GeneratePasswordTokenResponse, error) {
		_, rep, err := s.generatepasswordtoken.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.GeneratePasswordTokenResponse), nil
		}

		func encodeGeneratePasswordTokenResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.GeneratePasswordTokenResponse)
			return resp, nil
		}
	

	
		func (s *grpcServer) ResetPassword(ctx oldcontext.Context, req *pb.ResetPasswordRequest) (*pb.Account, error) {
		_, rep, err := s.resetpassword.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeResetPasswordResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) ConfirmAccount(ctx oldcontext.Context, req *pb.ConfirmAccountRequest) (*pb.Account, error) {
		_, rep, err := s.confirmaccount.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeConfirmAccountResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) Create(ctx oldcontext.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {
		_, rep, err := s.create.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeCreateResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) Update(ctx oldcontext.Context, req *pb.UpdateAccountRequest) (*pb.Account, error) {
		_, rep, err := s.update.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Account), nil
		}

		func encodeUpdateResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Account)
			return resp, nil
		}
	

	
		func (s *grpcServer) Delete(ctx oldcontext.Context, req *pb.DeleteAccountRequest) (*pb.Empty, error) {
		_, rep, err := s.delete.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Empty), nil
		}

		func encodeDeleteResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Empty)
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
package account_grpctransport



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

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.AccountServiceServer {
     	var options []grpctransport.ServerOption
        _ = options
	return &grpcServer{
		
			
				authenticate: grpctransport.NewServer(
					endpoints.AuthenticateEndpoint,
					decodeRequest,
					encodeAuthenticateResponse,
					options...,
				),
			
                
			
				validate: grpctransport.NewServer(
					endpoints.ValidateEndpoint,
					decodeRequest,
					encodeValidateResponse,
					options...,
				),
			
                
	}
}

type grpcServer struct {
	
		
			authenticate grpctransport.Handler
		
	
		
			validate grpctransport.Handler
		
	
}


	
		func (s *grpcServer) Authenticate(ctx oldcontext.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
		_, rep, err := s.authenticate.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.AuthResponse), nil
		}

		func encodeAuthenticateResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.AuthResponse)
			return resp, nil
		}
	

	
		func (s *grpcServer) Validate(ctx oldcontext.Context, req *pb.ValidateRequest) (*pb.Empty, error) {
		_, rep, err := s.validate.ServeGRPC(ctx, req)
			if err != nil {
				return nil, err
			}
			return rep.(*pb.Empty), nil
		}

		func encodeValidateResponse(ctx context.Context, response interface{}) (interface{}, error) {
			resp := response.(*pb.Empty)
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
