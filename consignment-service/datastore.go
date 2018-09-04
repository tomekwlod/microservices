package main

import (
	"gopkg.in/mgo.v2"
)

// Copy and Clone methods are very similar, but have a subtle, but important difference. Clone re-uses the same
// socket as master. Which reduces the overhead of spawning an entirely new socket. Which is perfect for fast write
// performance. However, longer operations, such as more complex queries or big data jobs etc, may cause blocking in
// other go routines attempting to use this socket.

// CreateSession creates the main session to our mongodb instance
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
