package initializer

import (
	"github.com/akshaybt001/DatingApp_NotificationService/internal/adapters"
	"github.com/akshaybt001/DatingApp_NotificationService/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func Initializer(db *mongo.Database) *service.NotificationService{
	adapter:=adapters.NewEmailAdapter(db)
	service:=service.NewEmailService(adapter)
	return service
}