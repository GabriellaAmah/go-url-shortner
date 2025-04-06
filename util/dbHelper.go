package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringId(id string) (primitive.ObjectID, error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}