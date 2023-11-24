
/*
	FILE: dice.go
	AUTHORS: Aaron N. (GalacticSandwich)

	...

	HISTORY:
		- 11/23/23 - File Created + Basic Code Added
*/

package util

import (
	"errors"
	"math/rand"
	"time"
)

const (
	// the number of dice in a set
	DICE_NUM = 5
	// the maximum value that case be rolled on a dice
	DICE_MAX_VAL = 6
)

// struct representing a set of die
type DiceSet struct {
	die		[ DICE_NUM ] int	// array representing values on each dice
	holds	[ DICE_NUM ] bool	// array representing the hold status of each dice
}

/*
	Reset a dice set, set all the dice numeric values to zero and reset all
	hold flags.
*/
func ( set *DiceSet ) reset () {
	// loop through and set each dice to zero, and the default hold status, 
	// which is false
	for i := 0; i < DICE_NUM; i++ {
		set.die[ i ] = 0
		set.holds[ i ] = false
	}
}

/*
	Generate a new empty dice set, and then return it.
*/
func make_set () DiceSet {
	// declare a new dice set
	var set DiceSet

	// reset the new set
	set.reset()

	// return the new set
	return set
}

/*
	Toggle a hold flag in a dice in a dice set. May return an error if an invalid
	index is provided, returns nil otherwise.
*/
func ( set *DiceSet ) toggle_hold ( n int ) error {
	// error check to see if the index is out of bounds
	if n >= DICE_MAX_VAL || n < 0 {
		return errors.New( "Dice index out of bounds" )
	}

	// error check to see if the dice has a null/unrolled value
	if set.die[ n ] == 0 {
		return errors.New( "Cannot hold unrolled dice" )
	}

	// set the dice hold flag to the opposite of its current, and return nil (for no error)
	set.holds[ n ] = !set.holds[ n ]

	return nil
}

/*
	Roll all unheld dice within a dice set.
*/
func ( set *DiceSet ) roll () {
	// get a new seed/source for the the random number generator
	src := rand.NewSource( time.Now().UnixNano() )
	// get a new random number generator using the new source
	gen := rand.New( src )

	// loop through the set, and set each unheld dice to a new random allowed value
	for i := 0; i < DICE_NUM; i++ {
		if !set.holds[ i ] {
			set.die[ i ] = gen.Intn( DICE_MAX_VAL ) + 1
		}
	}
}
