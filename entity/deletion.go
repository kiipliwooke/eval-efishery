package entity

type DeletionLoan struct {
	Documents       bool `json:"documents"`
	LoanApplication bool `json:"loanApplication"`
	Loan            bool `json:"loan"`
}
