package grpcapi

import (
	"context"
	"fmt"

	"github.com/fingo-martPedia/fingo-ums/cmd/proto/tokenvalidation"
	"github.com/fingo-martPedia/fingo-ums/constants"
	"github.com/fingo-martPedia/fingo-ums/helpers"
	"github.com/fingo-martPedia/fingo-ums/internal/interfaces/grpciface"
	"github.com/fingo-martPedia/fingo-ums/internal/services/grpcservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenValidationHandler struct {
	TokenValidationService grpciface.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServer
}

func NewTokenValidationHandler(svc *grpcservice.TokenValidationService) *TokenValidationHandler {
	return &TokenValidationHandler{TokenValidationService: svc}
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, status.Error(codes.InvalidArgument, "token is required")
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, status.Error(codes.Unauthenticated, "invalid or expired token")
	}

	return &tokenvalidation.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
			Email:    claimToken.Email,
		},
	}, nil
}
