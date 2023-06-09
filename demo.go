package main

//func main() {
//	fmt.Print("test")
//}

//func add(x, y int) int {
//	return x + y
//}
//
//func main() {
//	fmt.Print(add(42, 13))
//}

//func main() {
//	name := "陈俊豪"
//	age := 26
//	fmt.Print(name, age)
//}
import (
	"fmt"
	"net/http"
)

func main() {
	testHttp()
}
func testHttp() {
	//创建客户端
	client := http.Client{}
	//通过client去请求
	response, err := client.Get("https://www.zhoukezhang.com/")
	CheckErr(err)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
		//处理
	}
}

//检查错误
func CheckErr(err error) {
	//fmt.Println("09---------------")
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
