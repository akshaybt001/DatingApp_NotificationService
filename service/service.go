package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akshaybt001/DatingApp_NotificationService/internal/adapters"
	"github.com/akshaybt001/DatingApp_NotificationService/internal/helper/otp"
	"github.com/akshaybt001/DatingApp_proto_files/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (email *NotificationService) AddNotification(ctx context.Context,req *pb.AddNotificationRequest) (*pb.NoMessage,error){
	if req.UserId==""{
		return nil,fmt.Errorf("please provide a valid userID")

	}
	var message primitive.M
	if err:=json.Unmarshal([]byte(req.Message),&message);err!=nil{
		return nil,fmt.Errorf("failed to parse message JSON : %v",err)
	}
	if err:=email.adapter.AddNotification(req.UserId,message);err!=nil{
		return nil,err
	}
	return nil,nil
}

func (email *NotificationService) GetAllNotifications(req *pb.GetNotificationsByUserId,srv pb.Notification_GetAllNotificationsServer)error{
	notifications,err:=email.adapter.GetAllNotifications(req.UserId)
	if err!=nil{
		return err
	}
	for _,notification:=range notifications{
		message,ok:=notification["message"].(string)
		if !ok{
			return fmt.Errorf("message field is not a string in notification: %v",notification)
		}
		seen,ok:=notification["seen"].(bool)
		if !ok{
			return fmt.Errorf("unable to get seen status")
		}
		res:=&pb.NotificationResponse{
			Message: message,
			Seen: seen,
		}
		if err:=srv.Send(res);err!=nil{
			return err
		}
	}
	return nil
}