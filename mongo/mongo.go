package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION_NAME = "sessions"
)

//Connector wraps connection to mongoDB
type Connector struct {
	client *mongo.Client
	dbName string
}

//NewConnector creates a new connector to mongoDB
func NewConnector(dbURI, dbName string) (*Connector, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbURI))
	if err != nil {
		return &Connector{}, err
	}
	return &Connector{
		client: client,
		dbName: dbName,
	}, nil
}

//CreateOne adds a document to a collection
func (c *Connector) StartSession(sessionId string, document interface{}) error {
	col := c.client.Database(c.dbName).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := col.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	//id := res.InsertedID.(primitive.ObjectID).String()
	return nil
}

func (c *Connector) RetrieveData(sessionId string, result interface{}) {
	col := c.client.Database(c.dbName).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"sessionId", sessionId}}
	col.FindOne(ctx, filter).Decode(result)
}

func (c *Connector) EndSession(sessionId string) error {
	col := c.client.Database(c.dbName).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"sessionId", sessionId}}
	_, err := col.DeleteOne(ctx, filter)
	return err
}

func (c *Connector) UpdateSession(sessionId string, document interface{}) error {
	col := c.client.Database(c.dbName).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"sessionId", sessionId}}
	update := bson.D{
		{"$set", document},
	}
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// func (c *Connector) FindOne(collectionName string, ID string, r interface{}, opts ...*options.FindOneOptions) error {
// 	col := c.client.Database(c.dbName).Collection(collectionName)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	filter := bson.D{{"id", string(ID)}}
// 	return col.FindOne(ctx, filter, opts...).Decode(r)
// }

// //Update updates existing document in the database
// func (c *Connector) Update(ID string, collectionName string, changes interface{}) error {
// 	col := c.client.Database(c.dbName).Collection(collectionName)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	filter := bson.D{{"id", string(ID)}}
// 	update := bson.D{
// 		bson.E{"$set", changes},
// 	}
// 	_, err2 := col.UpdateOne(ctx, filter, update)
// 	return err2
// }

//CountDocuments returns count of filtered documents
func (c *Connector) SessionExists(sessionId string) (bool, error) {
	filter := bson.D{{"sessionId", sessionId}}
	col := c.client.Database(c.dbName).Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
