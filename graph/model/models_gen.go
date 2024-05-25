// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ManageEntity struct {
	RequestID  string `json:"requestId"`
	UserID     int    `json:"userId"`
	Signer     string `json:"signer"`
	EntityType string `json:"entityType"`
	EntityID   int    `json:"entityId"`
	Metadata   string `json:"metadata"`
	Action     string `json:"action"`
}

type ManageEntityFilter struct {
	UserID     *int    `json:"userId,omitempty"`
	Signer     *string `json:"signer,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	EntityID   *int    `json:"entityId,omitempty"`
	Action     *string `json:"action,omitempty"`
}

type Mutation struct {
}

type NewManageEntity struct {
	RequestID  string `json:"requestId"`
	UserID     int    `json:"userId"`
	Signer     string `json:"signer"`
	EntityType string `json:"entityType"`
	EntityID   int    `json:"entityId"`
	Metadata   string `json:"metadata"`
	Action     string `json:"action"`
}

type Query struct {
}

type Subscription struct {
}
