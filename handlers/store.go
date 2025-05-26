package handlers

import (
	"sync"

	"github.com/lacriman/tic-tac-toe/game"
)

var (
	games   = make(map[string]*game.Game)
	gameMux = sync.RWMutex{}
)
