package main 
import "fmt"

type animal interface{
    speak()

}

type dog struct{}
//type cat struct{}

func(d dog) speak()string{
    return "bho-bho"
}
func main(){
   d1:= dog{}
   fmt.Println(d1.speak())
}
