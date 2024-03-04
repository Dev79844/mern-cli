package main

import (
	cmd "github.com/Dev79844/mern-cli/cmd"
	// "os/exec"
	// "fmt"
)

func main(){
	cmd.Execute()
	// cmdFrontend := exec.Command("npm","create","vite@latest","blog","--","--template","react")
	// output, err := cmdFrontend.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(string(output))
}