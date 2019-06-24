package trick

import "time"
import log "github.com/sirupsen/logrus"

type TrickPolicy struct {
	Delay    int64 `yaml:"delay"`
	Keep     int64 `yaml:"keep"`
	Interval int64 `yaml:"interval"`
}

func (p *TrickPolicy) Apply(fail func() error, recover func() error) {
	time.AfterFunc(time.Millisecond * time.Duration(p.Delay), func() {
		if err := fail(); err != nil {
			log.Error("apply fail function error, ", err)
		}
		time.Sleep(time.Millisecond * time.Duration(p.Keep))
		if err := recover(); err != nil {
			log.Error("apply recover function error, ", err)
		}

	})
}
