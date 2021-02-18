package object

func SelectRotatorSocket(a Object, index int) (RotatorSocket, error) {
	ss := a.Sockets()
	c := 0
	var first RotatorSocket
	for _, s := range ss {
		r, ok := s.(RotatorSocket)
		if !ok {
			continue
		}

		if index < 0 && first == nil {
			first = r
		}

		if c == index {
			return r, nil
		} else {
			c++
		}
	}

	if c == 0 {
		return nil, NoSocketError
	}

	if index < 0 {
		if c == 1 {
			return first, nil
		} else {
			return nil, MultiSocketError
		}
	} else {
		return nil, NoSocketError
	}
}

func SelectOtherRotatorSocket(a Object, other RotatorSocket) (RotatorSocket, error) {
	ss := a.Sockets()
	c := 0
	var first RotatorSocket
	for _, s := range ss {
		r, ok := s.(RotatorSocket)
		if !ok {
			continue
		}

		if r == other {
			continue
		}

		if first == nil {
			first = r
		}

		c++
	}

	if c == 0 {
		return nil, NoSocketError
	}

	if c == 1 {
		return first, nil
	} else {
		return nil, MultiSocketError
	}
}

func ConnectRotator(a, b Object) error {
	return ConnectRotatorI(a, b, -1)
}

func ConnectRotatorI(a, b Object, indexOfA int) error {
	ra, err := SelectRotatorSocket(a, indexOfA)
	if err != nil {
		return err
	}

	rb, err := SelectRotatorSocket(b, -1)
	if err != nil {
		return err
	}

	return ra.Connect(rb)
}

func ChanRotator(a, b Object, r ...Reducer) error {
	socket, err := SelectRotatorSocket(a, -1)
	if err != nil {
		return err
	}

	for _, c := range r {
		s0, err := SelectRotatorSocket(c, 0)
		if err != nil {
			return err
		}

		s1, err := SelectRotatorSocket(c, 1)
		if err != nil {
			return err
		}

		err = socket.Connect(s0)
		if err != nil {
			return err
		}

		socket = s1
	}

	socketB, err := SelectRotatorSocket(b, -1)
	if err != nil {
		return err
	}

	return socket.Connect(socketB)
}

func CountRotatorChanController(a Object) (count int, first Object) {
	m := make(map[Object]bool)
	m[a] = a.IsController()
	ss := a.Sockets()
	for len(ss) > 0 {
		var s Socket
		s, ss = ss[0], ss[1:]
		if !s.IsConnected() {
			continue
		}

		r, ok := s.(RotatorSocket)
		if !ok {
			continue
		}

		t := r.Target().Source()
		if _, ok = m[t]; !ok {
			m[t] = t.IsController()
			ss = append(ss, t.Sockets()...)
		}
	}

	for obj, ctrl := range m {
		if !ctrl {
			continue
		}

		count++
		if first == nil {
			first = obj
		}
	}

	return
}

func SelectElectricSocket(a Object, index int) (ElectricSocket, error) {
	ss := a.Sockets()
	c := 0
	var first ElectricSocket
	for _, s := range ss {
		e, ok := s.(ElectricSocket)
		if !ok {
			continue
		}

		if index < 0 && first == nil {
			first = e
		}

		if c == index {
			return e, nil
		} else {
			c++
		}
	}

	if c == 0 {
		return nil, NoSocketError
	}

	if index < 0 {
		if c == 1 {
			return first, nil
		} else {
			return nil, MultiSocketError
		}
	} else {
		return nil, NoSocketError
	}
}

func ConnectElectric(a, b Object, indexOfA int) error {
	ea, err := SelectElectricSocket(a, indexOfA)
	if err != nil {
		return err
	}

	eb, err := SelectElectricSocket(b, -1)
	if err != nil {
		return err
	}

	return ea.Connect(eb)
}
