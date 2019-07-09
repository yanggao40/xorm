package xorm

import 	"xorm.io/core"


// Logger sets the logger on this session
func (session *Session) Logger(logger core.ILogger) *Session {
	session.logger = logger
	return session
}