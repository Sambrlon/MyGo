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

package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := [3]string{"1", "2", "3"}
	messages[1] = "5"
	fmt.Println(messages)
}

func printMessages(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	fmt.Println(messages)
	return nil
}
