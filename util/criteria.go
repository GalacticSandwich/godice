
package util

import (
	"slices"
)

func ( set DiceSet ) valid_groups ( groups int, members int ) bool {
	gmap := make( map [ int ] int )

	for i := 0; i < DICE_NUM; i++ {
		dval := set.die[ i ]
		val, exists := gmap[ dval ]

		gmap[ dval ] = 1
		if exists {
			gmap[ dval ] = val + 1
		}
	}

	gcount := 0

	for _, val := range gmap {
		if val >= members {
			gcount++
		}
	}

	return gcount >= groups
}

// func ( set DiceSet ) valid_group ( members int ) bool {
// 
// }

func ( set DiceSet ) valid_straight ( size int ) bool {
	dvals := set.die
	slices.Sort( dvals )

	numord := 0
	for i := 1; i < DICE_NUM; i++ {
		if numord == 0 || ( dvals[ i ] - dvals[ i - 1 ] ) == 1 {
			numord++
		}

		else {
			numord = 0
		}

		if numord == size {
			return true
		}
	}

	return false
}