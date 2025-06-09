package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Post struct {
	Title string
	Tags  []string
}

type Bucket struct {
	Values []int
}

type Counter struct {
	Count int
}

func main() {
	x := 5
	y := 10
	num := 7
	arr := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2}
	user := User{Name: "JOhn", Age: 20}

	post := Post{Title: "Go learning", Tags: []string{}}

	swap(&x, &y)
	increment(&num)
	doubleValues(arr)
	updateName(&user, "Doe")
	appendSlice(&nums, 3)
	post.addTag("Go is good")
	newuser := newUser("Maria", 32)
	refUser := &newuser

	newuser.Name = "Klara"

	demonstrateCopy()

	count := Counter{Count: 1}

	count.Increment()
	count.Increment()
	count.Increment()
	count.Increment()
	count.Print()

	next := CounterFactory()

	fmt.Println("Swap", x, y)
	fmt.Println("Increment", num)
	fmt.Println("Doubled", arr)
	fmt.Println("Update name", user)
	fmt.Println("Append slice", nums)
	fmt.Println("Adding tag", post)
	fmt.Println("New user", newuser, *refUser)

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func swap(x *int, y *int) {
	temp := *x

	*x = *y
	*y = temp
}

func increment(inc *int) int {
	*inc = *inc + 1
	return *inc
}

func doubleValues(arr []int) {

	for i := range arr {
		arr[i] *= 2
	}
}

func updateName(u *User, newName string) {
	u.Name = newName
}

func appendSlice(nums *[]int, num int) {
	*nums = append(*nums, num)
}

func (p *Post) addTag(tagName string) {
	p.Tags = append(p.Tags, tagName)
}

func newUser(name string, age int) *User {
	return &User{Name: "JOhn", Age: 32}
}

func deepCopy(deepValues *[]int, originalValues []int) {
	copy(*deepValues, originalValues)
}

func demonstrateCopy() {
	original := Bucket{Values: []int{1, 2, 3, 4}}
	shallow := original

	deep := Bucket{make([]int, len(original.Values))}

	deepCopy(&deep.Values, original.Values)

	original.Values[0] = 99

	fmt.Println(shallow.Values[0])
	fmt.Println(deep.Values[0])
}

func (c *Counter) Increment() {
	c.Count += 1
}

func (c Counter) Print() {
	fmt.Println("Incremented by", c.Count)
}

func CounterFactory() func() int {
	var count int

	return func() int {
		count += 1
		return count
	}
}
