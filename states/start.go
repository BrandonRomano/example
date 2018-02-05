package states

import (
	"github.com/fsm/fsm"
)

func GetStartState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateStart,
		EntryAction: func() error {
			traverser.Upsert(varCashInHand, 60)
			traverser.Upsert(varAccountBalance, 500)
			return nil
		},
		ReentryAction: func() error {
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			return GetEnterBankState(emitter, traverser)
		},
	}
}
