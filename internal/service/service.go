package service

import (
	"context"
	"fmt"
	"log"

	tgapi "github.com/erikqwerty/tgservice/pkg/tgapiv1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Tg struct {
	tgapi.UnimplementedTgServiceV1Server
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func NewTgService(botToken string, chatID int64) (*Tg, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации бота: %w", err)
	}
	return &Tg{Bot: bot, ChatID: chatID}, nil
}

func (tg *Tg) SendMessage(ctx context.Context, req *tgapi.SendMessageRequest) (*emptypb.Empty, error) {
	msg := tgbotapi.NewMessage(tg.ChatID, req.Message)
	if _, err := tg.Bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
		return nil, err
	}
	log.Printf("Сообщение отправлено: %s", req.Message)
	return &emptypb.Empty{}, nil
}
