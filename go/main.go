package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("path:", os.Getenv("INPUT_PATH"))
	fmt.Println("key:", os.Getenv("INPUT_KEY"))
	fmt.Println("key:", os.Getenv("INPUT_RESTORE-KEYS"))
}
