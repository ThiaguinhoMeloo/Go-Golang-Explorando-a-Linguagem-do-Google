package systemsandtypes

import "strings"

type People struct {
	Name     string
	LastName string
}

func (p People) GetFullName() string {
	return p.Name + " " + p.LastName
}

func (p *People) SetFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.Name = parts[0]
	p.LastName = parts[1]
}

// func main() {
// 	p1 := systemsandtypes.People{"Pedro", "Silva"}
// 	fmt.Println(p1.GetFullName())

// 	p1.SetFullName("Thiago Melo")
// 	fmt.Println(p1.GetFullName())
// }
