///package main

//import (
//"fmt"
//)

// func MultiTable(number int) string {
// var out string
//
//	for i := 1; i <= 10; i++ {
//		out += fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
//	}
//
// return out
// }
//
//	func main() {
//		ret := MultiTable(10)
//		fmt.Println(ret)
//	}
//
// func Rps(p1, p2 string) string {
// var out string
//
//	if p1 == "scissors" && p2 == "paper" {
//		out = "Player 1 won!"
//	} else if p1 == "scissors" && p2 == "rock" {
//
//		out = "Player 2 won!"
//	} else if p1 == "paper" && p2 == "paper" {
//
//		out = "Draw!"
//	}
//
// return out
// }
// func Derive(coefficient, exponent int) string {
// var out string
// out = fmt.Sprintf("%dxˆ%d", coefficient*exponent, exponent-1)
// return out
// }
// func multipleOfIndex(ints []int) []int {
// var out []int
//
//	for index, ut := range ints {
//		if index > 0 && ut%index == 0 {
//			out = append(out, ut)
//		}
//	}
//
// return out
// }
// type function func(int) int
// var  out func
// a := x * 5
// return a

//func first(second func(x int) int) func(x int) int {
//return func(x int) int {
//		fmt.Println("the result of the execution")
//		return second(7)
//}
//}

//func main() {
//result := func(x int) int {
//	a := x * 5
//	return a
//}
//var c int
////fmt.Println(first(result)(c))
//}

//func main() {
//friendsNames := []string{"Аня", "Коля", "Лёша", "Лена", "Миша"}
//friendsCities := []string{"Владивосток", "Красноярск", "Москва", "Обнинск", "Чебоксары"}

// Объявлен пустой словарь, его нужно будет наполнить элементами,
// каждый из которых составлен по схеме "имя: город"
//	friends := map[string]string{}

// Допишите ваш код сюда
//	for index, _ := range friendsNames {
//		friends[friendsNames[index]] = friendsCities[index]

//	}

//fmt.Println("Лена живёт в городе:", friends["Лена"])
//	fmt.Println(friends)
//}

//package main

//import (
//	"fmt"
//"sort"
////)

//func getSortedKeys(mp map[int]string) []int {
//new := make([]int, 0)
//for key := range mp {
//	new = append(new, key)
//}
//sort.Ints(new)
//return new
//}
//func main() {
//moves := map[int]string{
//		1991: "Терминатор 2: Судный день",
//		1984: "Терминатор",
//		2009: "Терминатор: Да придет спаситель",
//		2003: "Терминатор 3: Восстание машин",
//		2015: "Терминатор: Генезис",
//		2019: "Терминатор: Темные судьбы",
////	}

//	sortedMoves := getSortedKeys(moves)

//	fmt.Println("Как смотреть франшизу Терминатор")

//	fmt.Println()

//	for _, year := range sortedMoves {
//		fmt.Println("Год:", year, "Фильм:", moves[year])
//	}
//}

//package main

////import "fmt"

//func main() {
//favoriteSongs := map[string][]string{
//	"Серёга": {"Unforgiven", "Holiday", "Highway to Hell"},
//		"Соня":   {"Shake It Out", "The Show Must Go On", "Наше лето"},
//		"Дима":   {"Владимирский централ", "Мурка", "Третье сентября"},
//	}
//lovSonyaSong := favoriteSongs["Соня"]
// Ниже напишите код, который напечатает на экран, сколько у Димы любимых песен

// Ниже напишите код, который построчно напечатает
//fmt.Println(len(favoriteSongs["Дима"]))
//	for _, column := range favoriteSongs["Соня"] {
//		fmt.Println(column)
////	}
//fmt.Println(lovSonyaSong) // все любимые песни Сони
//}

//package main

//import (
//"fmt"
//"sort"
//)

//var Database = map[string]string{
//	"Серёга": "Омск",
//	"Соня":   "Москва",
//	"Дима":   "Челябинск",
//	"Алина":  "Красноярск",
//	"Егор":   "Пермь",
//	"Коля":   "Екатеринбург",
//}

// вернёт отсортированную коллекцию ключей передаваемой мапы
//func getSortedKeys(collection map[string]string) []string {
//	keys := make([]string, 0, len(collection))
//	for key := range collection {
//		keys = append(keys, key)
//	}
//	sort.Strings(keys)
//	return keys
//}

