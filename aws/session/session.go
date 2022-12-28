package session

import "github.com/aws/aws-sdk-go/aws/session"

// NewDynamicSession is a factory to emit sessions
// relevent to the current config, ie:
//
// dev, local, sandbox, production
func NewDynamicSession() *session.Session {
	return session.Must(session.NewSession())
}
