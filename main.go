package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// pierw()
	// druga()
	// stringi()
	Petle()
}


func Petle(){
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
	fmt.Println(strings.Repeat(RepeaterChar,21))
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
