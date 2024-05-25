package graph

import (
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/graph/model"
)

func APItoInner(api *model.ManageEntity) *core.ManageEntity {
	return &core.ManageEntity{
		RequestID:  api.RequestID,
		UserID:     api.UserID,
		Signer:     api.Signer,
		EntityType: api.EntityType,
		EntityID:   api.EntityID,
		Action:     api.Action,
		Metadata:   api.Metadata,
	}
}

func InnerToAPI(inner *core.ManageEntity) *model.ManageEntity {
	return &model.ManageEntity{
		RequestID:  inner.RequestID,
		UserID:     inner.UserID,
		Signer:     inner.Signer,
		EntityType: inner.EntityType,
		EntityID:   inner.EntityID,
		Action:     inner.Action,
		Metadata:   inner.Metadata,
	}
}

func NewAPItoAPI(newApi *model.NewManageEntity) *model.ManageEntity {
	return &model.ManageEntity{
		RequestID:  newApi.RequestID,
		UserID:     newApi.UserID,
		Signer:     newApi.Signer,
		EntityType: newApi.EntityType,
		EntityID:   newApi.EntityID,
		Metadata:   newApi.Metadata,
		Action:     newApi.Action,
	}
}
