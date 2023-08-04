// package main

// import "fmt"

// func main() {
// 	message := sayHello("Артём", 24)
// 	printMessage(message)

// 	// 	printMessage("вызов функции 1\n")
// 	// 	printMessage("вызов функции 2\n")
// 	// 	printMessage("вызов функции 3\n")
// }

// func printMessage(message string) {
// 	fmt.Println(message)
// }

// func sayHello(name string, age int) string {
// 	return fmt.Sprintf("Привет %s! Тебе %d лет!", name, age)
// }

// func enterTheClub(age int) (string, bool) {
// 	if age >= 18 {
// 		return "Добро пожаловать в клуб", true
// 	} else {
// 		return "Тебе нет 18", false
// 	}

// }

//----------------------------------------------------------------------

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	message, err := enterTheClub(55)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(message)
// 	// fmt.Println(entered)
// }

// func enterTheClub(age int) (string, error) {
// 	if age >= 18 && age < 45 {
// 		return "Добро пожаловать", nil
// 	} else if age >= 45 && age < 65 {
// 		return "Вам точно понравится такая музыка?", nil
// 	} else if age >= 65 {
// 		return "Для вас это слишком активный отдых", errors.New("yoy are too old")
// 	}

// 	return "Тебе нет 18", errors.New("you are too young")
// }

//----------------------------------------------------------------------

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	message, err := enterTheClub(55)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(message)
// 	// fmt.Println(entered)
// }

// func enterTheClub(age int) (string, error) {
// 	if age >= 18 && age < 45 {
// 		return "Добро пожаловать", nil
// 	} else if age >= 45 && age < 65 {
// 		return "Вам точно понравится такая музыка?", nil
// 	} else if age >= 65 {
// 		return "Для вас это слишком активный отдых", errors.New("yoy are too old")
// 	}

// 	return "Тебе нет 18", errors.New("you are too young")
// }

// func prediction(dayOfWeek string) string {
// 	if dayOfWeek == "понедельник" {
// 		return "Хорошего понедельника"
// 	} else if dayOfWeek == "вторник" {
// 		return "Хорошего вторника"
// 	} else if dayOfWeek == "среда" {
// 		return "Хорошей среды"
// 	} else if dayOfWeek == "четверг" {
// 		return "Хорошего четверга"
// 	} else if dayOfWeek == "пятница" {
// 		return "Хорошей пятницы"
// 	} else if dayOfWeek == "суббота" {
// 		return "Сегодня выходной"
// 	} else if dayOfWeek == "воскресенье" {
// 		return "Хороших выходных"
// 	}

// 	return ""
// }

// -----------------------------------------------------------------------

// Конструкция switch case

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	fmt.Println(prediction("пятница"))
// }

// func prediction(dayOfWeek string) (string, error) {
// 	switch dayOfWeek {
// 	case "понедельник":
// 		return "Хорошего понедельника", nil
// 	case "вторник":
// 		return "Хорошего вторника", nil
// 	case "среда":
// 		return "Хорошей среды", nil
// 	case "четверг":
// 		return "Хорошего четверга", nil
// 	case "пятница":
// 		return "Хорошей пятницы", nil
// 	default:
// 		return "Невалидный день недели", errors.New("invalid day of week")
// 	}
// }

// -----------------------------------------------------------------------

// Работа с циклом for

// package main

// import "fmt"

// func main() {
// 	fmt.Println(findMin(1, 2, 3, 4, 5, 6, 7, -2, -5, 0))
// 	fmt.Println(findMax(1, 2, 3, 4, 5, 6, 7, -2, -5, 0))

// }

// func findMin(numbers ...int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}

// 	min := numbers[0]

// 	for _, i := range numbers {
// 		if i < min {
// 			min = i
// 		}
// 	}
// 	return min

// }

// func findMax(numbers ...int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}

// 	max := numbers[0]

// 	for _, i := range numbers {
// 		if i > max {
// 			max = i
// 		}
// 	}
// 	return max

// }

// ------------------------------------------------------------------------

