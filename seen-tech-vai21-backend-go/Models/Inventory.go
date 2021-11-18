package Models

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID            primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	InventoryName string                 `json:"inventoryname,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Address       string                 `json:"address,omitempty"`
	Longitude     float64                `json:"longitude,omitempty"`
	Latitude      float64                `json:"latitude,omitempty"`
	Status        bool                   `json:"status,omitempty"`
	Contents      []InventoryContentsRow `json:"contents,omitempty" bson:"contents"`
}

type InventoryContentsRow struct {
	ProductRef     primitive.ObjectID
	RawMaterialRef primitive.ObjectID
	Amount         float64
}

type InventoryContentsRowPopulated struct {
	ProductRef     Product  `json:"productref,omitempty"`
	RawMaterialRef Material `json:"rawmaterialref,omitempty"`
	Amount         float64  `json:"amount,omitempty"`
}

func (obj *InventoryContentsRowPopulated) CloneFrom(other InventoryContentsRow) {
	obj.ProductRef = Product{}
	obj.RawMaterialRef = Material{}
	obj.Amount = other.Amount
}

func (obj Inventory) Validate() error {
	return validation.ValidateStruct(&obj,
		validation.Field(&obj.InventoryName, validation.Required),
	)
}

func (obj Inventory) GetBSONModificationObj() bson.M {
	self := bson.M{
		"inventoryname": obj.InventoryName,
		"type":          obj.Type,
		"address":       obj.Address,
		"longitude":     obj.Longitude,
		"latitude":      obj.Latitude,
		"status":        obj.Status,
	}
	return self
}

type InventorySearch struct {
	InventoryName       string `json:"inventoryname,omitempty"`
	InventoryNameIsUsed bool   `json:"inventorynameisused,omitempty"`
	Type                string `json:"type,omitempty"`
	TypeIsUsed          bool   `json:"typeisused,omitempty"`
	Status              bool   `json:"status,omitempty"`
	StatusIsUsed        bool   `json:"statusisused,omitempty"`
}

func (obj InventorySearch) GetBSONSearchObj() bson.M {
	self := bson.M{}
	if obj.InventoryNameIsUsed {
		regexPattern := fmt.Sprintf(".*%s.*", obj.InventoryName)
		self["inventoryname"] = bson.D{{"$regex", primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}
	if obj.TypeIsUsed {
		self["type"] = obj.Type
	}
	if obj.StatusIsUsed {
		self["status"] = obj.Status
	}
	return self
}
