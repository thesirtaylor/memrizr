package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thesirtaylor/memrizr/handler"
)

func main()  {
	log.Println("Starting server...")

	router := gin.Default();

	handler.NewHandler(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	//Graceful server shutdown
	go func() {
		//this is called in a dedicated goroutine because it is a blocking operation
		//this means that if srv.ListenAndServe() is called in the main goroutine, the program will be stuck at this point
		//and won't execute any code after it until the server is shutdown
		err := srv.ListenAndServe();
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf(("Failed to initialize server: %v"), err)
		}
	}()

	log.Printf("Listening on port %s", srv.Addr);

	//create a dedicated channel to listen for OS signals
	quit := make(chan os.Signal, 1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	//the 2 lines above do the same thing as process.on('SIGINT', callback) and process.on('SIGTERM', callback) in node.js

	//what is happening here is that the code creates a seperate channel called `quit`, 
	//and then tells the channel to listen for SIGINT and SIGTERM signals on the OS

	//this blocks the execution of the current goroutine until a signal is passed into the quit channel
	//that is, when the program receives either a SIGINT or SIGTERM signal from the OS, a signal would be sent on the quit channel

	//when the next line is reached in the code, the program will block(pause) at this line until a value is sent on the quit channel
	//once one of the signals (SIGTERM or SIGINT) is received, the program will continue executing the next line of code
	//the purpose of <-quit is to wait for and handle termination signals from the OS
	<-quit;
	//so if there is no 'SIGTERM' or 'SIGINT' signal from the OS, the '<-quit' operation will continue to block the execution of the program indefinitely
	//once a signal is received, the program will continue executing the next line of code
	//the context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second);
	defer cancel();//ensures resource cleanup

	log.Println("Shutting down server...");
	err := srv.Shutdown(ctx);
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err);
	}
}