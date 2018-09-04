package main

import mgo "gopkg.in/mgo.v2"

// CreateSession creates a master mongo session which later should be Cloned/Copied for your needs
func CreateSession(host string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
