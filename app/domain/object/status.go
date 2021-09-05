package object

type (
	Status struct {
		// The internal ID of the status
		ID int64 `json:"id"`

		// The accountId of the status
		AccountID AccountID `json:"account_id,omitempty" db:"account_id"`

		// The content of the status
		Content *string `json:"content,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
