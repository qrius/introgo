/* Peter initiates a lottery contract and puts 1Eth in it
John places a bet with 1 Eth and guesses a number "5"
Lottery contract now holds 2Eth
Peter executes "release" on the contract
contract executes an external JSON web service to generate a random number, then compare this with John's bet
if John guesses right, he gets back his own Eth as well as Peter's so he'd have 2Eth
if John guesses wrong, Peter gets back his own Eth as well as John's so Peter'd have Eth

*/

package main

type lottery []int

// func (l lottery) draw() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	// for i := range d {
// 	// 	newPosition := r.Intn(len(d) - 1)

// 	// 	d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }
