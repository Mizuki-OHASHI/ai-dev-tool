# 目標

DB 定義 と AP 定義 を入力として `publish_contents_templates` ブランチで追加された内容を再現すること。

## 入力

### DB 定義

- `model/publish_contents_templates.go`

  ```go
  package model

  import (
  	"time"

  	"go.mongodb.org/mongo-driver/bson/primitive"
  )

  type PublishContentsTemplates struct {
  	ID          primitive.ObjectID `bson:"_id"`
  	TenantID    string             `bson:"tenant_id"`
  	DisplayName string             `bson:"display_name"`
  	Channel     string             `bson:"channel"`
  	CreatedBy   string             `bson:"created_by"`
  	CreatedAt   time.Time          `bson:"created_at"`
  	UpdatedBy   string             `bson:"updated_by"`
  	UpdatedAt   time.Time          `bson:"updated_at"`
  	Contents    map[string]any     `bson:"contents"`
  	Remarks     string             `bson:"remarks"`
  	Tags        []string           `bson:"tags"`
  	DeletedAt   time.Time          `bson:"deleted_at"`
  }

  type PublishContentsTemplatesList = []PublishContentsTemplates
  ```

### API 定義

- `reference/publish_contents_templates.yaml`
  ```yaml
  openapi: 3.0.0
  x-stoplight:
    id: ww1q1v92674m2
  info:
    title: publish_contents_templates
    version: "1.0"
    description: "https://www.notion.so/supsys/c7d413c43f374fb28d4e75bd1104b4af"
  paths:
    /craft/v1/publish_contents_templates:
      get:
        summary: コンテンツテンプレート取得
        tags: []
        responses:
          "200":
            $ref: "#/components/responses/ListPublishContentsTemplatesResponse"
          "400":
            description: Bad Request
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "500":
            description: Internal Server Error
            headers: {}
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/ErrorResponse
        operationId: get-craft-v1-publish_contents_templates
        parameters:
          - $ref: ./common.yaml#/components/parameters/BrowserKeywords
          - $ref: ./common.yaml#/components/parameters/BrowserRows
          - $ref: ./common.yaml#/components/parameters/BrowserStartRow
          - $ref: ./common.yaml#/components/parameters/BrowserOrderBy
          - $ref: ./common.yaml#/components/parameters/BrowserFilter
        description: ""
      parameters: []
      post:
        summary: コンテンツテンプレート新規作成
        operationId: post-craft-v1-publish_contents_templates
        responses:
          "201":
            $ref: "#/components/responses/PublishContentsTemplatesIDResponse"
          "400":
            description: Bad Request
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "500":
            description: Internal Server Error
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/ErrorResponse
        parameters: []
        requestBody:
          $ref: "#/components/requestBodies/PublishContentsTemplatesRequest"
        description: ""
    "/craft/v1/publish_contents_templates/{id}":
      parameters:
        - schema:
            type: string
          name: id
          in: path
          required: true
      get:
        summary: コンテンツテンプレートの詳細取得
        operationId: get-craft-v1-publish_contents_templates-id
        responses:
          "200":
            $ref: "#/components/responses/GetPublishContentsTemplatesIDResponse"
          "400":
            description: Bad Request
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "404":
            description: Not Found
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "500":
            description: Internal Server Error
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/ErrorResponse
        parameters: []
      put:
        summary: コンテンツテンプレートの変更
        operationId: put-craft-v1-publish_contents_templates-id
        responses:
          "200":
            $ref: "#/components/responses/PublishContentsTemplatesIDResponse"
          "400":
            description: Bad Request
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "404":
            description: Not Found
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "500":
            description: Internal Server Error
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/ErrorResponse
        parameters: []
        requestBody:
          $ref: "#/components/requestBodies/PublishContentsTemplatesRequest"
      delete:
        summary: コンテンツテンプレートの削除
        operationId: delete-craft-v1-publish_contents_templates-id
        responses:
          "204":
            description: No Content
          "400":
            description: Bad Request
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "404":
            description: Not Found
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/BadRequestResponse
          "500":
            description: Internal Server Error
            content:
              application/json:
                schema:
                  $ref: ./common.yaml#/components/schemas/ErrorResponse
        parameters: []
  components:
    schemas:
      PublishContentsTemplatesInfo:
        title: PublishContentsTemplatesInfo
        type: object
        x-examples:
          example-1:
            id: "1001"
            display_name: コンテンツテンプレートのテスト
            channel: line
            created_by: user_1
            created_at: "2019-08-24T14:15:22Z"
            updated_by: user_2
            updated_at: "2019-08-24T14:15:22Z"
            remarks: 備考
            tag:
              - tag1
              - tag2
        description: List用
        properties:
          id:
            type: string
            description: ユニークID
          display_name:
            type: string
            description: 表示名
          channel:
            type: string
            description: チャンネル名
            enum:
              - fcm_push
              - line
              - email
              - sms
          updated_by:
            type: string
            description: 最終更新者
          updated_at:
            type: string
            format: date-time
            description: 最終更新日時
          tags:
            description: タグ
            type: array
            items:
              type: string
        required:
          - id
          - display_name
          - channel
      PublishContentsTemplates:
        title: PublishContentsTemplates
        type: object
        x-examples:
          example-1:
            id: "1001"
            display_name: コンテンツテンプレートのテスト
            channel: line
            created_by: user_1
            created_at: "2019-08-24T14:15:22Z"
            updated_by: user_2
            updated_at: "2019-08-24T14:15:22Z"
            contents:
              messages:
                - type: text
                  text: "Hello, world1"
                - type: text
                  text: "Hello, world2"
            remarks: 備考
            tag:
              - tag1
              - tag2
        description: Get用
        properties:
          id:
            type: string
            description: ユニークID
          display_name:
            type: string
            description: 表示名
          channel:
            type: string
            description: チャンネル名
            enum:
              - fcm_push
              - line
              - email
              - sms
          created_by:
            type: string
            description: 作成者
          created_at:
            type: string
            format: date-time
            description: 作成日時
          updated_by:
            type: string
            description: 最終更新者
          updated_at:
            type: string
            format: date-time
            description: 最終更新日時
          contents:
            description: コンテンツ
            type: object
          remarks:
            type: string
            description: 備考
          tags:
            description: タグ
            type: array
            items:
              type: string
        required:
          - id
          - display_name
          - channel
          - contents
      PublishContentsTemplatesBody:
        title: PublishContentsTemplatesBody
        type: object
        x-examples:
          example-1:
            display_name: コンテンツテンプレートのテスト
            channel: line
            contents:
              messages:
                - type: text
                  text: "Hello, world1"
                - type: text
                  text: "Hello, world2"
            remarks: 備考
            tag:
              - tag1
              - tag2
        description: Postのリクエスト用
        properties:
          display_name:
            type: string
            description: 表示名
          channel:
            type: string
            description: チャンネル名
            enum:
              - fcm_push
              - line
              - email
              - sms
          contents:
            description: コンテンツ
            type: object
          remarks:
            type: string
            description: 備考
          tags:
            description: タグ
            type: array
            items:
              type: string
        required:
          - display_name
          - channel
          - contents
    responses:
      ListPublishContentsTemplatesResponse:
        description: ""
        content:
          application/json:
            schema:
              type: object
              properties:
                meta:
                  $ref: ./common.yaml#/components/schemas/SearchMetadata
                data:
                  type: array
                  items:
                    $ref: "#/components/schemas/PublishContentsTemplatesInfo"
      PublishContentsTemplatesIDResponse:
        description: ""
        content:
          application/json:
            schema:
              type: object
              properties:
                id:
                  type: string
              required:
                - id
      GetPublishContentsTemplatesIDResponse:
        description: ""
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/PublishContentsTemplates"
    requestBodies:
      PublishContentsTemplatesRequest:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/PublishContentsTemplatesBody"
        description: ""
    parameters: {}
  ```

