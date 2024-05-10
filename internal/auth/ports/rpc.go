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

func (av *AuthServer) SignupWithPhoneNumber(ctx context.Context, re *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.SignUpResponse], error) {
	dto_obj, err := av.Application.AuthApplication.SignUp(ctx, re.Msg.Number)
	if err != nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, err)
	}
	res := connect.NewResponse(&authv1.SignUpResponse{
		UserId:    dto_obj.UserId,
		IsVerfied: dto_obj.IsVerified,
		CreatedAt: dto_obj.CreatedAt,
	})
	return res, nil
}
func (av *AuthServer) VerifyAccount(ctx context.Context, re *connect.Request[authv1.VerifyAccountRequest]) (*connect.Response[authv1.VerifyAccountResponse], error) {
	dto_obj, err := av.Application.AuthApplication.VerifyAccount(ctx, re.Msg.UserId, re.Msg.Code)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	res := connect.NewResponse(&authv1.VerifyAccountResponse{
		Message: dto_obj.Message,
	})
	return res, nil
}
func (av *AuthServer) Login(context.Context, *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
func (av *AuthServer) OTPGenerate(ctx context.Context, re *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
	dto_obj, err := av.Application.AuthApplication.GenerateOTP(ctx, re.Msg.Number)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	res := connect.NewResponse(&authv1.Response{
		Message: dto_obj.Message,
	})
	return res, nil
}
func (av *AuthServer) GetProfile(context.Context, *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.UserProfile], error) {
	res := connect.NewResponse(&authv1.UserProfile{
		PhoneNumber: "9677892850",
	})
	return res, nil
}
