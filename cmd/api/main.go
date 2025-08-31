package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ghady-Abboud/paper-trading.git/internal/server"
)

func main() {
	server.RestyClientInit()
	router := server.RegisterRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go server.HandleAlpacaWs(ctx)

	<-ctx.Done()
	log.Println("Shutdown Server...")
	ctxTimeout, cancelTimeOut := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeOut()
	if err := srv.Shutdown(ctxTimeout); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
