// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TransactionsResponse - OK
type TransactionsResponse struct {
	Data []Transaction `json:"data,omitempty"`
	// The maximum number of items that can be returned per page.
	PerPage *int64 `json:"per_page,omitempty"`
	// Could be used to retrieve the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Could be used to retrieve the previous page.
	PrevCursor *string `json:"prev_cursor,omitempty"`
}

func (o *TransactionsResponse) GetData() []Transaction {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TransactionsResponse) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *TransactionsResponse) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *TransactionsResponse) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}