## 出力

### DAO

- `dao/publish_contents_templates.go`

  ```go
  //go:generate mockgen -source publish_contents_templates.go -destination mock/publish_contents_templates.go -package mock
  package dao

  import (
  	"context"
  	"fmt"
  	"time"

  	"github.com/supsysjp/craft/dao/listutil"
  	"github.com/supsysjp/craft/infra"
  	"github.com/supsysjp/craft/model"
  	"github.com/supsysjp/craft/util"
  	"go.mongodb.org/mongo-driver/bson"
  	"go.mongodb.org/mongo-driver/bson/primitive"
  	"go.mongodb.org/mongo-driver/mongo"
  	"golang.org/x/xerrors"
  )

  type PublishContentsTemplatesDao interface {
  	List(ctx context.Context, req model.SearchQueryParameters) (model.PublishContentsTemplatesList, int64, error)
  	Fetch(ctx context.Context, tenantID, id string) (*model.PublishContentsTemplates, error)
  	Create(ctx context.Context, templates model.PublishContentsTemplates) (string, error)
  	Update(ctx context.Context, templates model.PublishContentsTemplates) (string, error)
  	Delete(ctx context.Context, tenantID, id string) error
  }

  type publishContentsTemplatesDao struct {
  	db        infra.MongoClientWrapper
  	tableName string
  }

  type publishContentsTemplatesMongoDBNoneID struct {
  	TenantID    string         `bson:"tenant_id"`
  	DisplayName string         `bson:"display_name"`
  	Channel     string         `bson:"channel"`
  	CreatedBy   string         `bson:"created_by"`
  	CreatedAt   time.Time      `bson:"created_at"`
  	UpdatedBy   string         `bson:"updated_by"`
  	UpdatedAt   time.Time      `bson:"updated_at"`
  	Contents    map[string]any `bson:"contents"`
  	Remarks     string         `bson:"remarks"`
  	Tags        []string       `bson:"tags"`
  	DeletedAt   time.Time      `bson:"deleted_at"`
  }

  func NewPublishContentsTemplatesDao(mongo infra.MongoClientWrapper) PublishContentsTemplatesDao {
  	return &publishContentsTemplatesDao{
  		db:        mongo,
  		tableName: "publish_contents_templates",
  	}
  }

  // deleted_atがあるものをどうするか
  func (d *publishContentsTemplatesDao) List(ctx context.Context, params model.SearchQueryParameters) (model.PublishContentsTemplatesList, int64, error) {
  	findOption, err := listutil.FindOption(params.Orderby, params.StartRow, params.Rows)
  	if err != nil {
  		return nil, 0, err
  	}

  	fmt.Println(time.Time{}.String())

  	filter, err := listutil.FilterOptions(params.Filter)
  	if err != nil {
  		return nil, 0, err
  	}

  	//FilterOptionsにはstring以外を入れることができない
  	options := (*filter)["$and"].([]bson.M)
  	options = append(options, bson.M{"deleted_at": bson.M{"$eq": time.Time{}}})
  	filterAddDeletedAt := bson.M{"$and": options}

  	findOption.SetProjection(bson.D{
  		{
  			Key:   "_id",
  			Value: 1,
  		},
  		{
  			Key:   "display_name",
  			Value: 1,
  		},
  		{
  			Key:   "channel",
  			Value: 1,
  		},
  		{
  			Key:   "created_by",
  			Value: 1,
  		},
  		{
  			Key:   "created_at",
  			Value: 1,
  		},
  		{
  			Key:   "updated_by",
  			Value: 1,
  		},
  		{
  			Key:   "updated_at",
  			Value: 1,
  		},
  		{
  			Key:   "tags",
  			Value: 1,
  		},
  		{
  			Key:   "remarks",
  			Value: 1,
  		},
  		{
  			Key:   "deleted_at",
  			Value: 1,
  		},
  	})

  	colls := d.db.Database(ctx).Collection(d.tableName)
  	cursor, err := colls.Find(ctx, filterAddDeletedAt, findOption)
  	if err != nil {
  		return nil, 0, xerrors.Errorf("list publish contents templates error: %w", err)
  	}
  	defer cursor.Close(ctx)

  	var results model.PublishContentsTemplatesList
  	if err := cursor.All(ctx, &results); err != nil {
  		return nil, 0, xerrors.Errorf("cursor publish contents templates error: %w", err)
  	}

  	size, err := colls.CountDocuments(ctx, filterAddDeletedAt)
  	if err != nil {
  		return nil, 0, xerrors.Errorf("count publish contents templates error: %w", err)
  	}

  	return results, size, nil
  }

  func (d *publishContentsTemplatesDao) Fetch(ctx context.Context, tenantID, id string) (*model.PublishContentsTemplates, error) {

  	objid, err := primitive.ObjectIDFromHex(id)
  	if err != nil {
  		return nil, xerrors.Errorf("convert object_id error: %w", err)
  	}

  	filter := bson.M{
  		"_id":       objid,
  		"tenant_id": tenantID,
  	}

  	colls := d.db.Database(ctx).Collection(d.tableName)
  	single := colls.FindOne(ctx, filter)
  	if single.Err() == mongo.ErrNoDocuments {
  		return nil, xerrors.Errorf("Not Found tenant_id: %s, id: %s", tenantID, id)
  	} else if single.Err() != nil {
  		return nil, xerrors.Errorf("fetch publish contents templates error: %w", err)
  	}

  	var result *model.PublishContentsTemplates
  	err = single.Decode(&result)
  	if err != nil {
  		return nil, xerrors.Errorf("publish contents templates decode error: %w", err)
  	}

  	if result.DeletedAt.IsZero() {
  		return result, xerrors.Errorf("fetch publish contents templates error: %w", single.Err())
  	}
  	return nil, nil
  }

  func (d *publishContentsTemplatesDao) Create(ctx context.Context, templates model.PublishContentsTemplates) (string, error) {
  	colls := d.db.Database(ctx).Collection(d.tableName)

  	currentTime := util.Now(ctx)

  	insertData := publishContentsTemplatesMongoDBNoneID{
  		TenantID:    templates.TenantID,
  		DisplayName: templates.DisplayName,
  		Channel:     templates.Channel,
  		CreatedBy:   templates.CreatedBy,
  		CreatedAt:   currentTime,
  		UpdatedBy:   templates.UpdatedBy,
  		UpdatedAt:   currentTime,
  		Contents:    templates.Contents,
  		Remarks:     templates.Remarks,
  		Tags:        templates.Tags,
  		DeletedAt:   time.Time{},
  	}

  	insertResult, err := colls.InsertOne(ctx, insertData)
  	if err != nil {
  		return "", xerrors.Errorf("insert publish contents templates error: %w", err)
  	}

  	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
  	if !ok {
  		return "", xerrors.Errorf("insert publish contents templates, but not found id: %v", insertResult.InsertedID)
  	}

  	return oid.Hex(), nil
  }

  func (d *publishContentsTemplatesDao) Update(ctx context.Context, templates model.PublishContentsTemplates) (string, error) {
  	filter := bson.M{
  		"_id":       templates.ID,
  		"tenant_id": templates.TenantID,
  	}

  	colls := d.db.Database(ctx).Collection(d.tableName)
  	single := colls.FindOne(ctx, filter)
  	if single.Err() == mongo.ErrNoDocuments {
  		return "", xerrors.Errorf("Not Found tenant_id: %s, id: %s", templates.TenantID, templates.ID)
  	} else if single.Err() != nil {
  		return "", xerrors.Errorf("fetch publish contents templates error: %w", single.Err())
  	}

  	var result *model.PublishContentsTemplates
  	err := single.Decode(&result)
  	if err != nil {
  		return "", xerrors.Errorf("publish contents templates decode error: %w", err)
  	}

  	if !result.DeletedAt.IsZero() {
  		return "", xerrors.Errorf("fetch publish contents templates error: %w", single.Err())
  	}

  	currentTime := util.Now(ctx)
  	update := bson.M{
  		"$set": bson.M{
  			"display_name": templates.DisplayName,
  			"channel":      templates.Channel,
  			"updated_by":   templates.UpdatedBy,
  			"updated_at":   currentTime,
  			"contents":     templates.Contents,
  			"remarks":      templates.Remarks,
  			"tags":         templates.Tags,
  		},
  	}
  	_, err = colls.UpdateOne(ctx, filter, update)
  	if err != nil {
  		return "", xerrors.Errorf("update publish contents templates error: %w", err)
  	}
  	return templates.ID.Hex(), nil
  }

  func (d *publishContentsTemplatesDao) Delete(ctx context.Context, tenantID, id string) error {

  	objid, err := primitive.ObjectIDFromHex(id)
  	if err != nil {
  		return xerrors.Errorf("convert object_id error: %w", err)
  	}

  	filter := bson.M{
  		"_id":       objid,
  		"tenant_id": tenantID,
  	}

  	colls := d.db.Database(ctx).Collection(d.tableName)
  	single := colls.FindOne(ctx, filter)
  	if single.Err() == mongo.ErrNoDocuments {
  		return xerrors.Errorf("Not Found tenant_id: %s, id: %s", tenantID, id)
  	} else if single.Err() != nil {
  		return xerrors.Errorf("fetch publish contents templates error: %w", single.Err())
  	}

  	var result *model.PublishContentsTemplates
  	err = single.Decode(&result)
  	if err != nil {
  		return xerrors.Errorf("publish contents templates decode error: %w", err)
  	}

  	if !result.DeletedAt.IsZero() {
  		return xerrors.Errorf("Not Found tenant_id: %s, id: %s", tenantID, id)
  	}

  	currentTime := util.Now(ctx)
  	update := bson.M{
  		"$set": bson.M{
  			"deleted_at": currentTime,
  		},
  	}
  	_, err = colls.UpdateOne(ctx, filter, update)
  	if err != nil {
  		return xerrors.Errorf("delete publish contents templates error: %w", err)
  	}

  	return nil
  }
  ```

