package main

type Director struct {
	builder BuilderI
}

func NewDirector(builder BuilderI) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildComputer() Computer {
	return d.builder.Build()
}

func (d *Director) SetBuilder(builder BuilderI) {
	d.builder = builder
}
