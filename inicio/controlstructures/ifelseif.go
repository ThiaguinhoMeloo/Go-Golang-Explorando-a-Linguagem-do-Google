package controlstructures

func NoteForConcept(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 6 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}

	// fmt.Println(controlstructures.NoteForConcept(9.8))
	// fmt.Println(controlstructures.NoteForConcept(6.9))
	// fmt.Println(controlstructures.NoteForConcept(2.1))
}
