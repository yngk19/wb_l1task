package t1

import "fmt"


// Создаём структуру Human
type Human struct {
	Name string
	Age int
	Gender string
}

// Определяем метод Greeting
func (h *Human) Greeting() {
	fmt.Println("Hello! I'm", h.Name)
}

// Создаём структуру Action и встраиваем в неё Human
type Action struct {
	Human
}

// Определяем метод Speak, который будет обращаться к полям структуры Human
func (a *Action) Speak() {
	fmt.Println("I'm", a.Age, "years old")
}


func main() {
	var Person = Action{
		Human: {
			Name: "Ivan",
			Age: 30,
			Gender: "Male",
		}
	}
	// Можем вызывать методы как у Action так и у встроенной структуры
	Person.Greeting()
	Person.Speak()
}