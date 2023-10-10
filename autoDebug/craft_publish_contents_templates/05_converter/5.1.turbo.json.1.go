package controller

import (
	"time"

	"github.com/your-package/model"
)

func ListFromPublishContentsTemplatesList(s *model.PublishContentsTemplatesList) ([]PublishContentsTemplatesInfo, error) {
	var result []PublishContentsTemplatesInfo

	for _, item := range *s {
		publishContentsTemplatesInfo := PublishContentsTemplatesInfo{
			Channel:     PublishContentsTemplatesInfoChannel(item.Channel),
			DisplayName: item.DisplayName,
			Id:          item.ID.Hex(),
			Tags:        &item.Tags,
			UpdatedAt:   item.UpdatedAt,
			UpdatedBy:   &item.UpdatedBy,
		}

		result = append(result, publishContentsTemplatesInfo)
	}

	return result, nil
}

func ListFromPublishContentsTemplatesList_want(s *model.PublishContentsTemplatesList) ([]PublishContentsTemplatesInfo, error) {
	// 通常nilがくることはない
	if s == nil {
		return nil, fmt.Errorf("model is nil")
	}
	list := make([]PublishContentsTemplatesInfo, len(*s))
	for i, v := range *s {
		list[i] = PublishContentsTemplatesInfo{
			Id:          v.ID.Hex(),
			DisplayName: v.DisplayName,
			Tags:        &v.Tags,
			UpdatedAt:   common.AddTimeZone(&v.UpdatedAt),
			UpdatedBy:   &v.UpdatedBy,
		}
		switch v.Channel {
		case "fcm_push":
			list[i].Channel = PublishContentsTemplatesInfoChannel("fcm_push")
		case "line":
			list[i].Channel = PublishContentsTemplatesInfoChannel("line")
		case "email":
			list[i].Channel = PublishContentsTemplatesInfoChannel("email")
		default:
			return nil, xerrors.Errorf("Invalid channel : %s", v.Channel)
		}

	}
	return list, nil
}
