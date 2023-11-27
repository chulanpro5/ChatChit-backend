package config

type MongoDb struct {
	Username     string
	Password     string
	ClusterURL   string
	DatabaseName string
	Options      string // Additional options if needed
}
