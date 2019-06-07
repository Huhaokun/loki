package trick

import "time"
import log "github.com/sirupsen/logrus"

type TrickPolicy struct {
	Delay    time.Duration `yaml:"delay"`
	Keep     time.Duration `yaml:"keep"`
	Interval time.Duration `yaml:"interval"`
}

func (p *TrickPolicy) Apply(fail func() error, recover func() error) {
	time.AfterFunc(p.Delay, func() {
		if err := fail(); err != nil {
			log.Error("apply fail function error, ", err)
		}
		time.Sleep(p.Keep)
		if err := recover(); err != nil {
			log.Error("apply recover function error, ", err)
		}

	})
}
