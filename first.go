package main

import (
	"fmt"
)

// mendefinisikan struct Team
type Team struct {
	name string
	scores []int
}


// method calculateAverage merupakan bagian dari struct Team
func (t Team) calculateAverage() int {
	total := 0
	for _, score := range t.scores {
		total += score
	}
	return total / len(t.scores)
}

// menerima 3 params, 2 object dari struct Team dan int value yang mengembalikan string
func determineWinner(scoresLumba Team, scoresKoala Team, minScore int) string {
	// memanggil method calculateAverage() dengan objek Team
	averageLumba := scoresLumba.calculateAverage()
	averageKoala := scoresKoala.calculateAverage()

	// blok pengkondisian pertam terhadapa minScore
	if averageLumba >= minScore && averageKoala >= minScore {
		
		if averageLumba > averageKoala { // kondisi score team lumba > team koala
			return "Tim Lumba-lumba"
		} else if averageLumba < averageKoala { // kondisi score team koala > team lumba
			return "Tim Koala"
		} else {
			return "Seri"
		}
	} else if averageLumba >= minScore {
		return "Tim Lumba-lumba"
	} else if averageKoala >= minScore {
		return "Tim Koala"
	} else {
		return "Tidak ada pemenang karena kedua tim tidak mencapai skor minimum"
	}
}

func main() {
	// pembuatan object dan assignment values 
	var teamLumba1 = Team{name: "Tim Lumba-lumba", scores: []int{96, 108, 89}}
	var teamKoala1 = Team{name: "Tim Koala", scores: []int{88, 91, 110}}

	var teamLumba2 = Team{name: "Tim Lumba-lumba", scores: []int{97, 112, 101}}
	var teamKoala2 = Team{name: "Tim Koala", scores: []int{109, 95, 123}}

	var teamLumba3 = Team{name: "Tim Lumba-lumba", scores: []int{97, 112, 101}}
	var teamKoala3 = Team{name: "Tim Koala", scores: []int{109, 95, 106}}

	// define minimum score
	minScore := 100

	// memanggil fungsi determineWinner dengan passing parameter object team dan minimum score
	winner1 := determineWinner(teamLumba1, teamKoala1, minScore)
	fmt.Println("Data 1:", winner1)

	
	winner2 := determineWinner(teamLumba2, teamKoala2, minScore)
	fmt.Println("Data Bonus 1:", winner2)

	
	winner3 := determineWinner(teamLumba3, teamKoala3, minScore)
	fmt.Println("Data Bonus 2:", winner3)
}
