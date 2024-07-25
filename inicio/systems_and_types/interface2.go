package systemsandtypes

type Sports interface {
	TurnOnTurbo()
}

type Ferrari2 struct {
	Model        string
	TurboOn      bool
	CurrentSpeed int
}

func (f *Ferrari2) TurnOnTurbo() {
	f.TurboOn = true
}
