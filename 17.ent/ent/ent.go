package ent

// nilableの設定の仕方について
type User struct {
	RequiredName string  `json:"required_name,omitempty"`
	OptionalName string  `json:"optional_name,omitempty"`
	NillableName *string `json:"nillable_name,omitempty"`
}
