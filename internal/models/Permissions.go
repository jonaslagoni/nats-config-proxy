package models

// Permissions represents a Permissions model.
type Permissions struct {
	Publish        *PermissionRules `json:"publish,omitempty"`
	Subscribe      *PermissionRules `json:"subscribe,omitempty"`
	AllowResponses *PermissionRules `json:"allow_responses,omitempty"`
}
