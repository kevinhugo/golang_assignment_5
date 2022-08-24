package main
import "fmt"

type human struct {
	name string
}

func main() {
	humans := []human{
		{name : "Mr"},
		{name : "Sir"},
		{name : "Bro"},
		{name : "Friend"},
		{name : "Kevin"},
		{name : "Hugo"},
		{name : "Ke"},
		{name : "Vin"},
		{name : "Hu"},
		{name : "Go"},
	}
	var pointers []*human
	for key, _ := range humans {
		pointers = append(pointers,&humans[key])
	}

	printNames := func (points []*human) {
		for index, eachHuman := range points {
			fmt.Println(index+1,eachHuman.name)
		}
	}

	printNames(pointers)


	// names := []string{"Mr","Sir","Bro","Friend","Kevin","Hugo","Ke","Vin","Hu","Go"}
	// var pointers []*string
	// for key, _ := range names {
	// 	pointers = append(pointers, &names[key])
	// }
	// printNames := func (points []*string) {
	// 	for index, eachPoint := range points {
	// 		fmt.Println(index+1,*eachPoint)
	// 	}
	// }
	// printNames(pointers)
}

// func printNames(names []string) {
// 	for index, eachName := range names {
// 		fmt.Println(index+1,eachName)
// 	}
// }