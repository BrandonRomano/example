package main

import (
	cachestore "github.com/fsm/cache-store"
	"github.com/fsm/cli"
	"github.com/fsm/example/states"
	"github.com/fsm/fsm"
)

func main() {
	startCLI()
}

func startCLI() {
	cli.Start(getStateMachine(), getStore())
}

func getStateMachine() fsm.StateMachine {
	return fsm.StateMachine{
		states.GetStartState,
		states.GetEnterBankState,
		states.GetWithdrawState,
		states.GetWithdrawResultState,
		states.GetReenterBankState,
	}
}

func getStore() fsm.Store {
	return cachestore.New()
}
