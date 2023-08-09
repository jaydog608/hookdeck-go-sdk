// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type CreateDestinationRequest struct {
	// Name for the destination <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the destination
	Description *string `json:"description,omitempty"`
	// Endpoint of the destination
	Url *string `json:"url,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *CreateDestinationRequestRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Limit event attempts to receive per period
	RateLimit              *int                         `json:"rate_limit,omitempty"`
	HttpMethod             *DestinationHttpMethod       `json:"http_method,omitempty"`
	AuthMethod             *DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	PathForwardingDisabled *bool                        `json:"path_forwarding_disabled,omitempty"`
}

type GetDestinationsRequest struct {
	Id         *string                    `json:"-"`
	Name       *string                    `json:"-"`
	Archived   *bool                      `json:"-"`
	ArchivedAt *time.Time                 `json:"-"`
	Url        *string                    `json:"-"`
	CliPath    *string                    `json:"-"`
	OrderBy    *string                    `json:"-"`
	Dir        *GetDestinationsRequestDir `json:"-"`
	Limit      *int                       `json:"-"`
	Next       *string                    `json:"-"`
	Prev       *string                    `json:"-"`
}

type UpdateDestinationRequest struct {
	// Name for the destination <span style="white-space: nowrap">`<= 155 characters`</span>
	Name *string `json:"name,omitempty"`
	// Description for the destination
	Description *string `json:"description,omitempty"`
	// Endpoint of the destination
	Url *string `json:"url,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *UpdateDestinationRequestRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Limit event attempts to receive per period
	RateLimit              *int                         `json:"rate_limit,omitempty"`
	HttpMethod             *DestinationHttpMethod       `json:"http_method,omitempty"`
	AuthMethod             *DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	PathForwardingDisabled *bool                        `json:"path_forwarding_disabled,omitempty"`
}

type UpsertDestinationRequest struct {
	// Name for the destination <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the destination
	Description *string `json:"description,omitempty"`
	// Endpoint of the destination
	Url *string `json:"url,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *UpsertDestinationRequestRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Limit event attempts to receive per period
	RateLimit              *int                         `json:"rate_limit,omitempty"`
	HttpMethod             *DestinationHttpMethod       `json:"http_method,omitempty"`
	AuthMethod             *DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	PathForwardingDisabled *bool                        `json:"path_forwarding_disabled,omitempty"`
}