- `dao/mock/publish_contents_templates.go`
  ```bash
  //go:generate mockgen -source publish_contents_templates.go -destination mock/publish_contents_templates.go -package mock
  ```

### Controller

- `controller/pulishcontentstemplates/controller.go`

  ```go
  //go:generate oapi-codegen -generate types,gin,spec,skip-prune -templates ./../../settings/oapi-codegen -import-mapping ./common.yaml:github.com/supsysjp/craft/controller/common -package publishcontentstemplates -o schema.gen.go ./../../reference/publish_contents_templates.yaml
  package publishcontentstemplates

  import (
  	"fmt"
  	"net/http"

  	gin "github.com/gin-gonic/gin"
  	"github.com/supsysjp/craft/controller/common"
  	"github.com/supsysjp/craft/logger"
  	"github.com/supsysjp/craft/model"
  	usecase "github.com/supsysjp/craft/usecase"
  )

  type publishcontentstemplates struct {
  	u usecase.PublishContentsTemplatesUsecase
  	ServerInterface
  }

  func NewPublishcontentstemplates(u usecase.PublishContentsTemplatesUsecase) ServerInterface {
  	return &publishcontentstemplates{u: u}
  }

  func (c *publishcontentstemplates) DeleteCraftV1PublishContentsTemplatesId(ctx *gin.Context, s1 string) {
  	au, err := common.GetAuthenticatedUser(ctx)
  	if err != nil {
  		detail := err.Error()
  		logger.Error(err)
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Get Authenticated User failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	err = c.u.Delete(ctx, au.TenantId, s1)
  	if err != nil {
  		logger.Error(err)
  		if err == usecase.ErrNotFoundPublishContentsTemplatesID {
  			ctx.JSON(
  				http.StatusNotFound,
  				common.NewBadRequest(
  					fmt.Sprintf("Not Found Publish Contents Templates ID Error: %s", s1),
  				),
  			)
  			return
  		}
  		errorMsg := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
  			Code:   "internal_server_error",
  			Msg:    "Internal Server Error",
  			Detail: &errorMsg,
  		})
  		return
  	}
  	ctx.Writer.WriteHeader(http.StatusNoContent)
  }

  func (c *publishcontentstemplates) GetCraftV1PublishContentsTemplates(ctx *gin.Context, s1 GetCraftV1PublishContentsTemplatesParams) {
  	req := model.SearchQueryRequestParameters{}
  	err := ctx.BindQuery(&req)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, common.NewBadRequest(err.Error()))
  		return
  	}

  	sqm, err := common.ToSearchQueryModel(req)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, common.NewBadRequest(err.Error()))
  		return
  	}

  	au, err := common.GetAuthenticatedUser(ctx)
  	if err != nil {
  		detail := err.Error()
  		logger.Error(err)
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Get Authenticated User failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}
  	l, total, canDownload, err := c.u.List(ctx, au.TenantId, *sqm)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, err.Error())
  		return
  	}

  	meta, err := common.ToSearchMetaData(req, total, canDownload)
  	if err != nil {
  		logger.Error(err)
  		detail := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Convert search meta data failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	res, err := ListFromPublishContentsTemplatesList(&l)
  	if err != nil {
  		logger.Error(err)
  		detail := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Convert scenario list failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	ctx.JSON(http.StatusOK, ListPublishContentsTemplatesResponse{
  		Meta: meta,
  		Data: &res,
  	})
  }

  func (c *publishcontentstemplates) GetCraftV1PublishContentsTemplatesId(ctx *gin.Context, s1 string) {
  	au, err := common.GetAuthenticatedUser(ctx)
  	if err != nil {
  		detail := err.Error()
  		logger.Error(err)
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Get Authenticated User failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	res, err := c.u.Fetch(ctx, au.TenantId, s1)
  	if err != nil {
  		logger.Error(err)
  		if err == usecase.ErrNotFoundPublishContentsTemplatesID {
  			ctx.JSON(
  				http.StatusNotFound,
  				common.NewBadRequest(
  					fmt.Sprintf("Not Found Publish Contents Templates ID Error: %s", s1),
  				),
  			)
  			return
  		}
  		errorMsg := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
  			Code:   "internal_server_error",
  			Msg:    "Internal Server Error",
  			Detail: &errorMsg,
  		})
  		return
  	}
  	resParam, err := GetFromPublishContentsTemplates(*res)
  	if err != nil {
  		logger.Error(err)
  		errorMsg := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
  			Code:   "internal_server_error",
  			Msg:    "Internal Server Error",
  			Detail: &errorMsg,
  		})
  		return
  	}
  	ctx.JSON(http.StatusOK, resParam)
  }

  func (c *publishcontentstemplates) PostCraftV1PublishContentsTemplates(ctx *gin.Context) {
  	var b *PublishContentsTemplatesBody
  	err := ctx.BindJSON(&b)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, common.NewBadRequest(err.Error()))
  		return
  	}

  	au, err := common.GetAuthenticatedUser(ctx)
  	if err != nil {
  		detail := err.Error()
  		logger.Error(err)
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Get Authenticated User failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	m := PostToPublishContentsTemplates(b, au.TenantId, *au.DisplayName)
  	id, err := c.u.Create(ctx, m)
  	if err != nil {
  		logger.Error(err)
  		errorMsg := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
  			Code:   "internal_server_error",
  			Msg:    "Internal Server Error",
  			Detail: &errorMsg,
  		})
  		return
  	}
  	res := PublishContentsTemplatesIDResponse{
  		Id: id,
  	}
  	ctx.JSON(http.StatusCreated, res)
  }

  func (c *publishcontentstemplates) PutCraftV1PublishContentsTemplatesId(ctx *gin.Context, s1 string) {
  	var b *PublishContentsTemplatesBody
  	err := ctx.BindJSON(&b)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, common.NewBadRequest(err.Error()))
  		return
  	}

  	au, err := common.GetAuthenticatedUser(ctx)
  	if err != nil {
  		detail := err.Error()
  		logger.Error(err)
  		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
  			Msg:    "Get Authenticated User failed",
  			Code:   "internal_server_error",
  			Detail: &detail,
  		})
  		return
  	}

  	m, err := PutToPublishContentsTemplates(b, s1, au.TenantId, *au.DisplayName)
  	if err != nil {
  		logger.Error(err)
  		ctx.JSON(http.StatusBadRequest, common.NewBadRequest(err.Error()))
  		return
  	}

  	id, err := c.u.Update(ctx, m)
  	if err != nil {
  		logger.Error(err)
  		errorMsg := err.Error()
  		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
  			Code:   "internal_server_error",
  			Msg:    "Internal Server Error",
  			Detail: &errorMsg,
  		})
  		return
  	}
  	res := PublishContentsTemplatesIDResponse{
  		Id: id,
  	}
  	ctx.JSON(http.StatusCreated, res)
  }
  ```

