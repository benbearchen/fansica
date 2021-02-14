package object

type PCU struct {
	SimpleElectric

	sockets []*SimpleElectricSocket
}

func NewPCU(n int) *PCU {
	pcu := new(PCU)
	s := make([]*SimpleElectricSocket, n)
	for i := range s {
		s[i] = NewSimpleElectricSocket(pcu)
	}

	pcu.sockets = s

	return pcu
}

func (pcu *PCU) Sockets() []Socket {
	sockets := make([]Socket, len(pcu.sockets))
	for i := range sockets {
		sockets[i] = pcu.sockets[i]
	}

	return sockets
}

func (pcu *PCU) Disband() {
	for _, s := range pcu.sockets {
		s.Disband()
	}
}

func (pcu *PCU) SetController(c bool) {
	// 假定 PCU 不会成为控制焦点
}

func (pcu *PCU) IsController() bool {
	return false
}