//func askAlice(query string) {
//	if query == "Сколько у меня друзей?" {
//		count := len(Database)
//		fmt.Println("У тебя", count, "друзей.")
//		return
//	}
//	if query == "Кто все мои друзья?" {
//		var friendsString string
//		for _, friend := range getSortedKeys(Database) {
//			friendsString += friend + " "
//		}
//		fmt.Println("Твои друзья:", friendsString)
//		return
//	}
//	if query == "Где все мои друзья?" {
//		var friendCity string
//		for _, city := range Database {
//			friendCity += city + " "
//		}
//		fmt.Println("Твои друзья в городах:", friendCity)
//		return
//	}
// если ни одно из условий не подошло
//	fmt.Println("<неизвестный запрос>")
//}

// если ни одно из условий не подошло

// Не изменяйте следующий код
//
//	func main() {
//		fmt.Println("Привет, я Алиса!")
//		askAlice("Сколько у меня друзей?")
//		askAlice("Кто все мои друзья?")
//		askAlice("Где все мои друзья?")
//	}
//package main

//import (
//	"fmt"
//	"strings"
//)

//func main() {
//	quote1 := "Работает? Не трогай"
//	quote2 := "Если твой код работает, значит, это хороший код"
//	quote3 := "Лень — главное достоинство программиста"

// вызовы функции готовы к работе, не изменяйте их!

// тут вызывается функция antepenultWord с аргументом quote1 и печатается результат её работы
//	fmt.Println(antepenultWord(quote1))

// то же, но с аргументом quote2
//	fmt.Println(antepenultWord(quote2))

// то же с аргументом quote3
//	fmt.Println(antepenultWord(quote3))
////}

//func antepenultWord(msg string) string {
//msgSlice := strings.Split(msg, " ")
//if len(msgSlice) < 3 { // здесь проверяется длина слайса, она должна быть больше или равна 3
//	return "Строка содержит меньше трёх слов" // если длина слайса меньше 3
//}
//i := msgSlice[len(msgSlice)-3]
//fmt.Println(len(msgSlice))
//return i
//}

// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// var database = map[string]string{
// 	"Серёга": "Омск",
// 	"Соня":   "Москва",
// 	"Миша":   "Оренбург",
// 	"Дима":   "Челябинск",
// 	"Алина":  "Красноярск",
// 	"Егор":   "Пермь",
// 	"Коля":   "Екатеринбург",
// }

// // новая функция, она возвращает правильное словосочетание, склоняя слово «друзья»
// // в зависимости от того, какое число передано в аргументе friendsCount
// func formatFriendsCount(friendsCount int) string {
// 	if friendsCount == 1 {
// 		return "1 друг"
// 	}
// 	if friendsCount >= 2 && friendsCount <= 4 {
// 		return fmt.Sprintf("%d друга", friendsCount)
// 	}
// 	return fmt.Sprintf("%d друзей", friendsCount)
// }

// func processAlice(query string) string {
// 	if query == "сколько у меня друзей?" {
// 		// вызовите функцию formatFriendsCount() и передайте в неё friendsCount;
// 		// отредактируйте строку ниже: в ней должно использоваться выражение,
// 		// которое вернёт функция formatFriendsCount()
// 		friendsCount := len(database) // friendsCount должна содержать общее количество друзей в database
// 		friend := formatFriendsCount(friendsCount)
// 		return fmt.Sprintf("У тебя %s.", friend)
// 	}
// 	if query == "кто все мои друзья?" {
// 		var friends []string
// 		for name := range database {
// 			friends = append(friends, name)
// 		}
// 		sort.Strings(friends)
// 		return fmt.Sprintf("Твои друзья: %s", strings.Join(friends, ", "))
// 	}
// 	if query == "где все мои друзья?" {
// 		uniqueCities := map[string]int{}
// 		for _, city := range database {
// 			// заполняем мапу без дублирования городов
// 			uniqueCities[city] = 1
// 		}
// 		var cities []string
// 		// получаем уникальные названия городов
// 		for city := range uniqueCities {
// 			cities = append(cities, city)
// 		}
// 		sort.Strings(cities)
// 		return fmt.Sprintf("Твои друзья в городах: %s", strings.Join(cities, ", "))
// 	}
// 	return "неизвестный запрос1"
// }
// func processFriend(name, query string) string {
// 	for n := range database {
// 		if name == n && query == "ты где?" {
// 			return fmt.Sprintf("%s в городе: %s", name, database[name])
// 		} else if name == n && query != "ты где?" {
// 			return "неизвестный запрос2"
// 		}
// 	}
// 	return fmt.Sprintf("У тебя нет друга по имени %s", name)
// }
// func processQuery(query string) string {
// 	division := strings.Split(query, ", ")
// 	if division[0] == "Алиса" {
// 		return fmt.Sprint(processAlice(division[1]))
// 	} else {
// 		return fmt.Sprint(processFriend(division[0], division[1]))
// 	}
// }

