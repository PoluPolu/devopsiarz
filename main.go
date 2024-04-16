package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type MyFuncDod = func(x, y int) int

func main() {
	// pierw()
	// druga()
	// stringi()
	// Petle()
	// SwitcheIfy()
	// zewnetrzna() //wywolanie go run *.go lub go run .

	// fmt.Println(dodawanko1(1,2))
	// txt, err := tekscior("Paawel", "Robert")
	// fmt.Println(txt, err)

	// if txt, err := tekscior("Paawel", "Robert"); txt == "Pawel Robert" {
	// 	fmt.Println(txt, err)
	// } else {
	// 	fmt.Println("Ni hujaaa a jest >5, err: ", err)
	// }\

	// var fd MyFuncDod = dodawanko1 //tu mozna dac od razy body
	// var fd2 MyFuncDod = func (s,d int) int {
	// 	return s + d
	// }
	// rdodawania := fd(1,2)
	// fmt.Println(rdodawania)
	// rdodawania2 := fd2(3,2)
	// fmt.Println(rdodawania2)

	// DevOpsFunc := func(num ...int) {
	// 	fmt.Println("Liczba cyderek: ",len(num))
	// }
	// DevOpsFunc(1)
	// DevOpsFunc(1,2,33,4,5,6)

	// inp := "Tekscik doi funkcyji"
	// func(str string){
	// 	fmt.Println(str)
	// }(inp)

	//struktury()

	// pointery()
	pointeryJamie()

	// slicyarrarje()
}

func slicyarrarje() {
	var MyArray [5][10]int
	var MySlice1 = make([]int, 10)
	MySlice2 := make([]int, 5)
	MySlice3 := []int{1,2,3}


	fmt.Println("MySlice == nil", MySlice1 == nil)
	fmt.Println(len(MySlice1))
	fmt.Println(cap(MySlice1))
	fmt.Println(MySlice1)
	MySlice1 = append(MySlice1, 1)
	fmt.Println(MySlice1)
	fmt.Println(len(MySlice1))
	fmt.Println(cap(MySlice1))
	fmt.Println("---------------------")

	fmt.Println(MySlice1)
	fmt.Println(MySlice2)
	fmt.Println(MySlice3)

	fmt.Println("---------------------")

	fmt.Println("Prin", len(MyArray))

	rows := 3
	cols := 4
	myArray2 := make([][]int, rows)
	for i := range myArray2 {
		myArray2[i] = make([]int, cols)
	}

	for i := range myArray2 {
		fmt.Println(i, myArray2[i])
	}

	fmt.Println(myArray2)

}

func pointery() {
	a := 100
	origPointer := &a
	fmt.Println("Starting value a: ", a)
	fmt.Println("origPointer: ", origPointer)
	fmt.Println("a: ", a)
	fmt.Println("origPointer val: ", *origPointer)

	fmt.Println(" ---1-- ")

	a++
	fmt.Println("a: ", a)
	fmt.Println("&a==origPointer", &a == origPointer)

	fmt.Println(" ---2-- ")

	a++
	fmt.Println("a: ", a)
	fmt.Println("&a==origPointer", &a == origPointer)

	fmt.Println(" ---3-- ")

	b := new(int)
	*b++
	c := *b
	fmt.Println("b (pointer): ", b)
	fmt.Println("b (value): ", *b)
	fmt.Println("c (value): ", c)
	fmt.Println("c (pointer): ", &c)

	fmt.Println(" ---4-- ")
	*b++

}

func tekscior(a, b string) (string, bool) {
	if len(a) == 5 {
		return a + "-" + b, true
	} else {
		return b + a, false
	}
}

func dodawanko1(x, y int) int {
	return x + y
}

func SwitcheIfy() {
	i := 7

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch {
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")
	case i == 6:
		fmt.Println("six")
		fallthrough
	case i == 7:
		fmt.Println("seven")
		// fallthrough
	case i > 5 && i < 8:
		fmt.Println("bigger than 5")
	default:
		fmt.Println("default")
	}

	var j interface{} = "asd"
	switch j.(type) {
	case nil:
		fmt.Println("nil")
	case bool:
		fmt.Println("boolean")
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}

	v := "21"
	if r, err := strconv.Atoi(v); err == nil {
		fmt.Printf("\nTyp: %T, wartosc: %d\n", r, r)
	}
}

func Petle() {
	samochody := []string{"Kia", "Merc", "BMW", "Audi", "Volvo", "Jeep", "Toyota"}
	for idx, samochody := range samochody {
		fmt.Printf("%d => %s\n", idx, samochody)
	}
	fmt.Println("------------")
	for _, samochody := range samochody {
		fmt.Printf("%s\n", samochody)
	}
	fmt.Println("------------")
	for i := -10; i < 1; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------")
	i := 0
	for {
		fmt.Println("Petla numer ", i)
		if i == 3 {
			fmt.Println("3 petla")
			time.Sleep(time.Second * 1)
			i++
			continue
		}
		if i == 7 {
			fmt.Println("7 petla - koniec")
			break
		}

		time.Sleep(time.Second * 1)
		i++
	}
}

