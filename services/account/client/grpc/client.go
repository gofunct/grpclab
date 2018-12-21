package account_clientgrpc

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



func New(conn *grpc.ClientConn, logger log.Logger) pb.AccountServiceServer {
        
		
			var listEndpoint endpoint.Endpoint
			{
				listEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"List",
					EncodeListRequest,
					DecodeListResponse,
					pb.ListResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var getbyidEndpoint endpoint.Endpoint
			{
				getbyidEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"GetById",
					EncodeGetByIdRequest,
					DecodeGetByIdResponse,
					pb.GetByIdResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var getbyemailEndpoint endpoint.Endpoint
			{
				getbyemailEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"GetByEmail",
					EncodeGetByEmailRequest,
					DecodeGetByEmailResponse,
					pb.GetByEmailResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var authenticatebyemailEndpoint endpoint.Endpoint
			{
				authenticatebyemailEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"AuthenticateByEmail",
					EncodeAuthenticateByEmailRequest,
					DecodeAuthenticateByEmailResponse,
					pb.AuthenticateByEmailResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var generatepasswordtokenEndpoint endpoint.Endpoint
			{
				generatepasswordtokenEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"GeneratePasswordToken",
					EncodeGeneratePasswordTokenRequest,
					DecodeGeneratePasswordTokenResponse,
					pb.GeneratePasswordTokenResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var resetpasswordEndpoint endpoint.Endpoint
			{
				resetpasswordEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"ResetPassword",
					EncodeResetPasswordRequest,
					DecodeResetPasswordResponse,
					pb.ResetPasswordResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var confirmaccountEndpoint endpoint.Endpoint
			{
				confirmaccountEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"ConfirmAccount",
					EncodeConfirmAccountRequest,
					DecodeConfirmAccountResponse,
					pb.ConfirmAccountResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var createEndpoint endpoint.Endpoint
			{
				createEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"Create",
					EncodeCreateRequest,
					DecodeCreateResponse,
					pb.CreateResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var updateEndpoint endpoint.Endpoint
			{
				updateEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"Update",
					EncodeUpdateRequest,
					DecodeUpdateResponse,
					pb.UpdateResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var deleteEndpoint endpoint.Endpoint
			{
				deleteEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"Delete",
					EncodeDeleteRequest,
					DecodeDeleteResponse,
					pb.DeleteResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        

        return &endpoints.Endpoints {
                
			
				ListEndpoint: listEndpoint,
			
                
			
				GetByIdEndpoint: getbyidEndpoint,
			
                
			
				GetByEmailEndpoint: getbyemailEndpoint,
			
                
			
				AuthenticateByEmailEndpoint: authenticatebyemailEndpoint,
			
                
			
				GeneratePasswordTokenEndpoint: generatepasswordtokenEndpoint,
			
                
			
				ResetPasswordEndpoint: resetpasswordEndpoint,
			
                
			
				ConfirmAccountEndpoint: confirmaccountEndpoint,
			
                
			
				CreateEndpoint: createEndpoint,
			
                
			
				UpdateEndpoint: updateEndpoint,
			
                
			
				DeleteEndpoint: deleteEndpoint,
			
                
        }
}


	
		func EncodeListRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.ListRequest)
			return req, nil
		}

		func DecodeListResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.ListResponse)
			return response, nil
		}
	

	
		func EncodeGetByIdRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.GetByIdRequest)
			return req, nil
		}

		func DecodeGetByIdResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.GetByIdResponse)
			return response, nil
		}
	

	
		func EncodeGetByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.GetByEmailRequest)
			return req, nil
		}

		func DecodeGetByEmailResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.GetByEmailResponse)
			return response, nil
		}
	

	
		func EncodeAuthenticateByEmailRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.AuthenticateByEmailRequest)
			return req, nil
		}

		func DecodeAuthenticateByEmailResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.AuthenticateByEmailResponse)
			return response, nil
		}
	

	
		func EncodeGeneratePasswordTokenRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.GeneratePasswordTokenRequest)
			return req, nil
		}

		func DecodeGeneratePasswordTokenResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.GeneratePasswordTokenResponse)
			return response, nil
		}
	

	
		func EncodeResetPasswordRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.ResetPasswordRequest)
			return req, nil
		}

		func DecodeResetPasswordResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.ResetPasswordResponse)
			return response, nil
		}
	

	
		func EncodeConfirmAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.ConfirmAccountRequest)
			return req, nil
		}

		func DecodeConfirmAccountResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.ConfirmAccountResponse)
			return response, nil
		}
	

	
		func EncodeCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.CreateRequest)
			return req, nil
		}

		func DecodeCreateResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.CreateResponse)
			return response, nil
		}
	

	
		func EncodeUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.UpdateRequest)
			return req, nil
		}

		func DecodeUpdateResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.UpdateResponse)
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
	

package account_clientgrpc

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



func New(conn *grpc.ClientConn, logger log.Logger) pb.AccountServiceServer {
        
		
			var authenticateEndpoint endpoint.Endpoint
			{
				authenticateEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"Authenticate",
					EncodeAuthenticateRequest,
					DecodeAuthenticateResponse,
					pb.AuthenticateResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        
		
			var validateEndpoint endpoint.Endpoint
			{
				validateEndpoint = grpctransport.NewClient(
					conn,
					"account.AccountService",
					"Validate",
					EncodeValidateRequest,
					DecodeValidateResponse,
					pb.ValidateResponse{},
					append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.ContextToGRPC()))...,
				).Endpoint()
			}
		
        

        return &endpoints.Endpoints {
                
			
				AuthenticateEndpoint: authenticateEndpoint,
			
                
			
				ValidateEndpoint: validateEndpoint,
			
                
        }
}


	
		func EncodeAuthenticateRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.AuthenticateRequest)
			return req, nil
		}

		func DecodeAuthenticateResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.AuthenticateResponse)
			return response, nil
		}
	

	
		func EncodeValidateRequest(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(*pb.ValidateRequest)
			return req, nil
		}

		func DecodeValidateResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
			response := grpcResponse.(*pb.ValidateResponse)
			return response, nil
		}
	