// func main() {
// 	fmt.Println("Привет, я Алиса!")
// 	fmt.Println(processQuery("Алиса, сколько у меня друзей?"))
// 	fmt.Println(processQuery("Алиса, кто все мои друзья?"))
// 	fmt.Println(processQuery("Алиса, где все мои друзья?"))
// 	fmt.Println(processQuery("Алиса, кто виноват?"))
// 	fmt.Println(processQuery("Соня, ты где?"))
// 	fmt.Println(processQuery("Коля, что делать?"))
// 	fmt.Println(processQuery("Антон, ты где?"))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// 	//"time"
// )

// func whatTime(friend string) time.Time {
// 	var database = map[string]string{
// 		"Серёга": "Омск",
// 		"Соня":   "Москва",
// 		"Дима":   "Челябинск",
// 		"Алина":  "Красноярск",
// 		"Егор":   "Пермь",
// 	}
// 	var offsetUTC = map[string]int{
// 		"Санкт-Петербург": 3,
// 		"Москва":          3,
// 		"Самара":          4,
// 		"Новосибирск":     7,
// 		"Екатеринбург":    5,
// 		"Нижний Новгород": 3,
// 		"Казань":          3,
// 		"Челябинск":       5,
// 		"Омск":            6,
// 		"Ростов-на-Дону":  3,
// 		"Уфа":             5,
// 		"Красноярск":      7,
// 		"Пермь":           5,
// 		"Воронеж":         3,
// 		"Волгоград":       3,
// 		"Краснодар":       3,
// 		"Калининград":     2,
// 	}
// 	city := database[friend]
// 	utc := time.Now().UTC()
// 	time := utc.Add(time.Duration(offsetUTC[city]) * time.Hour)
// 	// напишите недостающий код
// 	// ...
// 	return time
// }

// func main() {
// 	t := whatTime("Соня")
// 	fmt.Printf("У Сони %d ч.\n", t.Hour())
// }

//

// package main

// import (
// 	"fmt"
// )

// // func twoSum(nums []int, target int) []int {
// // 	result := make([]int, 0)
// // outerLoop:
// // 	for ind1, x := range nums {
// // 		for ind2, y := range nums {
// // 			if x+y == target && ind1 != ind2 {
// // 				result = append(result, ind1+1)
// // 				result = append(result, ind2+1)
// // 				break outerLoop
// // 			}
// // 		}
// // 	}
// // 	return result

// // }

// //	func isPalindrome(x int) bool {
// //		var number, rem, temporary int
// //		var reverse int = 0
// //		number = x
// //		temporary = number
// //		for {
// //			rem = number % 10
// //			reverse = reverse*10 + rem
// //			number /= 10
// //			if number == 0 {
// //				break
// //			}
// //		}
// //		if temporary == reverse {
// //			return true
// //		} else {
// //			return false
// //		}
// //	}
// func latinToint(x string) int {
// 	in := []rune(x)
// 	latInt := map[string]int{
// 		"M": 1000,
// 		"D": 500,
// 		"C": 100,
// 		"L": 50,
// 		"X": 10,
// 		"V": 5,
// 		"I": 1,
// 	}
// 	var out int
// 	for N := range len(in) - 1 {
// 		if latInt[string(in[N])] < latInt[string(in[(N+1)])] {
// 			out -= latInt[string(in[N])]
// 		} else {
// 			out += latInt[string(in[N])]
// 		}
// 	}
// 	return out + latInt[string(in[len(in)-1])]
// }
// func main() {
// 	fmt.Print(latinToint("MCMXCVI"))
// }

// package main

// import "fmt"

// func main() {
// 	response := 424562
// 	secInDay := 60 * 60 * 24 // секунд в сутках

// 	days := response / secInDay
// 	hours := response % secInDay / 3600
// 	minutes := response % 3600 / 60
// 	seconds := response % 60

