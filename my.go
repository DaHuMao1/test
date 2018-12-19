package main
import "fmt"
func main(){
    var arr=make([]int,2)
    arr[0]=0
    arr[1]=1
    arr = append(arr,3,4,5,6)
    for i,e:= range arr{
        fmt.Println(i)
        fmt.Println(e)
    }
}