- `controller/pulishcontentstemplates/converter.go`

  ```go
  package publishcontentstemplates

  import (
  	"fmt"

  	"github.com/supsysjp/craft/controller/common"
  	"github.com/supsysjp/craft/model"
  	"go.mongodb.org/mongo-driver/bson/primitive"
  	"golang.org/x/xerrors"
  )

  func ListFromPublishContentsTemplatesList(s *model.PublishContentsTemplatesList) ([]PublishContentsTemplatesInfo, error) {
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

  func GetFromPublishContentsTemplates(m model.PublishContentsTemplates) (*PublishContentsTemplates, error) {

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

  func PostToPublishContentsTemplates(r *PublishContentsTemplatesBody, tenantID, userName string) model.PublishContentsTemplates {
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

  func PutToPublishContentsTemplates(r *PublishContentsTemplatesBody, ID, tenantID, userName string) (model.PublishContentsTemplates, error) {

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
  ```

- `controller/pulishcontentstemplates/schema.gen.go`
  ```bash
  go:generate oapi-codegen -generate types,gin,spec,skip-prune -templates ./../../settings/oapi-codegen -import-mapping ./common.yaml:github.com/supsysjp/craft/controller/common -package publishcontentstemplates -o schema.gen.go ./../../reference/publish_contents_templates.yaml
  ```

### Main

- `main.go`
  PENDING...
