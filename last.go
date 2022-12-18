package main
import "fmt"
func main(){

  var stroka, result string

    fmt.Scan(&stroka)

    result = stroka[len(stroka)-1:]


    fmt.Println(result)
}


// s      := "12121211122"
// first3 := s[0:3]
// last3  := s[len(s)-3:]


//Дано натуральное число, выведите его последнюю цифру.