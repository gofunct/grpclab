package account_httptransport



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



	
		func MakeListHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeListRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.ListAccountsRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeGetByIdHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeGetByIdRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeGetByIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.GetByIdRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeGetByEmailHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeGetByEmailRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeGetByEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.GetByEmailRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeAuthenticateByEmailHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeAuthenticateByEmailRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeAuthenticateByEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.AuthenticateByEmailRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeGeneratePasswordTokenHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeGeneratePasswordTokenRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeGeneratePasswordTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.GeneratePasswordTokenRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeResetPasswordHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeResetPasswordRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeResetPasswordRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.ResetPasswordRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeConfirmAccountHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeConfirmAccountRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeConfirmAccountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.ConfirmAccountRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeCreateHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeCreateRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.CreateAccountRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeUpdateHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeUpdateRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.UpdateAccountRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeDeleteHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeDeleteRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.DeleteAccountRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(svc pb.AccountServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
		
			log.Println("new HTTP endpoint: \"/List\" (service=Account)")
			mux.Handle("/List", MakeListHandler(svc, endpoints.ListEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/GetById\" (service=Account)")
			mux.Handle("/GetById", MakeGetByIdHandler(svc, endpoints.GetByIdEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/GetByEmail\" (service=Account)")
			mux.Handle("/GetByEmail", MakeGetByEmailHandler(svc, endpoints.GetByEmailEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/AuthenticateByEmail\" (service=Account)")
			mux.Handle("/AuthenticateByEmail", MakeAuthenticateByEmailHandler(svc, endpoints.AuthenticateByEmailEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/GeneratePasswordToken\" (service=Account)")
			mux.Handle("/GeneratePasswordToken", MakeGeneratePasswordTokenHandler(svc, endpoints.GeneratePasswordTokenEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/ResetPassword\" (service=Account)")
			mux.Handle("/ResetPassword", MakeResetPasswordHandler(svc, endpoints.ResetPasswordEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/ConfirmAccount\" (service=Account)")
			mux.Handle("/ConfirmAccount", MakeConfirmAccountHandler(svc, endpoints.ConfirmAccountEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Create\" (service=Account)")
			mux.Handle("/Create", MakeCreateHandler(svc, endpoints.CreateEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Update\" (service=Account)")
			mux.Handle("/Update", MakeUpdateHandler(svc, endpoints.UpdateEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Delete\" (service=Account)")
			mux.Handle("/Delete", MakeDeleteHandler(svc, endpoints.DeleteEndpoint))
		
	
	return nil
}
package account_httptransport



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



	
		func MakeAuthenticateHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeAuthenticateRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeAuthenticateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.AuthRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	

	
		func MakeValidateHandler(svc pb.AccountServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
			return httptransport.NewServer(
				endpoint,
				decodeValidateRequest,
				encodeResponse,
				[]httptransport.ServerOption{}...,
			)
		}

		func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
			var req pb.ValidateRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			return &req, nil
		}
	


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(svc pb.AccountServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
		
			log.Println("new HTTP endpoint: \"/Authenticate\" (service=Account)")
			mux.Handle("/Authenticate", MakeAuthenticateHandler(svc, endpoints.AuthenticateEndpoint))
		
	
		
			log.Println("new HTTP endpoint: \"/Validate\" (service=Account)")
			mux.Handle("/Validate", MakeValidateHandler(svc, endpoints.ValidateEndpoint))
		
	
	return nil
}
