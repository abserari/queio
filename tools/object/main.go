package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.design/x/tgstore"
	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	tgs := tgstore.New()

	tgs.BotToken = os.Getenv("TELEGRAM_TECHCATS_BOT_TOKEN")
	if tgs.BotToken == "" {
		fmt.Println("no bot token")
		return
	}
	tgs.ChatID = -1001179220569

	objectKey := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(objectKey); err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()

	_, err := tgs.Upload(
		context.TODO(),
		objectKey,
		strings.NewReader("Hello, 世界"),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upload time:", time.Since(startTime))

	startTime = time.Now()

	// downloadedObject, err := tgs.Download(
	// 	context.TODO(),
	// 	object.ID,
	// 	objectKey,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Download time:", time.Since(startTime))

	// startTime = time.Now()

	// rc, err := downloadedObject.NewReader(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rc.Close()

	// b, err := ioutil.ReadAll(rc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Read time:", time.Since(startTime))

	// fmt.Println("Content:", string(b))
}
