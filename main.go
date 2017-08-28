package main

import (
	"fmt"
	"os"
)

func main() {
	msg := os.Getenv("QIITA_ACCESS_TOKEN")
	fmt.Println(msg)
}

func getQiitaItems() {

}
