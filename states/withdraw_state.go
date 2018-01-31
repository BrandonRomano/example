package states

import (
	"strconv"

	"github.com/fsm/fsm"
)

func GetWithdrawState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateWithdraw,
		EntryAction: func() error {
			emitter.Emit("Ok I can help you with that, how much would you like to withdraw?")
			return nil
		},
		ReentryAction: func() error {
			emitter.Emit("I'm not quite sure I understand...")
			emitter.Emit("How much would you like to withdraw?")
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			switch v := input.(type) {
			case string:
				amnt, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return nil
				}
				traverser.Upsert(varWithdrawAmount, amnt)
				return GetWithdrawResultState(emitter, traverser)
			}
			return nil
		},
	}
}
