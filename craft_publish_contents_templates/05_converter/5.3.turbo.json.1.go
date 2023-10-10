package controller

import (
	"time"

	"github.com/example/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostToPublishContentsTemplates(r *PublishContentsTemplatesBody, tenantID, userName string) model.PublishContentsTemplates {
	now := time.Now()

	contents := make(map[string]any)
	for key, value := range r.Contents {
		contents[key] = value
	}

	return model.PublishContentsTemplates{
		ID:          primitive.NewObjectID(),
		TenantID:    tenantID,
		DisplayName: r.DisplayName,
		Channel:     r.Channel.Channel,
		CreatedBy:   userName,
		CreatedAt:   now,
		UpdatedBy:   userName,
		UpdatedAt:   now,
		Contents:    contents,
		Remarks:     *r.Remarks,
		Tags:        *r.Tags,
		DeletedAt:   time.Time{},
	}
}

func PutToPublishContentsTemplates(r *PublishContentsTemplatesBody, ID, tenantID, userName string) (model.PublishContentsTemplates, error) {
	now := time.Now()

	contents := make(map[string]any)
	for key, value := range r.Contents {
		contents[key] = value
	}

	publishContentsTemplates := model.PublishContentsTemplates{
		ID:          primitive.ObjectID{},
		TenantID:    tenantID,
		DisplayName: r.DisplayName,
		Channel:     r.Channel.Channel,
		CreatedBy:   "",
		CreatedAt:   time.Time{},
		UpdatedBy:   userName,
		UpdatedAt:   now,
		Contents:    contents,
		Remarks:     *r.Remarks,
		Tags:        *r.Tags,
		DeletedAt:   time.Time{},
	}

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return model.PublishContentsTemplates{}, err
	}

	publishContentsTemplates.ID = objectID

	return publishContentsTemplates, nil
}

func PostToPublishContentsTemplates_want(r *PublishContentsTemplatesBody, tenantID, userName string) model.PublishContentsTemplates {
	template := model.PublishContentsTemplates{
		TenantID:    tenantID,
		Channel:     string(r.Channel),
		Contents:    r.Contents,
		CreatedBy:   userName,
		DisplayName: r.DisplayName,
		Remarks:     *r.Remarks,
		Tags:        *r.Tags,
		UpdatedBy:   userName,
	}
	return template
}

func PutToPublishContentsTemplates_want(r *PublishContentsTemplatesBody, ID, tenantID, userName string) (model.PublishContentsTemplates, error) {

	objid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return model.PublishContentsTemplates{}, xerrors.Errorf("convert object_id error: %w", err)
	}

	template := model.PublishContentsTemplates{
		ID:          objid,
		TenantID:    tenantID,
		Channel:     string(r.Channel),
		Contents:    r.Contents,
		DisplayName: r.DisplayName,
		Remarks:     *r.Remarks,
		Tags:        *r.Tags,
		UpdatedBy:   userName,
	}
	return template, nil
}
