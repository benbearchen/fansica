package object

import "testing"

import (
	"fmt"
	"runtime"
)

func testCallerLine(d int) string {
	_, _, line, ok := runtime.Caller(d)
	if ok {
		return fmt.Sprintf("L%d: ", line)
	} else {
		return ""
	}
}

func TestSocketFunc(t *testing.T) {
	wheel, _ := ParseWheel("wheel0", "255/50R20")
	eng := EasyMakeEngine("eng0", 6500, 220, 96)
	b := EasyMakeBattery("bat", 40, 220)
	m := EasyMakeMotor("mg", 100, 240, 12000)
	pcu := NewPCU("control", 4)
	reducer := NewSingleReducer("sr", 3.104)

	checkSelectRotatorSocket := func(a Object, i int, s Socket, e error) {
		tag := testCallerLine(2)
		r, err := SelectRotatorSocket(a, i)
		if err != e {
			t.Errorf("%sSelectRotatorSocket(%v, %d) want error %v, return %v", tag, a, i, e, err)
		}

		if r != s {
			t.Errorf("%sSelectRotatorSocket(%v, %d) return unexpected socket", tag, a, i)
		}

		if r == nil {
			return
		}

		ss := a.Sockets()
		si := 0
		match := -1
		for _, socket := range ss {
			_, ok := socket.(RotatorSocket)
			if !ok {
				continue
			}

			if s == socket {
				if match >= 0 {
					t.Errorf("%s%v has duplicate sockets?", tag, a)
				} else if si == i {
					match = 1
				} else {
					match = 0
				}
			}

			si++
		}

		if match < 0 {
			t.Errorf("%sSelectRotatorSocket(%v, %d) return wild socket", tag, a, i)
		}
	}

	checkSelectRotatorSocket(wheel, -1, wheel.socket, nil)
	checkSelectRotatorSocket(wheel, 0, wheel.socket, nil)
	checkSelectRotatorSocket(wheel, 1, nil, NoSocketError)

	checkSelectRotatorSocket(eng, -1, eng.socket, nil)
	checkSelectRotatorSocket(eng, 0, eng.socket, nil)
	checkSelectRotatorSocket(eng, 1, nil, NoSocketError)

	checkSelectRotatorSocket(b, -1, nil, NoSocketError)
	checkSelectRotatorSocket(b, 0, nil, NoSocketError)

	checkSelectRotatorSocket(m, -1, m.socketR, nil)
	checkSelectRotatorSocket(m, 0, m.socketR, nil)
	checkSelectRotatorSocket(m, 1, nil, NoSocketError)

	checkSelectRotatorSocket(pcu, -1, nil, NoSocketError)
	checkSelectRotatorSocket(pcu, 0, nil, NoSocketError)

	checkSelectRotatorSocket(reducer, -1, nil, MultiSocketError)
	checkSelectRotatorSocket(reducer, 0, reducer.sockets[0], nil)
	checkSelectRotatorSocket(reducer, 1, reducer.sockets[1], nil)
	checkSelectRotatorSocket(reducer, 2, nil, NoSocketError)

	checkOtherRotator := func(a Object, other RotatorSocket, s RotatorSocket, e error) {
		tag := testCallerLine(2)
		r, err := SelectOtherRotatorSocket(a, other)
		if err != e {
			t.Errorf("%sSelectOtherRotatorSocket(%v, *) want error %v, return %v", tag, a, e, err)
		}

		if r != s {
			t.Errorf("%sSelectOtherRotatorSocket(%v, *) return unexpected socket", tag, a)
		}
	}

	checkOtherRotator(wheel, wheel.socket, nil, NoSocketError)
	checkOtherRotator(wheel, nil, wheel.socket, nil)
	checkOtherRotator(wheel, eng.socket, wheel.socket, nil)

	checkOtherRotator(eng, eng.socket, nil, NoSocketError)
	checkOtherRotator(eng, nil, eng.socket, nil)
	checkOtherRotator(eng, wheel.socket, eng.socket, nil)

	checkOtherRotator(b, nil, nil, NoSocketError)
	checkOtherRotator(b, wheel.socket, nil, NoSocketError)

	checkOtherRotator(m, m.socketR, nil, NoSocketError)
	checkOtherRotator(m, nil, m.socketR, nil)
	checkOtherRotator(m, wheel.socket, m.socketR, nil)

	checkOtherRotator(pcu, nil, nil, NoSocketError)
	checkOtherRotator(pcu, wheel.socket, nil, NoSocketError)

	checkOtherRotator(reducer, reducer.sockets[0], reducer.sockets[1], nil)
	checkOtherRotator(reducer, reducer.sockets[1], reducer.sockets[0], nil)
	checkOtherRotator(reducer, nil, nil, MultiSocketError)
	checkOtherRotator(reducer, wheel.socket, nil, MultiSocketError)
}
