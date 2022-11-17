package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	fmt.Println("path:", os.Getenv("INPUT_PATH"))
	fmt.Println("key:", os.Getenv("INPUT_KEY"))
	fmt.Println("key:", os.Getenv("INPUT_RESTORE-KEYS"))
	
	fmt.Printf(`::set-output name=cache-hit::1`)

	out, err := exec.Command("ls", "-ltr").Output()
	fmt.Printf("1 %s 11 %s", out, err)
}