// Анонимные функции

// package main

// import "fmt"

// func main() {
// 	func() {
// 		fmt.Println("анонимная функция")
// 	}()

// }

// Замыкание

// package main

// import "fmt"

// var msg string

// func init() {
// 	msg = "from init()"
// }

// func main() {
// 	fmt.Println(msg)

// 	inc := increment()
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(inc())
// 	fmt.Println(inc())

// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// 	fmt.Println(increment2())
// }
// func increment() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func increment2() int {
// 	count := 0
// 	count++
// 	return count
// }
// -------------------------------------------------------------------------

// Указатели

// package main

// import "fmt"

// func main() {
// 	message := "Скоро я стану ниндзя!"
// 	fmt.Println(message)

// 	changeMessage(&message) // & референсинг

// 	fmt.Println(message)
// }

// func changeMessage(message *string) {
// 	*message += " (из фунции changeMessage())" // * дереференсинг, указатель на строку
// }
// ------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	number := 5
// 	var p *int

// 	p = &number

// 	fmt.Println(p)
// 	fmt.Println(&number)

// 	*p = 10

// 	fmt.Println(number)

// }

// func changeMessage(message *string) {
// 	*message += " (из фунции changeMessage())" // * дереференсинг, указатель на строку
// }
// ------------------------------------------------------------------------

// Массивы

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	messages := [3]string{"1", "2", "3"}

// 	printMessages(messages)

// 	fmt.Println(messages)
// }

// func printMessages(messages [3]string) error {
// 	if len(messages) == 0 {
// 		return errors.New("empty array")
// 	}

// 	messages[1] = "5"
// 	fmt.Println(messages)

// 	return nil
// }
// ------------------------------------------------------------------------

// Слайсы

// package main

// import "fmt"
// import "errors"

// func main() {
// 	messages := make([]string, 10000)
// 	fmt.Println(len(messages))
// 	fmt.Println(cap(messages))
// 	messages = append(messages, "10001")
// 	fmt.Println(len(messages))
// 	fmt.Println(cap(messages))
// }

// func printMessages(messages []string) error {
// 	if len(messages) == 0 {
// 		return errors.New("empty array")
// 	}

// 	messages[1] = "5"
// 	fmt.Println(messages)

// 	return nil
// }
// ------------------------------------------------------------------------

// Слайсы в матрице
// package main

// import "fmt"

// func main() {
// 	matrix := make([][]int, 10)

// 	counter := 0
// 	for x := 0; x < 10; x++{
// 		matrix[x] = make([]int, 10)
// 		for y := 0; y < 10; y++{
// 			counter++
// 			matrix[x][y] = counter
// 		}
// 		fmt.Println(matrix[x])
// 	}
// }
//------------------------------------------------------------------------

// Цикл for и слайсы

// package main

// import "fmt"

// func main(){
// 	// messages := []string {
// 	// 	"message 1",
// 	// 	"message 2",
// 	// 	"message 3",
// 	// 	"message 4",
// 	// }
// 	// for _, message := range messages{
// 	// 	fmt.Println(message)
// 	// }

// 	counter := 0

// 	for {
// 		if counter ==100{
// 			break
// 		}

// 		counter++
// 		fmt.Println(counter)
// 	}
// }
// ---------------------------------

// 4.08
// паника

// package main

// import "fmt"

// func main() {
// 	defer handlerPanic()

// 	messages := []string{
// 		"message 1",
// 		"message 2",
// 		"message 3",
// 		"message 4",
// 	}

// 	messages[4] = "message 5"

// 	fmt.Println(messages)

// }

// func handlerPanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println(r)
// 	}

// 	fmt.Println("handlerPanic() выполнен успешно")

// }
// -------------------------

// Мапы
// package main

// import "fmt"

// func main() {
// 	users := map[string]int{
// 		"Artem":  16,
// 		"Ruben":  23,
// 		"Gergen": 22,
// 	}

// 	users["Kutyika"] = 21

// 	delete(users, "Ruben")

// 	var usersa map[string]int
// 	usersa = make(map[string]int)

