package controller

import (
	"time"

	"github.com/your-module-path/model"
)

func GetFromPublishContentsTemplates(m model.PublishContentsTemplates) (*PublishContentsTemplates, error) {
	p := &PublishContentsTemplates{
		ID:          m.ID.Hex(),
		TenantID:    m.TenantID,
		DisplayName: m.DisplayName,
		Channel:     PublishContentsTemplatesInfoChannel(m.Channel),
		CreatedBy:   m.CreatedBy,
		CreatedAt:   m.CreatedAt,
		UpdatedBy:   m.UpdatedBy,
		UpdatedAt:   m.UpdatedAt,
		Contents:    m.Contents,
		Remarks:     m.Remarks,
		Tags:        m.Tags,
		DeletedAt:   m.DeletedAt,
	}

	return p, nil
}

func GetFromPublishContentsTemplates_want(m model.PublishContentsTemplates) (*PublishContentsTemplates, error) {

	var channel PublishContentsTemplatesChannel
	switch m.Channel {
	case "fcm_push":
		channel = PublishContentsTemplatesChannel("fcm_push")
	case "line":
		channel = PublishContentsTemplatesChannel("line")
	case "email":
		channel = PublishContentsTemplatesChannel("email")
	default:
		return nil, xerrors.Errorf("Invalid channel : %s", m.Channel)
	}

	template := &PublishContentsTemplates{
		Channel:     channel,
		Contents:    m.Contents,
		CreatedAt:   &m.CreatedAt,
		CreatedBy:   &m.CreatedBy,
		DisplayName: m.DisplayName,
		Id:          m.ID.Hex(),
		Remarks:     &m.Remarks,
		Tags:        &m.Tags,
		UpdatedAt:   &m.UpdatedAt,
		UpdatedBy:   &m.UpdatedBy,
	}
	return template, nil
}
