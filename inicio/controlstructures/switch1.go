package controlstructures

func NoteForConceptSwitch1(n float64) string {
	var note = int(n)
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}

	// fmt.Println(controlstructures.NoteForConceptSwitch1(9.8))
	// fmt.Println(controlstructures.NoteForConceptSwitch1(6.9))
	// fmt.Println(controlstructures.NoteForConceptSwitch1(11.1))
}
