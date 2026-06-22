// 声明当前包，可执行程序必须是 main
package main

// 导入依赖包
import (
	"fmt"
	"math/rand"
	"time"
)

// 定义结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 结构体绑定方法
func (s Student) ShowInfo() {
	fmt.Printf("学生姓名：%s，年龄：%d，分数：%.2f\n", s.Name, s.Age, s.Score)
}

// 普通函数：两数求和
func add(a, b int) int {
	return a + b
}

// 多返回值函数
func getMinMax(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	fmt.Println("========== Go 完整示例程序 ==========")

	// 1. 基础变量与输出
	var msg string = "Hello Goland"
	num := 666
	fmt.Println(msg, num)

	// 2. 调用普通函数
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// 3. 切片 + 多返回值函数
	numList := []int{12, 45, 7, 89, 23}
	minVal, maxVal := getMinMax(numList)
	fmt.Printf("数组最小值：%d，最大值：%d\n", minVal, maxVal)

	// 4. 结构体 + 方法调用
	stu := Student{
		Name:  "张三",
		Age:   18,
		Score: 92.5,
	}
	stu.ShowInfo()

	// 5. if 判断
	if stu.Score >= 90 {
		fmt.Println("成绩优秀")
	} else if stu.Score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	// 6. for 循环
	fmt.Println("\n循环打印 1~5：")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}

	// 7. 随机数演示
	rand.Seed(time.Now().UnixMilli())
	randomNum := rand.Intn(100)
	fmt.Printf("\n随机生成0-99数字：%d\n", randomNum)
}
