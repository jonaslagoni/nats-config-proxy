
package api

// PermissionRules represents a PermissionRules model.
type PermissionRules struct {
  Allow []string `json:"allow,omitempty"`
  Deny []string `json:"deny,omitempty"`
}