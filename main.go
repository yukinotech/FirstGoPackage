package FirstGoPackage

import (
	"fmt"
)

func main(){
	fmt.Println("Hello, World!")
	FirstGoPackage.Say('good bye')
}

func Say(words string){ 
	fmt.Println(words)
}