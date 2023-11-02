
package util

import (
	"math/rand"
	"time"
)

const (
	MaxSideVal = 6
	NumDice = 5
)

type DiceSet = struct {
	values	[ NumDice ] int
	holds	[ NumDice ] bool
	score	int
}

func BuildSet() DiceSet {
	ds := DiceSet { score: 0 }

	for i := 0; i < NumDice; i++ {
		ds.values[ i ] = 0
		ds.holds[ i ] = false
	}

	return ds
}

func ( ds DiceSet ) Roll( n int ) {
	if !ds.holds[ n ] {
		ds.values[ n ] = rand.Intn( MaxSideVal ) + 1
	}
}

func ( ds DiceSet ) RollEach() {
	for i := 0; i < NumDice; i++ {
		ds.Roll( i )
	}
}

func ( ds DiceSet ) Hold( n int ) {
	ds.holds[ n ] = !ds.holds[ n ]
}
