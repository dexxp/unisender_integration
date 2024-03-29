package main

import (
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.NewConfig()
	application := app.NewApp(cfg)

	fmt.Println("Server starting on: " + cfg.Address)

	go func() {
		err := application.RunHTTPServer()
		if err != nil {
			panic("Не удалось запустить http server")
		}
	}()

	go func() {
		err := application.RunGRPCServer()
		if err != nil {
			panic("Не удалось запустить http server")
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
}
