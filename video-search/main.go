package main

import (
	"log"
	"context"
	"time"
	"fmt"
	"github.com/ktamashun/flickrphotosearch/video-search/app"
)

func main() {
	application := app.NewApp()

	if err := application.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	sig := <-application.Done()
	fmt.Println(fmt.Sprintf("Captured signal: %s", sig))

	stopCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := application.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
