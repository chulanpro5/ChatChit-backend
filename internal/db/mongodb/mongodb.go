package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"test-chat/internal/config"
	"test-chat/internal/entity"
)

type MongoDb struct {
	*mongo.Client
}

func NewMongoDb(config *config.Config) (*MongoDb, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/%s%s",
		config.MongoDb.Username, config.MongoDb.Password, config.MongoDb.ClusterURL, config.MongoDb.DatabaseName, config.MongoDb.Options)).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
		return nil, err
	}

	return &MongoDb{client}, nil
}

func (m *MongoDb) Ping() {
	fmt.Println("Pinging your deployment...")
	// Send a ping to confirm a successful connection
	if err := m.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (m *MongoDb) Close() {
	m.Disconnect(context.Background())
	fmt.Println("Connection to MongoDB closed.")
}

// InsertMessage Add message to database (collection: messages)
func (m *MongoDb) InsertMessage(message *entity.Message) error {
	collection := m.Database("ChatChit").Collection("Message")
	_, err := collection.InsertOne(context.Background(), *message)
	if err != nil {
		return err
	}
	return nil
}

// GetAllMessages Get all messages from database (collection: messages)
func (m *MongoDb) GetAllMessages() ([]entity.Message, error) {
	collection := m.Database("ChatChit").Collection("Message")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var messages []entity.Message
	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
