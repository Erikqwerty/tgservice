package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/erikqwerty/tgservice/internal/config"
	"github.com/erikqwerty/tgservice/internal/service"
	tgapi "github.com/erikqwerty/tgservice/pkg/tgapiv1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "conf/", "Path to configure file")
}

func main() {
	flag.Parse()
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tgService, err := service.NewTgService(conf.TgApikey, conf.ChatID)
	if err != nil {
		log.Fatalf("Ошибка создания сервиса Telegram: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	tgapi.RegisterTgServiceV1Server(s, tgService)

	log.Printf("Server listening at :%v", conf.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Faider to server: ", err)
	}
}
