package Controllers

import (
	"SEEN-TECH-VAI21-BACKEND-GO/DBManager"
	"SEEN-TECH-VAI21-BACKEND-GO/Models"
	"SEEN-TECH-VAI21-BACKEND-GO/Utils"
	"context"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func isInventoryNameExisting(collection *mongo.Collection, InventoryName string) bool {

	var filter bson.M = bson.M{
		"inventoryname": InventoryName,
	}
	var results []bson.M
	b, results := Utils.FindByFilter(collection, filter)
	return (b && len(results) > 0)
}

func InventoryNew(c *fiber.Ctx) error {
	Collection := DBManager.SystemCollections.Inventory

	var self Models.Inventory
	c.BodyParser(&self)

	err := self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}
	if isInventoryNameExisting(Collection, self.InventoryName) {
		c.Status(500)
		return errors.New("Inventory Name is already exist")
	}

	res, err := Collection.InsertOne(context.Background(), self)
	if err != nil {
		c.Status(500)
		return err
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)

	return nil

}

func InventoryGetAll(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Inventory

	results := []bson.M{}

	var searchParams Models.InventorySearch
	c.BodyParser(&searchParams)

	cur, err := collection.Find(context.Background(), searchParams.GetBSONSearchObj())
	if err != nil {
		c.Status(500)
		return err
	}
	defer cur.Close(context.Background())
	cur.All(context.Background(), &results)
	response, _ := json.Marshal(bson.M{
		"results": results,
	})
	c.Set("content-type", "application/json")
	c.Status(200).Send(response)

	return nil
}
func InventorySetStatus(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Inventory
	if c.Params("id") == "" || c.Params("new_status") == "" {
		c.Status(404)
		return errors.New("all params not sent correctly")
	}
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	var newValue = true
	if c.Params("new_status") == "Inactive" {
		newValue = false
	}
	updateData := bson.M{
		"$set": bson.M{
			"status": newValue,
		},
	}
	_, updateErr := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, updateData)
	if updateErr != nil {
		c.Status(500)
		return errors.New("an error occurred when modifing inventory status")
	}
	c.Status(200).Send([]byte("status modified successfully"))
	return nil
}
func InventoryModify(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Inventory
	var self Models.Inventory
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}
	updateQuery := bson.M{
		"$set": self.GetBSONModificationObj(),
	}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": self.ID}, updateQuery)
	if err != nil {
		c.Status(500)
		return err
	} else {
		c.Status(200)
	}
	return nil
}

func InventoryGetById(id primitive.ObjectID) (Models.Inventory, error) {
	collection := DBManager.SystemCollections.Inventory
	filter := bson.M{"_id": id}
	var self Models.Inventory
	_, results := Utils.FindByFilter(collection, filter)
	if len(results) == 0 {
		return self, errors.New("obj not found")
	}
	bsonBytes, _ := bson.Marshal(results[0]) // Decode
	bson.Unmarshal(bsonBytes, &self)         // Encode
	return self, nil
}
