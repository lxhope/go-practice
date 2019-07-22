package example

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Food struct {
	Name string `json: "name"`
}

type Result struct {
	Code    int
	Message string
}

func (p *Person) Walk(step int) (res *Result) {
	fmt.Printf("%s walks %d steps\n", p.Name, step)
	return &Result{1, ""}
}

func (p *Person) Speak(word string) (res *Result) {
	fmt.Printf("%s speaks the work %s\n", p.Name, word)
	return &Result{2, "ok"}
}

func (p *Person) Eat(food *Food) (res *Result) {
	fmt.Printf("%s eat food %s\n", p.Name, food.Name)
	return &Result{3, "succ"}
}
