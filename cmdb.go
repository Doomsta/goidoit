package goidoit

import (
	"context"
)

type CmdbService service

type ReadObjectTypeCategoriesCatgResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Const       string `json:"const"`
	MultiValue  string `json:"multi_value"`
	SourceTable string `json:"source_table"`
	Parent      string `json:"parent,omitempty"`
}

type ReadObjectTypeCategoriesResponse struct {
	Catg []ReadObjectTypeCategoriesCatgResponse `json:"catg"`
}

func (c *CmdbService) ReadObjectTypeCategories(ctx context.Context, objType string) (*ReadObjectTypeCategoriesResponse, error) {
	return parse[*ReadObjectTypeCategoriesResponse](c.client.Request(ctx, "cmdb.object_type_categories.read", struct {
		Type string `json:"type"`
	}{objType}))
}

type GetDialogResponse struct {
	Id    string `json:"id"`
	Const string `json:"const"`
	Title string `json:"title"`
}

func (c *CmdbService) GetDialog(ctx context.Context, category, property string) ([]GetDialogResponse, error) {
	return parse[[]GetDialogResponse](c.client.Request(ctx, "cmdb.dialog.read", struct {
		Category string `json:"category"`
		Property string `json:"property"`
	}{category, property}))
}

func (c *CmdbService) GetReport(ctx context.Context, ID int) ([]map[string]string, error) {
	return parse[[]map[string]string](c.client.Request(ctx, "cmdb.reports.read", struct {
		ID int `json:"id"`
	}{ID}))
}

type GetObjectResponse struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Sysid           string `json:"sysid"`
	Type            int    `json:"type"`
	Created         string `json:"created"`
	TypeTitle       string `json:"type_title"`
	TypeGroupTitle  string `json:"type_group_title"`
	Status          int    `json:"status"`
	CmdbStatus      int    `json:"cmdb_status"`
	CmdbStatusTitle string `json:"cmdb_status_title"`
	Image           string `json:"image"`
}

type GetObjectsFilter struct {
	IDs   []int  `json:"ids,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

func (c *CmdbService) GetObjects(ctx context.Context, filter GetObjectsFilter) ([]GetObjectResponse, error) {
	return parse[[]GetObjectResponse](c.client.Request(ctx, "cmdb.objects.read", struct {
		Filter any `json:"filter"`
	}{Filter: filter}))
}

type StatusReadResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Constant string `json:"constant"`
	Color    string `json:"color"`
	Editable bool   `json:"editable"`
}

func (c *CmdbService) StatusRead(ctx context.Context) ([]StatusReadResponse, error) {
	return parse[[]StatusReadResponse](c.client.Request(ctx, "cmdb.status.read", nil))
}
