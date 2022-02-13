
package api

// Identity represents a Identity model.
type Identity struct {
  Username string `json:"username,omitempty"`
  Password string `json:"password,omitempty"`
  Nkey string `json:"nkey,omitempty"`
  Permissions string `json:"permissions,omitempty"`
}