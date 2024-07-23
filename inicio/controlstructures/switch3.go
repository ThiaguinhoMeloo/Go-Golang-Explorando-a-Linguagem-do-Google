package controlstructures

func Type(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo não reconhecido."
	}

	// fmt.Println(controlstructures.Type(2.3))
	// fmt.Println(controlstructures.Type(1))
	// fmt.Println(controlstructures.Type("Teste"))
	// fmt.Println(controlstructures.Type(func() {}))
	// fmt.Println(controlstructures.Type(time.Now()))
}
