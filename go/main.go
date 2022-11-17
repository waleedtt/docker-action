package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("path:", os.Getenv("path"))
	fmt.Println("key:", os.Getenv("key"))
	fmt.Println("key:", os.Getenv("restore-keys"))
}
