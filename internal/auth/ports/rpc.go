package ports

import (
	"context"
	"login_api/internal/auth/container"
	authv1 "login_api/internal/common/genproto/auth/api/protobuf"

	"connectrpc.com/connect"
)

type AuthServer struct {
	Application container.Application
}

func NewAuthServer(application container.Application) *AuthServer {
	return &AuthServer{Application: application}

}

func (av *AuthServer) SignupWithPhoneNumber(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
	res := connect.NewResponse(&authv1.Response{
		Message: "success",
	})
	return res, nil
}
func (av *AuthServer) VerifyAccount(context.Context, *connect.Request[authv1.OTP]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
func (av *AuthServer) Login(context.Context, *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
func (av *AuthServer) OTPGenerate(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
	res := connect.NewResponse(&authv1.Response{
		Message: "success",
	})
	return res, nil
}
func (av *AuthServer) GetProfile(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
