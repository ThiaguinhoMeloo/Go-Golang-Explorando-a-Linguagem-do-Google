package systemsandtypes

type Car struct {
	Name         string
	CurrentSpeed int
}

type Ferrari struct {
	Car     // campo anonimo
	TurboOn bool
}

// func main() {
// 	f := systemsandtypes.Ferrari{}
// 	f.Name = "F40"
// 	f.CurrentSpeed = 0
// 	f.TurboOn = true

// 	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.Name, f.TurboOn)
// }
