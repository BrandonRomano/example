package states

import (
	"strings"

	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
)

func GetEnterBankState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: stateEnterBank,
		EntryAction: func() error {
			emitter.Emit("Greetings valued customer, welcome to the First Bank of FSM!")
			emitter.Emit(emitable.QuickReply{
				Message:       "What can I do for you today?",
				Replies:       []string{"withdraw", "deposit", "view balance"},
				RepliesFormat: "You can either %v.",
			})
			return nil
		},
		ReentryAction: func() error {
			emitter.Emit("That's not something I can do!")
			emitter.Emit(emitable.QuickReply{
				Message:       "I can help you with the following options:",
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
