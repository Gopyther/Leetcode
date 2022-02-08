package main

import "fmt"

// The actual code for creating an Array to hold DVD's.
var dvdCollection = [15]DVD{}

// A simple definition for a DVD.
type DVD struct {
	name        string
	releaseYear int
	director    string
}

func (d *DVD) AddDVD(name string, releaseYear int, director string) {
	d.name = name
	d.releaseYear = releaseYear
	d.director = director
}

func (d DVD) String() string {
	if d.name != "" {
		s := fmt.Sprintf("%s, directed by %s, released in %d", d.name, d.director, d.releaseYear)
		return s
	}
	return "null"
}

func main() {
	// Firstly, we need to actually create a DVD object for The Avengers.
	avengersDVD := DVD{"The Avengers", 2012, "Joss Whedon"}
	// Next, we'll put it into the 8th place of the Array. Remember, because we
	// started numbering from 0, the index we want is 7.
	dvdCollection[7] = avengersDVD

	incrediblesDVD := DVD{"The Incredibles", 2004, "Brad Bird"}
	findingDoryDVD := DVD{"Finding Dory", 2016, "Andrew Stanton"}
	lionKingDVD := DVD{"The Lion King", 2019, "Jon Favreau"}

	// Put "The Incredibles" into the 4th place: index 3.
	dvdCollection[3] = incrediblesDVD

	// Put "Finding Dory" into the 10th place: index 9.
	dvdCollection[9] = findingDoryDVD

	// Put "The Lion King" into the 3rd place: index 2.
	dvdCollection[2] = lionKingDVD

	starWarsDVD := DVD{"Star Wars", 1977, "George Lucas"}
	dvdCollection[3] = starWarsDVD

	// Print out what's in indexes 7, 10, and 3.
	fmt.Println(dvdCollection[7])
	fmt.Println(dvdCollection[10])
	fmt.Println(dvdCollection[3])

	// Will print:

	// The Avengers, directed by Joss Whedon, released in 2012
	// null
	// Star Wars, directed by George Lucas, released in 1977
}
