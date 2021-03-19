package main
 
import "fmt"
 
func main() {
    fmt.Println(calculation("Rectangle",20, 30))
    fmt.Println(calculation("Square",20))
}
 
func calculation(str string, y ...int) int {  // if we pass spread operater first it will throw an panic error 
 
    area := 1
 
    for _, val := range y {
        if str == "Rectangle" {
            area *= val
        } else if str == "Square" {
            area = val * val
        }
    }
    return area
}