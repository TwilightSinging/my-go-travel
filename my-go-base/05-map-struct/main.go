package main

import (
	"fmt"
	"sort"
)

/*
设计一个 Student 结构体（姓名、年龄、成绩 map），
实现一个 Classroom 类型管理学生列表，支持添加学生、按成绩排名、计算平均分。
*/
type Student struct {
	Name  string
	Age   int
	Score map[string]int //存科目：成绩
}
type Classroom struct {
	Students []Student
}

func (c *Classroom) AddStudent(s Student) {
	c.Students = append(c.Students, s)
}
func (c *Classroom) RankStudents() {
	sum := make(map[string]int, len(c.Students))
	//先算总分，而且防止因为排序Students数组导致sum[*Student]的地址变化
	for _, stu := range c.Students {
		sum[stu.Name] = 0
		for _, score := range stu.Score {
			sum[stu.Name] += score
		}
	}
	sort.Slice(c.Students, func(i int, j int) bool {

		fmt.Printf("sum[%s]:%d sum[%s]:%d\n", c.Students[i].Name, sum[c.Students[i].Name], c.Students[j].Name, sum[c.Students[j].Name])
		return sum[c.Students[i].Name] > sum[c.Students[j].Name]
	})
}
func (c *Classroom) CalculateAverageScore() {
	s := 0
	for i := 0; i < len(c.Students); i++ {
		for _, score := range c.Students[i].Score {
			s += score
		}
		fmt.Printf("Student[%s] average score:%.2f\n", c.Students[i].Name, float64(s)/float64(len(c.Students[i].Score)))
		s = 0
	}

}
func main() {
	c := Classroom{}
	c.AddStudent(Student{Name: "张三", Age: 18, Score: map[string]int{"语文": 65, "数学": 60, "英语": 85}})
	c.AddStudent(Student{Name: "李四", Age: 19, Score: map[string]int{"语文": 90, "数学": 85, "英语": 80}})
	c.AddStudent(Student{Name: "王五", Age: 20, Score: map[string]int{"语文": 85, "数学": 80, "英语": 75}})
	fmt.Println("==========排名前==========")
	for _, stu := range c.Students {
		fmt.Println(stu.Name, stu.Age, stu.Score)
		//Println会自动换行
	}
	c.RankStudents()
	fmt.Println("==========排名后==========")
	for _, stu := range c.Students {
		fmt.Println(stu.Name, stu.Age, stu.Score)
	}
	//fmt.Println(c.CalculateAverageScore())
	c.CalculateAverageScore()
}
