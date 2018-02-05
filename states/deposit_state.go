package states

import (
	"strconv"

	"github.com/fsm/fsm"
)

func GetDepositState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateDeposit,
		EntryAction: func() error {
			emitter.Emit("Ok I can help you with that, how much would you like to deposit?")
			return nil
		},
		ReentryAction: func() error {
			emitter.Emit("I'm not quite sure I understand...")
			emitter.Emit("How much would you like to deposit?")
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			switch v := input.(type) {
			case string:
				amnt, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					return nil
				}
				traverser.Upsert(varTransactionAmount, int(amnt))
				return GetDepositResultState(emitter, traverser)
			}
			return nil
		},
	}
}
