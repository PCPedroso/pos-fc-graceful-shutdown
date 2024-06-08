package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 4)
		w.Write([]byte("Olá mundo!"))
	})

	go func() {
		fmt.Println("Servidor rodando em 'http://localhost:8080'")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Não foi possível subir o servidor web %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("Parando o servidor...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Não foi possível parar o servidor de forma segura: %v\n", err)
	}
	fmt.Println("Servidor parado")
}
