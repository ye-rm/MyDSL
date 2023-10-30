package main

import(
	"os"
	"fmt"
	"os/user"
	"awesomeDSL/repl"
)

func main(){
	user,err:=user.Current()
	if err!=nil{
		panic(err)
	}

	fmt.Printf("Hello %s! This is awesomeDSL to help you build your chatting bot\n",user.Username)
	fmt.Print("Please type in commands\n")
	repl.Start(os.Stdin,os.Stdout)
}