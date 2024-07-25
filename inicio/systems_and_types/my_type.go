package systemsandtypes

type Note float64

func (n Note) InBetween(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

// func NoteForConcept(n Note) string {
// 	if n.InBetween(9.0, 10.0) {
// 		return "A"
// 	} else if n.InBetween(7.0, 8.99) {
// 		return "B"
// 	} else if n.InBetween(5.0, 7.99) {
// 		return "C"
// 	} else if n.InBetween(3.0, 4.99) {
// 		return "D"
// 	} else {
// 		return "E"
// 	}
// }

func NoteForConcept(n Note) string {
	switch {
	case n.InBetween(9.0, 10.0):
		return "A"
	case n.InBetween(7.0, 8.99):
		return "B"
	case n.InBetween(5.0, 7.99):
		return "C"
	case n.InBetween(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

// func main() {
// 	fmt.Println(systemsandtypes.NoteForConcept(9.8))
// 	fmt.Println(systemsandtypes.NoteForConcept(6.9))
// 	fmt.Println(systemsandtypes.NoteForConcept(2.1))
// }
