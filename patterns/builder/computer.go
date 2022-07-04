package main

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type BuilderI interface {
	CPU(val string) BuilderI
	RAM(val int) BuilderI
	MB(val string) BuilderI

	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() BuilderI {
	return &computerBuilder{}
}

func (c *computerBuilder) CPU(val string) BuilderI {
	c.cpu = val
	return c
}
func (c *computerBuilder) RAM(val int) BuilderI {
	c.ram = val
	return c
}
func (c *computerBuilder) MB(val string) BuilderI {
	c.mb = val
	return c
}

func (c *computerBuilder) Build() Computer {
	return Computer{
		CPU: c.cpu,
		RAM: c.ram,
		MB:  c.mb,
	}
}

type officeComputerBuilder struct {
	computerBuilder
}

func NewOfficeComputerBuilder() BuilderI {
	return &officeComputerBuilder{}
}

func (o *officeComputerBuilder) Build() Computer {
	if o.cpu == "" {
		o.cpu = "Celeron"
	}
	if o.ram == 0 {
		o.ram = 2
	}

	if o.mb == "" {
		o.mb = "ASRock"
	}

	return Computer{
		CPU: o.cpu,
		RAM: o.ram,
		MB:  o.mb,
	}
}