func stringi() {
	MojStr := "PawcioSrapcio"
	fmt.Printf("Moj string to typ: %T\n", MojStr)

	MojStrLenght := len(MojStr)
	fmt.Println("Lenght of Mojstr: ", MojStrLenght)

	for i := 0; i < len(MojStr); i++ {
		// fmt.Printf("%x \n", MojStr[i])
		fmt.Printf("=> %x => %s\n", MojStr[i], string(MojStr[i]))
	}

	fmt.Println("--------------UNICODY--------------------------------------------------")
	MyStr := "㋛" //unicody to runy
	fmt.Println("Lenght of MyStr (emoji): ", len(MyStr))
	fmt.Println("Lenght of MyStr (emoji): ", utf8.RuneCountInString(MyStr))

	for i := 0; i < len(MyStr); i++ {
		fmt.Printf("Rune code point 0: %v\n", MyStr[i])
	}

	fmt.Println("-------zamiana liter---------------------------------------------------------")
	MyStr = "DevOpsiarz"
	fmt.Println("MyStr 2nd element: ", string(MyStr[2]))
	// MyStr[2] = "V" - tak sie nie da

	fmt.Println("---Slice----zamiana liter----modyfikacja-----------------------------------------------------")
	MyStr = "My String"
	ByteStr := []byte("This is string")

	fmt.Printf("MyStr is (%v) and type: %T\n", MyStr, MyStr)
	fmt.Printf("ByteStr is (%v) and type: %T\n", ByteStr, ByteStr)

	ByteFromStr := []byte(MyStr)
	fmt.Printf("ByteFromMyStr is (%v) and type: %T\n", string(ByteFromStr), ByteFromStr)

	ByteFromStr[0] = 'Y'
	fmt.Printf("ByteFromMyStr is (%v) and type: %T\n", string(ByteFromStr), ByteFromStr)

	fmt.Println("---wybieranie elementu-----------------------------------------------------")
	MyStr = "Pawel Polanowski"
	fmt.Println("My String is:", MyStr)
	fmt.Println("First element: ", string(MyStr[0]))
	fmt.Println("Last element: ", string(MyStr[len(MyStr)-1]))
	fmt.Println("First two element: ", string(MyStr[0:2]))
	fmt.Println("From 2nd element to the end: ", string(MyStr[2:]))
	fmt.Println("Up tp 7th element: ", string(MyStr[:7]))

	fmt.Println("---Splitowanie-----------------------------------------------------")
	MyStr = "Pawel Polanowski"
	Lists := strings.Split(MyStr, "a")
	fmt.Printf("Lists type: %T\n", Lists)
	fmt.Printf("Listselemnt type: %T\n", Lists[0])
	fmt.Println(Lists)

	fmt.Println("---int na string-----------------------------------------------------")
	MyInt := 100
	MyStringInt := "1000000"

	fmt.Printf("MyInt type: %T\n", MyInt)
	fmt.Printf("MyStringInt type: %T\n", MyStringInt)

	StringFormInt := strconv.Itoa(MyInt)
	fmt.Printf("StringFormInt valuee: %s, type: %T\n", StringFormInt, StringFormInt)

	IntFromString, _ := strconv.Atoi(MyStringInt)
	fmt.Printf("IntFromString valuee: %d, type: %T\n", IntFromString, IntFromString)

	fmt.Println("---Join strings-----------------------------------------------------")
	//najwolniejsza
	MyStr = "Pawel Polanowski"
	MyStr += " i Beata\n"

	fmt.Println(MyStr)

	MyStr = strings.Join([]string{"Pawel", "Polanowski", "oraz", "Beata"}, " ")
	fmt.Println(MyStr)

	//najlepsza i najszybsza
	NewStr := &strings.Builder{}
	NewStr.WriteString("Pawel Polanowski")
	NewStr.WriteString(" a takze Beata\n")
	fmt.Println(NewStr.String())
	fmt.Printf("Type: %T\n", NewStr.String())
	fmt.Printf("Type: %T\n", NewStr)

	fmt.Println("---szukanie powtarzanie-----------------------------------------------------")
	MyStr = "Pawel Polanowski i Beata\n"
	if strings.Contains(MyStr, "Gawe") {
		fmt.Println("Jest Pawel")
	} else {
		fmt.Println("Nie ma Gawela")
	}

	RepeaterChar := "Pa"
	fmt.Println(strings.Repeat(RepeaterChar, 21))
}

func druga() {
	var x int32 = 1
	var y int32 = 1

	suma := x + y
	fmt.Println("Suma wynosi:", suma)
	fmt.Printf("suma jest typu %T\n", suma)

	a := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("a jest typu %T\n", a)

	fmt.Println(a)
}

func pierw() {
	fmt.Println("Czesc Polu")
}
