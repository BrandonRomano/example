package states

import (
	"fmt"

	"github.com/fsm/fsm"
)

func GetViewBalanceState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateViewBalance,
		EntryAction: func() error {
			accountBalance, _ := traverser.Fetch(varAccountBalance)
			cashInHand, _ := traverser.Fetch(varCashInHand)
			emitter.Emit(fmt.Sprintf("You currently have an account balance of $%v.", accountBalance))
			emitter.Emit(fmt.Sprintf("You also have $%v in cash.", cashInHand))
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
