package Controllers

import (
	"context"
	"encoding/json"
	"errors"

	"SEEN-TECH-VAI21-BACKEND-GO/DBManager"
	"SEEN-TECH-VAI21-BACKEND-GO/Models"
	"SEEN-TECH-VAI21-BACKEND-GO/Utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SalesNew(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Sales
	var self Models.Sales
	c.BodyParser(&self)
	err := self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}

	res, err := collection.InsertOne(context.Background(), self)
	if err != nil {
		c.Status(500)
		return err
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)
	return nil
}

func SalesGetAll(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Sales

	// Fill the received search obj data
	var self Models.SalesSearch
	c.BodyParser(&self)

	var results []bson.M
	b, results := Utils.FindByFilter(collection, self.GetSalesSearchBSONObj())
	if !b {
		err := errors.New("db error")
		c.Status(500).Send([]byte(err.Error()))
		return err
	}

	// Decode
	response, _ := json.Marshal(
		bson.M{"result": results},
	)
	c.Set("Content-Type", "application/json")
	c.Status(200).Send(response)

	return nil
}

func SalesGetAllPopulated(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Sales
	var self Models.SalesSearch
	c.BodyParser(&self)
	b, results := Utils.FindByFilter(collection, self.GetSalesSearchBSONObj())
	if !b {
		c.Status(500)
		return errors.New("object is not found")
	}
	byteArr, _ := json.Marshal(results)
	var ResultDocs []Models.Sales
	json.Unmarshal(byteArr, &ResultDocs)
	populatedResult := make([]Models.SalesPopulated, len(ResultDocs))

	for i, v := range ResultDocs {
		populatedResult[i], _ = SalesGetByIdPopulated(v.ID, &v)
	}
	allpopulated, _ := json.Marshal(bson.M{"result": populatedResult})
	c.Set("Content-Type", "application/json")
	c.Send(allpopulated)
	return nil
}

func SalesGetByIdPopulated(objID primitive.ObjectID, ptr *Models.Sales) (Models.SalesPopulated, error) {
	var SalesDoc Models.Sales
	if ptr == nil {
		SalesDoc, _ = SalesGetById(objID)
	} else {
		SalesDoc = *ptr
	}

	populatedResult := Models.SalesPopulated{}
	populatedResult.CloneFrom(SalesDoc)

	var err error

	// populate for CustomerRef
	if SalesDoc.CustomerRef != primitive.NilObjectID {
		populatedResult.CustomerRef, err = ContactGetById(SalesDoc.CustomerRef)
		if err != nil {
			return populatedResult, err
		}
	}

	// populate for InventoryRef
	if SalesDoc.InventoryRef != primitive.NilObjectID {
		populatedResult.InventoryRef, err = InventoryGetById(SalesDoc.InventoryRef)
		if err != nil {
			return populatedResult, err
		}
	}

	// populate for PriceListRef
	if SalesDoc.PriceListRef != primitive.NilObjectID {
		populatedResult.PriceListRef, err = PriceListGetById(SalesDoc.PriceListRef)
		if err != nil {
			return populatedResult, err
		}
	}

	// populate for QuotationRef
	if SalesDoc.QuotationRef != primitive.NilObjectID {
		populatedResult.QuotationRef, err = SalesGetById(SalesDoc.QuotationRef)
		if err != nil {
			return populatedResult, err
		}
	}

	// populate for SalesOrderRef
	if SalesDoc.SalesOrderRef != primitive.NilObjectID {
		populatedResult.SalesOrderRef, err = SalesGetById(SalesDoc.SalesOrderRef)
		if err != nil {
			return populatedResult, err
		}
	}

	// populate for Products of Invoice row array
	populatedResult.Products = make([]Models.InvoiceRowPopulated, len(SalesDoc.Products))

	for i := range SalesDoc.Products {
		populatedResult.Products[i].CloneFrom(SalesDoc.Products[i])
		populatedResult.Products[i].ProductRef, err = ProductGetById(SalesDoc.Products[i].ProductRef)
		if err != nil {
			return populatedResult, err
		}
	}

	return populatedResult, nil
}

func SalesGetById(id primitive.ObjectID) (Models.Sales, error) {
	collection := DBManager.SystemCollections.Sales
	filter := bson.M{"_id": id}
	var self Models.Sales
	_, results := Utils.FindByFilter(collection, filter)
	if len(results) == 0 {
		return self, errors.New("obj not found")
	}
	bsonBytes, _ := bson.Marshal(results[0]) // Decode
	bson.Unmarshal(bsonBytes, &self)         // Encode
	return self, nil
}

// for modify Sales
func SalesModify(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Sales
	if c.Params("id") == "" {
		c.Status(404)
		return errors.New("id param needed")
	}
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))

	_, err := SalesGetById(objID)
	if err != nil {
		return err
	}
	var self Models.Sales
	c.BodyParser(&self)
	err = self.Validate()
	if err != nil {
		c.Status(500)
		return err
	}
	// need to add validate here for product using loop
	for _, product := range self.Products {
		err := product.Validate()
		if err != nil {
			return err
		}
	}

	updateData := bson.M{
		"$set": self.GetModifcationBSONObj(),
	}
	_, updateErr := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, updateData)
	if updateErr != nil {
		c.Status(500)
		return errors.New("an error occurred when modifying product")
	}

	c.Status(200).Send([]byte("Modified Successfully"))
	return nil
}

func SetStatus(c *fiber.Ctx) error {
	collection := DBManager.SystemCollections.Sales
	if c.Params("id") == "" || c.Params("new_status") == "" || c.Params("type") == "" {
		c.Status(404)
		return errors.New("all params not sent correctly")
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))

	salesDB, _ := SalesGetById(objID)
	newValue := salesDB.Status

	if c.Params("type") != salesDB.Type {
		c.Status(500)
		return errors.New("type doesn't match")
	}

	// for Quatation
	if c.Params("type") == "Quatation" {
		if salesDB.Status == "created" && (c.Params("new_status") == "sent" || c.Params("new_status") == "decline") {
			newValue = c.Params("new_status")
		} else if salesDB.Status == "sent" && (c.Params("new_status") == "confirmed" || c.Params("new_status") == "decline") {
			newValue = c.Params("new_status")
		} else {
			c.Status(500)
			return errors.New("can't change status")
		}
	}

	// for order
	if c.Params("type") == "order" {
		if salesDB.Status == "created" && (c.Params("new_status") == "chiped" || c.Params("new_status") == "decline") {
			newValue = c.Params("new_status")
		} else {
			c.Status(500)
			return errors.New("can't change status")
		}
	}

	// need to add Delivery type status here

	updateData := bson.M{
		"$set": bson.M{
			"status": newValue,
		},
	}
	_, updateErr := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, updateData)
	if updateErr != nil {
		c.Status(500)
		return errors.New("an error occurred when modifying quatation status")
	} else {
		c.Status(200).Send([]byte("Modified Successfully"))
		return nil
	}
}
