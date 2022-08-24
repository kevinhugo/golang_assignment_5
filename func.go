package main
import "fmt"
import "os"
import "strconv"

type human struct {
	name string
	address string
	job string
	reason string
}

func main() {

	var number int = agrToInt(os.Args[1])

	humans := map[int]human{
		1 : {name : "Mr", address : "Sun", job : "Barber", reason : "Gabut"},
		2 : {name : "Sir", address : "Mercury", job : "Professional Barber", reason : "Gabut"},
		3 : {name : "Bro", address : "Venus", job : "Master Gabut", reason : "Gabut"},
		4 : {name : "Friend", address : "Earth", job : "Master Gabut", reason : "Gabut"},
		5 : {name : "Kevin", address : "Mars", job : "Master Gabut", reason : "Gabut"},
		6 : {name : "Hugo", address : "Jupiter", job : "Master Gabut", reason : "Gabut"},
		7 : {name : "Ke", address : "Saturn", job : "Master Gabut", reason : "Gabut"},
		8 : {name : "Vin", address : "Uranus", job : "Master Gabut", reason : "Gabut"},
		9 : {name : "Hu", address : "Neptune", job : "Master Gabut", reason : "Gabut"},
		10 : {name : "Go", address : "Pluto", job : "Master Gabut", reason : "Gabut"},
	}
	
	// humans := []map[int]human{}
	
	// var pointers map[int]*human
	// for key, _ := range humans {
	// 	pointers = append(pointers,&humans[key])
	// }

	// printNames := func (points []*human) {
	// 	for index, eachHuman := range points {
	// 		fmt.Println(index+1,eachHuman.name)
	// 	}
	// }

	// printNames(pointers)

	fmt.Println("You have chosen name : ",humans[number].name)
	fmt.Println("You have chosen address : ",humans[number].address)
	fmt.Println("You have chosen job : ",humans[number].job)
	fmt.Println("You have chosen reason : ",humans[number].reason)
}

func agrToInt(number string) int {
	value, err := strconv.Atoi(number)
    if err != nil {
        // ... handle error
        panic(err)
    }

	return value
}