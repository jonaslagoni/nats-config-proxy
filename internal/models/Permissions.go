package models

// Permissions represents a Permissions model.
type Permissions struct {
	Publish   *PermissionRules `json:"publish,omitempty"`
	Subscribe *PermissionRules `json:"subscribe,omitempty"`
	Responses *PermissionRules `json:"responses,omitempty"`
}
