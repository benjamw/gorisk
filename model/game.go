package model

import (
	"time"

	"github.com/benjamw/gorisk/risk"
)

// Game State Values
const (
	Waiting = iota
	Placing
	Playing
	Finished
)

type Game struct {
	Base
	Name           string
	Players        map[int]Player
	AvailableCards []risk.Card
	State          int
	CurrentPlayer  int
	paused         bool
	passwordHash   string
	createDate     time.Time
	modifyDate     time.Time
	lastMove       time.Time
	capacity       int
	host           Player
	extraInfo      extraGameInfo
}

type extraGameInfo struct {
	fortify            bool
	multipleFortify    bool
	connectedFortify   bool
	placeInitialArmies bool
	initialArmyLimit   int
	kamikaze           bool
	warmonger          bool
	fowArmies          bool
	fowColors          bool
	tradeNumber        int
	customTrades       []tradeEntry
	tradeCardBonus     int
	conquer            conquer
	customRules        string
}

type conquer struct {
	watch        string // conquer_type from original WebRisk
	conquestsPer int
	perNumber    int
	skip         int
	startAt      int
	minimum      int
	maximum      int
}

type tradeEntry struct {
	start int
	end   int
	step  int
	times int
}
