//proverb package generates a proverb in the format below, given a list of inputs, generate the relevant proverb. For example, given the list ["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"], you will output the full text of this proverbial rhyme:

// For want of a nail the shoe was lost.
// For want of a shoe the horse was lost.
// For want of a horse the rider was lost.
// For want of a rider the message was lost.
// For want of a message the battle was lost.
// For want of a battle the kingdom was lost.
// And all for the want of a nail.

// Note that the list of inputs may vary; your solution should be able to handle lists of arbitrary length and content. No line of the output text should be a static, unchanging string; all should vary according to the input given.
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	fmt.Println(rhyme)

	s := make([]string, 0)
	//case of 0 input
	if len(rhyme) == 0 {
		// s := make([]string, 0)
		// s = append(s, "")
		return rhyme
	}

	//case of n input
	if len(rhyme) > 0 {
		for i := 0; i < (len(rhyme) - 1); i++ {
			t := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
			s = append(s, t)
		}
		t := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		s = append(s, t)
	}

	return s

}
