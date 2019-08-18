package connectors

import "models"

var Mongo mongo

type mongo struct {

}

func InitMongo() *mongo{


	Mongo = mongo{
	}

	return &Mongo
}

func(m *mongo) GetEntityById(collection, id string) interface{}{
	qa := models.QuizAnswers{}
	return qa.GetMock()
}

func(m *mongo) UpdateEntityByIdAndField(collection, id, field, value string) error{
	return nil
}