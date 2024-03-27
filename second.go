package main

import (
	"fmt"
	"math"
)

// defien struct Person
type Person struct {
	massa  float64
	tinggi float64
}

// method hitungBMI milik struct Person
func (p Person) hitungBMI() float64 {
	// math.Pow menghitung pangkat
	var bmi = p.massa / math.Pow(p.tinggi, 2)
	return bmi
}

func isHigher(mark, john float64) {
	// mengembalikan return value boolean dengan operator perbandingan
	markHigherBMI1 := mark > john
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John?", markHigherBMI1)

}

func main() {
	// pembuatan objek dan inisiasi value
	var mark1 = Person{massa: 78.0, tinggi: 1.69}
	var mark2 = Person{massa: 95.0, tinggi: 1.88}

	var john1 = Person{massa: 92.0, tinggi: 1.95}
	var john2 = Person{massa: 85.0, tinggi: 1.76}

	// memanggil method hitungBMI() dengan objek struct Person
	bmiMark1 := mark1.hitungBMI()
	bmiJohn1 := john1.hitungBMI()

	bmiMark2 := mark2.hitungBMI()
	bmiJohn2 := john2.hitungBMI()

	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark : %.2f\n", bmiMark1)
	fmt.Printf("BMI John : %.2f\n", bmiJohn1)
	// memanggil fungsi isHigher() dengan passing 2 parameter
	isHigher(bmiMark1, bmiJohn1)

	fmt.Println("\nData 2:")
	fmt.Printf("BMI Mark : %.2f\n", bmiMark2)
	fmt.Printf("BMI John : %.2f\n", bmiJohn2)
	isHigher(bmiMark2, bmiJohn2)
}
