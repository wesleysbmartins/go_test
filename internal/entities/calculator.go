package entities

type Calculator struct {
	A      int `json:"a"`
	B      int `json:"b"`
	Result int `json:"result"`
}

type ICalculator interface {
	Sum()
	Subtraction()
	Multiplication()
	Division()
}

func (c *Calculator) Sum() {
	c.Result = c.A + c.B
}

func (c *Calculator) Subtraction() {
	c.Result = c.A - c.B
}

func (c *Calculator) Multiplication() {
	c.Result = c.A * c.B
}

func (c *Calculator) Division() {
	c.Result = c.A / c.B
}
