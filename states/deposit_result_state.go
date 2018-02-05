package states

import (
	"fmt"

	"github.com/fsm/fsm"
)

func GetDepositResultState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateDepositResult,
		EntryAction: func() error {
			// Get Values
			depositAmount, _ := traverser.Fetch(varTransactionAmount)
			cashInHand, _ := traverser.Fetch(varCashInHand)
			accountBalance, _ := traverser.Fetch(varAccountBalance)

			// If we have enough cash in hand
			if depositAmount.(int) <= cashInHand.(int) {
				// Update cash in hand
				newCashInHand := cashInHand.(int) - depositAmount.(int)
				traverser.Upsert(varCashInHand, newCashInHand)

				// Update account balance
				newAccountBalance := accountBalance.(int) + depositAmount.(int)
				traverser.Upsert(varAccountBalance, newAccountBalance)

				// Prompt user
				emitter.Emit(fmt.Sprintf("Okay, $%v has been added to your balance.", depositAmount))
				emitter.Emit(fmt.Sprintf("Your account balance is now $%v.", newAccountBalance))
			} else {
				// Not enough cash
				emitter.Emit(fmt.Sprintf("Sorry, it looks like you only have $%v, so you can't deposit that much.", cashInHand))
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
