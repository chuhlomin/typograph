package typograph

type Typograph struct {
}

type Processor interface {
	Process(in []byte) []byte
}

func NewTypograph() *Typograph {
	return &Typograph{}
}

func (t *Typograph) Process(in []byte) []byte {
	pipeline := []Processor{
		&Symbols{},
		&Space{},
		&NBSP{},
	}

	for _, p := range pipeline {
		in = p.Process(in)
	}

	return in
}
