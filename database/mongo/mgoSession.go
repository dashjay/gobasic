package mongo

import (
	"github.com/globalsign/mgo"

	"github.com/dashjay/gobasic/log"
)

// MgoSession is a tool to connect the mongodb use mgo
type MgoSession struct {
	s      *mgo.Session
	dbname string
}

func NewDefaultMgoSession(config *mgo.DialInfo) *MgoSession {
	var (
		err   error
		Mongo *mgo.Session
	)
	Mongo, err = mgo.DialWithInfo(config)
	if err != nil {
		panic(err)
	}
	Mongo.SetMode(mgo.Monotonic, true)
	log.Infof("Connect MongoDB success with addresses:%v", config.Addrs)
	return &MgoSession{s: Mongo, dbname: config.Database}
}

// CopySession can copy a connection but reuse the auth info....
func (m *MgoSession) CopySession() *MgoSession {
	return &MgoSession{s: m.s.Copy(), dbname: m.dbname}
}

// Close terminate the connection.
func (m *MgoSession) Close() {
	m.s.Close()
}

// C directly get collection from session.
func (m *MgoSession) C(cname string) *mgo.Collection {
	return m.s.DB(m.dbname).C(cname)
}
