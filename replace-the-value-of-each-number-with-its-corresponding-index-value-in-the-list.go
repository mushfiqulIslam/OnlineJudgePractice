import "fmt"


func replace(arr []int) string {
    for i := 0; i < len(arr); i+= 2 {
        indexValue := arr[i]
        arr[i] = arr[(i+1) % len(arr)]
        arr[(i+1) % len(arr)] = indexValue 
    }
	output := ""
	for i := 0; i < len(arr); i++ {
        output += fmt.Sprintf("%v",arr[i])
    }
    return output
}
func main() {
  fmt.Println("Hello World!")
  fmt.Println(replace([]int{3, 2, 0,1}))
  fmt.Println(replace([]int{4, 3, 2, 0,1}))
  fmt.Println(replace([]int{4, 3, 0, 5, 1, 2}))
}