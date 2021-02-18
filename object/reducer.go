package object

type SingleReducer struct {
	SimpleNamer
	SimpleRotator

	ratio float64

	sockets [2]*SimpleRotatorSocket
}

func NewSingleReducer(name string, ratio float64) *SingleReducer {
	r := new(SingleReducer)
	r.name = name
	r.ratio = ratio
	r.sockets[0] = NewSimpleRotatorSocket(r)
	r.sockets[1] = NewSimpleRotatorSocket(r)

	return r
}

func (r *SingleReducer) String() string {
	return r.NamedString(r)
}

func (r *SingleReducer) Sockets() []Socket {
	return []Socket{r.sockets[0], r.sockets[1]}
}

func (r *SingleReducer) Disband() {
	r.sockets[0].Disband()
	r.sockets[1].Disband()
}

func (r *SingleReducer) SetController(c bool) {
	// 不会成为控制焦点
}

func (r *SingleReducer) IsController() bool {
	return false
}

func (r *SingleReducer) Ratio() float64 {
	return r.ratio
}
