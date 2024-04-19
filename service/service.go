package service

import (
	"context"

	"github.com/akshaybt001/DatingApp_NotificationService/internal/adapters"
	"github.com/akshaybt001/DatingApp_NotificationService/internal/helper/otp"
	"github.com/akshaybt001/DatingApp_proto_files/pb"
)

type NotificationService struct {
	adapter adapters.EmailInterface
	pb.UnimplementedNotificationServer
}

func NewEmailService(adapter adapters.EmailInterface)*NotificationService{
	return &NotificationService{
		adapter: adapter,
	}
}

func (email *NotificationService) SendOTP(ctx context.Context,req *pb.SendOtpRequest) (*pb.NoMessage,error){
	otp.SendOTP(req.Email)
	return &pb.NoMessage{},nil
}

func (email *NotificationService) VerifyOTP(ct context.Context,req *pb.VerifyOtpRequest)(*pb.VerifyOtpResponse,error){
	verified:=otp.VerifyOTP(req.Email,req.Otp)
	res:=&pb.VerifyOtpResponse{
		Verified: verified,
	}
	return res,nil
}