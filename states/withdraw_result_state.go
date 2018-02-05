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
			withdrawAmount, _ := traverser.Fetch(varTransactionAmount)
			accountBalance, _ := traverser.Fetch(varAccountBalance)
			cashInHand, _ := traverser.Fetch(varCashInHand)

			// If there's enough money in the account
			if withdrawAmount.(int) <= accountBalance.(int) {
				// Update balance
				newAccountBalance := accountBalance.(int) - withdrawAmount.(int)
				traverser.Upsert(varAccountBalance, newAccountBalance)

				// Update cash in hand
				newCashInHand := cashInHand.(int) + withdrawAmount.(int)
				traverser.Upsert(varCashInHand, newCashInHand)

				// Prompt user
				emitter.Emit(fmt.Sprintf("Okay, here is $%v", withdrawAmount))
				emitter.Emit(fmt.Sprintf("Your balance is now $%v", newAccountBalance))
			} else {
				// Not enough cash
				emitter.Emit("Sorry, your account has insufficent funds to complete the transaction.")
			}

			// Switch to reenter bank
			traverser.Delete(varTransactionAmount)
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
