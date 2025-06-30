package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/akshaykathwate/students_api/internal/config"
	"github.com/akshaykathwate/students_api/internal/config/http/handlers/student"
)

func main() {

	// load config 
	cfg:=config.MustLoad()

	// database Setup
	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New())

	server := http.Server{
		Addr: cfg.Httpserver.Addr,
		Handler: router,
	}

	fmt.Println("Server Started At ... %s",cfg.Httpserver.Addr)

	done := make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt,syscall.SIGINT, syscall.SIGTERM)

	go func(){
		err:= server.ListenAndServe();
		if err!=nil{
			log.Fatalf("Falied to Start Server")
		}
	}()

	<- done
	
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}