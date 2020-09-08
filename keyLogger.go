package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	fmt.Println("Press ESC to quit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		_, err := fmt.Fprintf(f, "%v", string(event.Rune))
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
