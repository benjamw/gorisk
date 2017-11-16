package model

import (
	"time"

	"github.com/benjamw/gorisk/risk"
)

// Player State Values
const (
	_ = iota // Waiting (in game.go)
	_        // Placing (in game.go)
	Trading
	Attacking
	Occupying
	Fortifying
	Resigned
	Dead
)

type GamePlayer struct {
	Player
	order     int
	cards     []risk.Card
	armies    int
	state     int
	extraInfo extraPlayerInfo
	moveDate  time.Time
}

type extraPlayerInfo struct {
	conquered int
	forced    bool
	getCard   bool
	occupy    bool
	round     int
	turn      int
}
