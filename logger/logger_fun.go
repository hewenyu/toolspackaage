package logger

// INFO ------------------------

/*
INFO
*/
func Info(args ...interface{}) {
	h.Info(args)
}

/*
INFO
*/
func (h *Hliog) Info(args ...interface{}) {
	h.base.Info(args)
}

// DEBUG ------------------------

/*
DEBUG
*/
func Debug(args ...interface{}) {
	h.Debug(args)
}

/*
DEBUG
*/
func (h *Hliog) Debug(args ...interface{}) {
	h.base.Debug(args)
}

// func (s *HsLogs) With(args ...interface{}) *Hliog {
// 	return &Hliog{base: s.base.With(s.sweetenFields(args)...)}
// }

// // Debug uses fmt.Sprint to construct and log a message.
// func (s *HsLogs) Debug(args ...interface{}) {
// 	s.log(DebugLevel, "", args, nil)
// }

// // Info uses fmt.Sprint to construct and log a message.
// func (s *SugaredLogger) Info(args ...interface{}) {
// 	s.log(InfoLevel, "", args, nil)
// }

// // Warn uses fmt.Sprint to construct and log a message.
// func (s *SugaredLogger) Warn(args ...interface{}) {
// 	s.log(WarnLevel, "", args, nil)
// }

// // Error uses fmt.Sprint to construct and log a message.
// func (s *SugaredLogger) Error(args ...interface{}) {
// 	s.log(ErrorLevel, "", args, nil)
// }

// // DPanic uses fmt.Sprint to construct and log a message. In development, the
// // logger then panics. (See DPanicLevel for details.)
// func (s *SugaredLogger) DPanic(args ...interface{}) {
// 	s.log(DPanicLevel, "", args, nil)
// }

// // Panic uses fmt.Sprint to construct and log a message, then panics.
// func (s *SugaredLogger) Panic(args ...interface{}) {
// 	s.log(PanicLevel, "", args, nil)
// }

// // Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
// func (s *SugaredLogger) Fatal(args ...interface{}) {
// 	s.log(FatalLevel, "", args, nil)
// }

// // Debugf uses fmt.Sprintf to log a templated message.
// func (s *SugaredLogger) Debugf(template string, args ...interface{}) {
// 	s.log(DebugLevel, template, args, nil)
// }

// // Infof uses fmt.Sprintf to log a templated message.
// func (s *SugaredLogger) Infof(template string, args ...interface{}) {
// 	s.log(InfoLevel, template, args, nil)
// }

// // Warnf uses fmt.Sprintf to log a templated message.
// func (s *SugaredLogger) Warnf(template string, args ...interface{}) {
// 	s.log(WarnLevel, template, args, nil)
// }

// // Errorf uses fmt.Sprintf to log a templated message.
// func (s *SugaredLogger) Errorf(template string, args ...interface{}) {
// 	s.log(ErrorLevel, template, args, nil)
// }

// // DPanicf uses fmt.Sprintf to log a templated message. In development, the
// // logger then panics. (See DPanicLevel for details.)
// func (s *SugaredLogger) DPanicf(template string, args ...interface{}) {
// 	s.log(DPanicLevel, template, args, nil)
// }

// // Panicf uses fmt.Sprintf to log a templated message, then panics.
// func (s *SugaredLogger) Panicf(template string, args ...interface{}) {
// 	s.log(PanicLevel, template, args, nil)
// }

// // Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
// func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {
// 	s.log(FatalLevel, template, args, nil)
// }

// // Debugw logs a message with some additional context. The variadic key-value
// // pairs are treated as they are in With.
// //
// // When debug-level logging is disabled, this is much faster than
// //  s.With(keysAndValues).Debug(msg)
// func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {
// 	s.log(DebugLevel, msg, nil, keysAndValues)
// }

// // Infow logs a message with some additional context. The variadic key-value
// // pairs are treated as they are in With.
// func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {
// 	s.log(InfoLevel, msg, nil, keysAndValues)
// }

// // Warnw logs a message with some additional context. The variadic key-value
// // pairs are treated as they are in With.
// func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {
// 	s.log(WarnLevel, msg, nil, keysAndValues)
// }

// // Errorw logs a message with some additional context. The variadic key-value
// // pairs are treated as they are in With.
// func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {
// 	s.log(ErrorLevel, msg, nil, keysAndValues)
// }

// // DPanicw logs a message with some additional context. In development, the
// // logger then panics. (See DPanicLevel for details.) The variadic key-value
// // pairs are treated as they are in With.
// func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {
// 	s.log(DPanicLevel, msg, nil, keysAndValues)
// }

// // Panicw logs a message with some additional context, then panics. The
// // variadic key-value pairs are treated as they are in With.
// func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {
// 	s.log(PanicLevel, msg, nil, keysAndValues)
// }

// // Fatalw logs a message with some additional context, then calls os.Exit. The
// // variadic key-value pairs are treated as they are in With.
// func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {
// 	s.log(FatalLevel, msg, nil, keysAndValues)
// }

// // Sync flushes any buffered log entries.
// func (s *SugaredLogger) Sync() error {
// 	return s.base.Sync()
// }
