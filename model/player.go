package model

import (
	"time"
)

type Player struct {
	Base
	Username     string
	Name         string
	Timezone     time.Location
	Wins         int
	Kills        int
	Losses       int
	LastOnline   time.Time
	email        string
	preferences  preferences
	passwordHash string
	admin        bool
	loggedIn     bool `datastore:"-"`
	currentGames int  `datastore:"-"`
}

type preferences struct {
	allowEmail   bool
	inviteOptOut bool
	maxGames     int
}
