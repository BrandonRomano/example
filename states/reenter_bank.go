package states

import (
	"strings"

	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
)

func GetReenterBankState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateReenterBank,
		EntryAction: func() error {
			emitter.Emit(emitable.QuickReply{
				Message:       "Is there anything else I can do for you today?",
				Replies:       []string{"withdraw", "deposit", "view balance"},
				RepliesFormat: "You can either %v.",
			})
			return nil
		},
		ReentryAction: func() error {
			emitter.Emit("That's not something I can do!")
			emitter.Emit(emitable.QuickReply{
				Message:       "Just let me know if there's anything else I can do to help.",
				Replies:       []string{"withdraw", "deposit", "view balance"},
				RepliesFormat: "You can either %v.",
			})
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			switch v := input.(type) {
			case string:
				switch strings.ToLower(v) {
				case "withdraw":
					return GetWithdrawState(emitter, traverser)
				case "deposit":
					return nil // TODO
				case "balance":
					fallthrough
				case "view balance":
					return nil // TODO
				}
			}
			return nil
		},
	}
}
