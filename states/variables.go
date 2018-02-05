package states

const (
	varCashInHand        = "cash-in-hand"
	varAccountBalance    = "account-balance"
	varTransactionAmount = "transaction-amount"

	stateStart          = "start"
	stateWithdraw       = "withdraw"
	stateWithdrawResult = "withdraw-result"
	stateDeposit        = "deposit"
	stateDepositResult  = "deposit-result"
	stateViewBalance    = "view-balance"
	stateEnterBank      = "enter-bank"
	stateReenterBank    = "reenter-bank"
)
