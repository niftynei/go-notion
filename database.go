package notion

import "time"

// Database objects describe the property schema of a database in Notion.
// Page are the items (or children) in a database.
// Page property values must conform to the property objects laid out in the parent database object.
type Database struct {
	Object         ObjectType          `json:"object,omitempty"`
	ID             string              `json:"id,omitempty"`
	CreatedTime    time.Time           `json:"created_time,omitempty"`
	LastEditedTime time.Time           `json:"last_edited_time,omitempty"`
	Title          []*RichText         `json:"title,omitempty"`
	Properties     map[string]Property `json:"properties,omitempty"`
}

// PropertyType is type of database Property.
type PropertyType string

// PropertyType enums.
const (
	PropertyTitle          PropertyType = "title"
	PropertyRichText       PropertyType = "rich_text"
	PropertyNumber         PropertyType = "number"
	PropertySelect         PropertyType = "select"
	PropertyMultiSelect    PropertyType = "multi_select"
	PropertyDate           PropertyType = "date"
	PropertyPeople         PropertyType = "people"
	PropertyFile           PropertyType = "file"
	PropertyCheckbox       PropertyType = "checkbox"
	PropertyURL            PropertyType = "url"
	PropertyEmail          PropertyType = "email"
	PropertyPhoneNumber    PropertyType = "phone_number"
	PropertyFormula        PropertyType = "formula"
	PropertyRelation       PropertyType = "relation"
	PropertyRollup         PropertyType = "rollup"
	PropertyCreatedTime    PropertyType = "created_time"
	PropertyCreatedBy      PropertyType = "created_by"
	PropertyLastEditedTime PropertyType = "last_edited_time"
	PropertyLastEditedBy   PropertyType = "last_edited_by"
	PropertyUniqueID       PropertyType = "unique_id"
)

// NumberFormat is format of number property value.
type NumberFormat string

// NumberFormat enums.
const (
	NumberFormatNumber           NumberFormat = "number"
	NumberFormatNumberWithCommas NumberFormat = "number_with_commas"
	NumberFormatPercent          NumberFormat = "percent"
	NumberFormatDollar           NumberFormat = "dollar"
	NumberFormatEuro             NumberFormat = "euro"
	NumberFormatPound            NumberFormat = "pound"
	NumberFormatYen              NumberFormat = "yen"
	NumberFormatRuble            NumberFormat = "ruble"
	NumberFormatRupee            NumberFormat = "rupee"
	NumberFormatWon              NumberFormat = "won"
	NumberFormatYuan             NumberFormat = "yuan"
)

// Property is mix type of database property.
type Property struct {
	ID       string       `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Title    *struct{}    `json:"title,omitempty"`
	RichText *struct{}    `json:"rich_text,omitempty"`
	Number   *struct {
		Format NumberFormat `json:"format,omitempty"`
	} `json:"number,omitempty"`
	Select *struct {
		Options []*SelectOption `json:"options,omitempty"`
	} `json:"select,omitempty"`
	MultiSelect *struct {
		Options []*SelectOption `json:"options,omitempty"`
	} `json:"multi_select,omitempty"`
	Checkbox    *struct{} `json:"checkbox,omitempty"`
	URL         *struct{} `json:"url,omitempty"`
	Email       *struct{} `json:"email,omitempty"`
	PhoneNumber *struct{} `json:"phone_number,omitempty"`
	Formula     *struct {
		Expression string `json:"expression,omitempty"`
	} `json:"formula,omitempty"`
	Relation *struct {
		DatabaseID         string `json:"database_id,omitempty"`
		SyncedPropertyName string `json:"synced_property_name,omitempty"`
		SyncedPropertyID   string `json:"synced_property_id,omitempty"`
	} `json:"relation,omitempty"`
	Rollup *struct {
		RelationPropertyName string `json:"relation_property_name,omitempty"`
		RelationPropertyID   string `json:"relation_property_id,omitempty"`
		RollupPropertyName   string `json:"rollup_property_name,omitempty"`
		RollupPropertyID     string `json:"rollup_property_id,omitempty"`
		Function             string `json:"function,omitempty"`
	} `json:"rollup,omitempty"`
	People         *struct{} `json:"people,omitempty"`
	Date           *struct{} `json:"date,omitempty"`
	File           *struct{} `json:"files,omitempty"`
	CreatedTime    *struct{} `json:"created_time,omitempty"`
	CreatedBy      *struct{} `json:"created_by,omitempty"`
	LastEditedTime *struct{} `json:"last_edited_time,omitempty"`
	LastEditedBy   *struct{} `json:"last_edited_by,omitempty"`
}

// SelectOption is the option value of single selector and multi selector.
type SelectOption struct {
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Color Color  `json:"color,omitempty"`
}
