package logger

// func (s *HsLogs) With(args ...interface{}) *Hliog {
// 	return &Hliog{base: s.base.With(s.sweetenFields(args)...)}
// }

/*
INFO
*/
func (s *Hliog) Info(args ...interface{}) {
	s.base.Info(args)
}

/*
DEBUG
*/
func (s *Hliog) Debug(args ...interface{}) {
	s.base.Debug(args)
}

/*
Warn
*/
func (s *Hliog) Warn(args ...interface{}) {
	s.base.Warn(args)
}

/*
DPanic
*/
func (s *Hliog) DPanic(args ...interface{}) {
	s.base.DPanic(args)
}

/*
Error
*/
func (s *Hliog) Error(args ...interface{}) {
	s.base.Error(args)
}

/*
Panic
*/
func (s *Hliog) Panic(args ...interface{}) {
	s.base.Panic(args)
}

/*
Fatal
*/
func (s *Hliog) Fatal(args ...interface{}) {
	s.base.Fatal(args)
}

// ------------------------------------format---------------------------------------------------

/*
INFOf
*/
func (s *Hliog) Infof(template string, args ...interface{}) {
	s.base.Infof(template, args)
}

/*
DEBUGf
*/
func (s *Hliog) Debugf(template string, args ...interface{}) {
	s.base.Debugf(template, args)
}

/*
Warnf
*/
func (s *Hliog) Warnf(template string, args ...interface{}) {
	s.base.Warnf(template, args)
}

/*
DPanicf
*/
func (s *Hliog) DPanicf(template string, args ...interface{}) {
	s.base.DPanicf(template, args)
}

/*
Errorf
*/
func (s *Hliog) Errorf(template string, args ...interface{}) {
	s.base.Errorf(template, args)
}

/*
Panicf
*/
func (s *Hliog) Panicf(template string, args ...interface{}) {
	s.base.Panicf(template, args)
}

/*
Fatalf
*/
func (s *Hliog) Fatalf(template string, args ...interface{}) {
	s.base.Fatalf(template, args)
}

// func (s *SugaredLogger) Sync() error {
// 	return s.base.Sync()
// }
