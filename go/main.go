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

	out, _ := exec.Command("ls", "-la").Output()
	fmt.Printf("%s", out)

	out1, _ := exec.Command("ls", "-la", "/").Output()
	fmt.Printf("%s", out1)

	out2, _ := exec.Command("pwd").Output()
	fmt.Printf("%s", out2)

	out3, _ := exec.Command("ls", "-la". '~/').Output()
	fmt.Printf("%s", out3)

}
