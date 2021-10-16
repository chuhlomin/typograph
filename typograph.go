package main

type Typograph struct {
}

func NewTypograph() *Typograph {
	return &Typograph{}
}

func (t *Typograph) Process(in string) (out string) {
	return Symbols(in)
}
