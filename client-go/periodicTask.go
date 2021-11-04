package main

import (
	"fmt"
	"time"
)

func main() {
	listTicker := time.NewTicker(1 *time.Second)

	go func() {
		for {
			select {
			case <-listTicker.C:
                 fmt.Println("ticker called!")
				}
			}
		}()
	time.Sleep(5 * time.Second)

}
