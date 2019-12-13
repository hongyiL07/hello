package main

import (
	"time"
)

func main() {
	//将零个或者多个任意类型的命名变量组合在一起的聚合数据类型，每个变量都叫做 结构体的成员

	//定义一个叫 Employee 的结构体和一个结构体变量 dilbert
	type Employee struct {
		ID        int
		Name      string
		Address   string
		Bob       time.Time
		Position  string
		Salary    int
		ManggerID int
	}
	var dilbert Employee
	// dilbert 的每个成员都通过点号方式来访问
	dilbert.Salary -= 50000

	//或者获取幸运变量的地址，然后通过指针来访问它
	position := &dilbert.Position
	*position = "Senior" + *position

	//点号同样可以用在结构体指针上
	var employeeTheMonth *Employee = &dilbert
	employeeTheMonth.Position += "(proactive team player)"
	//等价于以下
	(*employeeTheMonth).Position += "(proactive team player)"

	//没有任何超越变量的结构体称为 空结构体，写作 struct{}
}

//利用二叉树来实现插入排序
type tree struct {
	value       int
	left, rigth *tree
}

//就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	apppendValues(values[:0], root)
}

//apppendValues 将元素按照顺序追加到 values 里面，然后返回结果 slice
func apppendValues(values []int, t *tree) []int {
	if t != nil {
		values = apppendValues(values, t.left)
		values = append(values, t.value)
		values = apppendValues(values, t.rigth)
	}
	return values
}
func add(t *tree, value int) *tree {
	if t == nil {
		//等价于返回  &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.rigth = add(t.rigth, value)
	}
	return t
}
