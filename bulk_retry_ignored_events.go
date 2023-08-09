// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type CreateIgnoredEventBulkRetryRequest struct {
	// Filter by the bulk retry ignored event query object
	Query *CreateIgnoredEventBulkRetryRequestQuery `json:"query,omitempty"`
}

type GetIgnoredEventBulkRetriesRequest struct {
	CancelledAt       *time.Time                            `json:"-"`
	CompletedAt       *time.Time                            `json:"-"`
	CreatedAt         *time.Time                            `json:"-"`
	Id                *string                               `json:"-"`
	QueryPartialMatch *bool                                 `json:"-"`
	InProgress        *bool                                 `json:"-"`
	OrderBy           *string                               `json:"-"`
	Dir               *GetIgnoredEventBulkRetriesRequestDir `json:"-"`
	Limit             *int                                  `json:"-"`
	Next              *string                               `json:"-"`
	Prev              *string                               `json:"-"`
}
