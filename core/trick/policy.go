package trick

import "time"

type TrickPolicy struct {
	delay    time.Duration
	keep     time.Duration
	interval time.Duration
}

func (p *TrickPolicy) Apply(func fail, func recover) {
	time.AfterFunc(p.delay, func() {
		fail()
		time.Sleep(p.keep)
		recover()
	})	
}
