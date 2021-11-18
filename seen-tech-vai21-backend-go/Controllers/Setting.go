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
)

func SettingNew(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Setting
	var self Models.Setting
	c.BodyParser(&self)
	res, err := collection.InsertOne(context.Background(), self)
	if err != nil {
		c.Status(500)
		return err
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)
	return nil
}

func settingGetAll(self *Models.SettingSearch) ([]bson.M, error) {
	collection := DBManager.SystemCollections.Setting
	var results []bson.M
	b, results := Utils.FindByFilter(collection, self.GetSettingSearchBSONObj())
	if !b {
		return results, errors.New("no object found")
	}
	return results, nil
}

func SettingGetAll(c *fiber.Ctx) error {
	var self Models.SettingSearch
	c.BodyParser(&self)
	results, err := settingGetAll(&self)
	if err != nil {
		c.Status(500)
		return err
	}
	response, _ := json.Marshal(bson.M{"result": results})
	c.Set("Content-Type", "application/json")
	c.Status(200).Send(response)
	return nil
}
