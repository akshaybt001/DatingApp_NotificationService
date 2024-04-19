package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/akshaybt001/DatingApp_NotificationService/db"
	"github.com/akshaybt001/DatingApp_NotificationService/initializer"
	"github.com/akshaybt001/DatingApp_proto_files/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}
	addr := os.Getenv("DB_KEY")
	DB, err := db.InitMongoDB(addr)
	if err != nil {
		log.Fatal("error connecting to database")
	}
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("failed to listen on port 8083")
	}
	fmt.Println("notification service  listening on port 8083")
	services := initializer.Initializer(DB)
	server := grpc.NewServer()
	pb.RegisterNotificationServer(server, services)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to listen on port 8083")
	}
}
