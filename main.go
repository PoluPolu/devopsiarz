package main

import "fmt"

func main(){
	// pierw()
	// druga()
	stringi()
}

func stringi(){
	MojStr := "PawcioSrapcio"
	fmt.Printf("Moj string to typ: %T\n", MojStr)

	MojStrLenght := len(MojStr)
	fmt.Println("Lenght of Mojstr: ", MojStrLenght)

}

func druga() {
	var x int32 = 1 
	var y int32 = 1
	
	suma := x + y
	fmt.Println("Suma wynosi:",suma)
	fmt.Printf("suma jest typu %T\n",suma)

	a := []string{"a","b","c","d","e"}
	fmt.Printf("a jest typu %T\n",a)

	fmt.Println(a)
}

func pierw() {
	fmt.Println("Czesc Polu")
}