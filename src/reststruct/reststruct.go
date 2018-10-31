package reststruct

type StructForREST struct {
	ArenaID   int     `json:"arenaID"`
	AccountID int     `json:"accountID"`
	ParryType string  `json:"parryTypeID"`
	BetValue  float32 `json:"betValue"`
}
