package repository

import (
	"fmt"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
var (
	Countyrepository countyrepository = countyrepository{}
)

type countyrepository struct{}

func (r *countyrepository) Create(county *model.County) (*httperrors.HttpError) {
	if err1 := county.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("county")
	_, err := collection.InsertOne(ctx, county)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create county Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *countyrepository) GetOne(id string) (county *model.County, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("county")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&county)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return county, nil	
}

func (r *countyrepository) GetAll(countys []model.County) ([]model.County, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("county")
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if err != nil { 
		return nil,	httperrors.NewNotFoundError("no results found")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
	err := cur.Decode(&countys)
		if err != nil { 
			return nil,	httperrors.NewNotFoundError("Error while decoding results!")
		}
	// do something with result....
	}
	if err := cur.Err(); err != nil {
		return nil,	httperrors.NewNotFoundError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}	
	DbClose(c)
    return countys, nil

}

func (r *countyrepository) Update(id string, county *model.County) (*httperrors.HttpError) {
	ucounty := &model.County{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("county")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&ucounty)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if county.Name  == "" {
		county.Name = ucounty.Name
	}
	if county.Title  == "" {
		county.Title = ucounty.Title
	}

	if county.Description  == "" {
		county.Description = ucounty.Description
	}
	_, err = collection.UpdateOne(ctx, filter, county)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of county Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r countyrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("county")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}

