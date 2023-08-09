// This file was auto-generated by Fern from our API Definition.

package api

type CreateIntegrationRequest struct {
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs  *CreateIntegrationRequestConfigs `json:"configs,omitempty"`
	Provider *IntegrationProvider             `json:"provider,omitempty"`
	// List of features to enable (see features list above)
	Features []IntegrationFeature `json:"features,omitempty"`
}

type GetIntegrationsRequest struct {
	Label    *string              `json:"-"`
	Provider *IntegrationProvider `json:"-"`
}

type UpdateIntegrationRequest struct {
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs  *UpdateIntegrationRequestConfigs `json:"configs,omitempty"`
	Provider *IntegrationProvider             `json:"provider,omitempty"`
	// List of features to enable (see features list above)
	Features []IntegrationFeature `json:"features,omitempty"`
}
