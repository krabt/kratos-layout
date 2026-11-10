package service

import (
	"context"
	"fmt"
	v1 "kratos-layout/api/hello/v1"
	"kratos-layout/internal/biz"
)

type HelloService struct {
	v1.UnimplementedHelloServiceServer
	uc *biz.UserBiz
}

func NewHelloService(uc *biz.UserBiz) *HelloService {
	return &HelloService{
		uc: uc,
	}
}

func (s *HelloService) Hello(ctx context.Context, req *v1.HelloReq) (*v1.HelloResp, error) {
	_, _ = s.uc.GetUserByEmail(ctx, req.Name)
	// if err != nil {
	// 	return nil, err
	// }
	return &v1.HelloResp{
		Message: fmt.Sprintf("Hello, %s!", "kkkk"),
	}, nil
}