// 	usersa["Alena"] = 10

// 	for key, value := range users  {
// 		fmt.Println(key, value)
// 	}

// 	for key, value := range usersa {
// 		fmt.Println(key, value)
// 	}
// }
// ----------------------------------------------------

// // Структура

// package main

// import "fmt"

// type Age int

// func (a Age) isAdult() bool {
// 	return a >= 18
// }

// type User struct {
// 	name   string
// 	age    Age
// 	sex    string
// 	weight int
// 	height int
// }

// // func (u User) printUserInfo(name string) { // Является методом структуры
// // 	u.name = name
// // 	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
// // }

// func (u *User) setName(name string) {
// 	u.name = name
// }

// func (u User) getName() string {
// 	return u.name
// }

// type DumbDatabase struct {
// 	m map[string]string
// }

// func NewBumbDatabase() *DumbDatabase { // Инициализация мапа
// 	return &DumbDatabase{
// 		m: make(map[string]string),
// 	}

// }

// func NewUser(name, sex string, age, weight, height int) User {
// 	return User{
// 		name:   name,
// 		age:    Age(age),
// 		sex:    sex,
// 		weight: weight,
// 		height: height,
// 	}
// }

// func main() {
// 	user1 := User{"Artem", 24, "Male", 77, 175}           // В структуре значения передаются в {} скобках
// 	user2 := NewUser("BumbleBizy", "Female", 22, 55, 155) // В конструкторе значения передаются в () скобках

// 	// fmt.Printf("%+v\n", user1)
// 	// fmt.Printf("%+v\n", user2)

// 	// fmt.Println(user1.name, user1.age)
// 	// fmt.Println(user2.name, user2.age)

// 	// printUserInfo(user1)
// 	// printUserInfo(user2)

// 	user1.setName("Tagir")
// 	user2.setName("Kisell")

// 	fmt.Println(user1.getName())
// 	fmt.Println(user1.age, user1.age.isAdult())
// 	fmt.Println(user2.getName())

// }

// // func printUserInfo(user User) { // Является функцией
// // 	fmt.Println(user.name, user.age, user.sex, user.weight, user.height)
// // }
// ----------------------------------------------------------------

// // Интерфейсы

// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"github.com/zhashkevych/scheduler"
// )

// func main() {
// 	ctx := context.Background()

// 	t := time.Now()
// 	fmt.Println(t)

// 	worker := scheduler.NewScheduler()
// 	worker.Add(ctx, parseSubscriptionData, time.Second*5)
// 	worker.Add(ctx, sendStatistics, time.Second*10)

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt, os.Interrupt)

// 	<-quit
// 	worker.Stop()
// }

// func parseSubscriptionData(ctx context.Context) {
// 	time.Sleep(time.Second * 1)
// 	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
// }

// func sendStatistics(ctx context.Context) {
// 	time.Sleep(time.Second * 5)
// 	fmt.Printf("statistics sent at %s\n", time.Now().String())
// }

// func main() {
// 	square := shape.NewSquare(5)
// 	// ciecle := Circle{8}

// 	printShapeArea(square)
// 	// printShapeArea(ciecle)

// 	printInterface(square)
// 	// printInterface(ciecle)

// 	scheduler := scheduler.NewScheduler()

// }

// func printShapeArea(s shape.Shape) { // Принимает интерфейс в качетсве аргумента
// 	fmt.Println(s.Area())
// 	fmt.Println(s.Perimeter())
// }

// func printInterface(i interface{}) {
// 	switch value := i.(type) {
// 	case int:
// 		fmt.Println("int", value)

// 	case bool:
// 		fmt.Println("bool", value)
// 	default:
// 		fmt.Println("unknown type", value)
// 	}
// 	// fmt.Printf("%+v\n", i)
// }

// //---------------------------------------------------------------

// Дженерики

package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(c, "2"))
}

func sum[V Number](input []V) V {
	var result V

	for _, number := range input {
		result += number
	}
	return result
}

func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}

	return false
}
