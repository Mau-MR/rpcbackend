package service

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

//NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

//Unary returns a server interceptor function to authenticate an authorize unary rpc
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}

}
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("--> unary interceptor: ", info.FullMethod)
		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, stream)

	}

}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		//means everyone can access to that service
		return nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "matadata is not provided")
	}
	//TODO: Check how it looks the metadata
	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is  invalid: %v", err)
	}
	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}
	return status.Error(codes.PermissionDenied, "insuficient credentials")
}
