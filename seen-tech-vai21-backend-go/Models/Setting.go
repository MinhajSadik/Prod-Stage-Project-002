package Models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductSerial int                `json:"productserial,omitempty"`
}

type SettingSearch struct {
	IDIsUsed            bool               `json:"idisused,omitempty" bson:"idisused,omitempty"`
	ID                  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductSerial       int                `json:"productserial,omitempty"`
	ProductSerialIsUsed bool               `json:"productserialisused,omitempty"`
}

func (obj Setting) GetIdString() string {
	return obj.ID.String()
}

func (obj Setting) GetId() primitive.ObjectID {
	return obj.ID
}

func (obj SettingSearch) GetSettingSearchBSONObj() bson.M {
	self := bson.M{}
	if obj.IDIsUsed {
		self["_id"] = obj.ID
	}

	if obj.ProductSerialIsUsed {
		self["productserial"] = obj.ProductSerial
	}

	return self
}
