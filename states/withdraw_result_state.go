package states

import (
	"fmt"

	"github.com/fsm/fsm"
)

func GetWithdrawResultState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateWithdrawResult,
		EntryAction: func() error {
			// Get Values
			withdraw, _ := traverser.Fetch(varWithdrawAmount)
			balance, _ := traverser.Fetch(varAccountBalance)

			// Update balance
			if withdraw.(int) < balance.(int) {
				// Update balance
				newBalance := balance.(int) - withdraw.(int)
				traverser.Upsert(varAccountBalance, newBalance)

				// Prompt user
				emitter.Emit(fmt.Sprintf("Okay, here is $%v", withdraw))
				emitter.Emit(fmt.Sprintf("Your balance is now $%v", newBalance))
			} else {
				// Not enough cash
				emitter.Emit("Sorry, your account has insufficent funds to complete the transaction.")
			}

			// Switch to reenter bank
			traverser.Delete(varWithdrawAmount)
			traverser.SetCurrentState(stateReenterBank)
			return nil
		},
		ReentryAction: func() error {
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			return nil
		},
	}
}
