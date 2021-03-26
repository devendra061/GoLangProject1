package main
import (
	"fmt"
	"github.com/devendra061/GoLangProject1/greetings"
)

func main(){
	message := greetings.Hello("Devendra")
	fmt.Println(message)
}