//		fmt.Println(days, hours, minutes, seconds)
//	}
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(5 > 4)
// }

// package main

// import "fmt"

// func main() {
// 	numbers := []int{2, -4, 5, 0, 16, 44, 0, -1, 32, 4, 0}
// 	var number int
// 	for _, number = range numbers {
// 		if number > 0 {
// 			fmt.Println("Число положительное")
// 		} else if number < 0 {
// 			fmt.Println("Число отрицательное")
// 		} else {
// 			fmt.Println("Ноль")
// 		}
// 	} // место для вашего кода
// }

// package main

// import "fmt"

// func main() {
// 	movie := []rune("Джонни Мнемоник")

// 	// %c используется для вывода символа
// 	fmt.Printf("%c%c%c%c%c", movie[0], movie[11], movie[10], movie[13], movie[14])
// 	// добавьте выводы остальных 4 букв
// 	// ...
// }

// package main

// import "fmt"

// func main() {
// 	slice := []int{789, 0, -45, 123, -77, 1029, 56, 39}
// 	var imax int
// 	var imin int

// 	for i, v := range slice {
// 		if v > slice[imax] {
// 			imax = i
// 		} else if v < slice[imin] {
// 			imin = i
// 		}
// 	}
// 	fmt.Printf("Max slice[%d] = %d\n", imax, slice[imax])
// 	fmt.Printf("Min slice[%d] = %d\n", imin, slice[imin])
// }

// package main

// import "fmt"

// func detectLoop(list []int) bool {
// 	var (
// 		index   int
// 		counter int
// 	)
// 	for index >= 0 && index < len(list) {
// 		counter++
// 		index = list[index]
// 		if counter > len(list) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func main() {
// 	list := []int{1, 5, 4, 8, 7, 6, 2, 3}
// 	loop := []int{1, 5, 4, 8, 7, 1, 2, 3, 9}

// 	fmt.Println("list имеет петлю?", detectLoop(list))
// 	fmt.Println("loop имеет петлю?", detectLoop(loop))
// }

// package main

// import (
//     "fmt"
// )

// const (
//     ok = iota
//     cancelled
//     error
// )

// func CodeToMessage(code int) string {
// 	switch code {
// 		case 0:
// 		return ok
// 		case 1:
// 		return cancelled
// 		case 2:
// 		return error
// 	}
// 	return fmt.Sprintln("unknown status")
// }

// func main() {
//     fmt.Println(CodeToMessage(0))
//     fmt.Println(CodeToMessage(1))
//     fmt.Println(CodeToMessage(2))
//     fmt.Println(CodeToMessage(45))
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	first := 1
// 	secondStr := "2"
// 	thirdFloat := 3.0

// 	ifirst := int(first)
// 	isecondStr, _ := strconv.Atoi(secondStr)
// 	ithirdFloat := int(thirdFloat)
// 	result := ifirst + isecondStr + ithirdFloat
// 	fmt.Printf("Суммой чисел %d, %d и %d будет %d", ifirst, isecondStr, ithirdFloat, result)
// }

// package main

// import "fmt"

// // maxInt возвращает большее из всех чисел
// // или 0, если аргументы не были переданы
// func maxInt(numbers ...int) int {
// 	var imax int
// 	var out int

// 	for i, v := range numbers {
// 		if v > numbers[imax] {
// 			imax = i
// 			out = numbers[i]
// 		}

// 	}
// 	return out
// }

// func main() {
// 	max := maxInt(10, 11, -4, 15, 0)
// 	// ещё проверяем работу функции при вызове без параметров
// 	fmt.Println(max, maxInt())
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func checkLang(lang string) {
// 	slice := strings.Split(lang, ", ")
// 	for _, v := range slice {
// 		v = strings.ToLower(v)
// 		if v == "python" {
// 			fmt.Printf("%s — интерпретируемый!\n", v)
// 		} else {
// 			fmt.Printf("%s — компилируемый!\n", v)
// 		}
// 	}
// 	return
// }

// func main() {
// 	// строка, которую нужно разбить на слайс
// 	languages := "C++, Go, Rust, Java, Python"

// 	str := "— это языки программирования!"
// 	// выводим сообщение, объединив две строки
// 	fmt.Println(languages, str)

// 	// ниже используйте for для получения каждого элемента слайса и вывода его на экран
// 	checkLang(languages)
// }

//

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Привет, Практикум!")
}
