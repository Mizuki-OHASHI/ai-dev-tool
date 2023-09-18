## `controller/converter`

#### 概要

converter.go ないの関数を 1 つないし 2 つずつ生成させる.

| No. | input                        | output                                                        | remarks |
| :-- | :--------------------------- | :------------------------------------------------------------ | :------ |
| 5.1 | model, part of schema.gen.go | ListFromPublishContentsTemplatesList                          |         |
| 5.2 | model, part of schema.gen.go | GetFromPublishContentsTemplates                               |         |
| 5.3 | model, part of schema.gen.go | PostToPublishContentsTemplates, PutToPublishContentsTemplates |         |

#### 所感

- 関数名, 引数の型, 返り値の型だけで概ね大枠は生成される.
- 典型的なエラーハンドリングはされるが, 少し込み入ったものはされない.

  ```go
  // want
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

  // got
  PublishContentsTemplatesInfoChannel(m.Channel)
  ```

- ポインタまわりが少し違う.

  ```go
  // want
  &PublishContentsTemplates{
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

  // got
  &PublishContentsTemplates{
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
  ```
