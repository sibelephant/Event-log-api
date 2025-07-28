package database

import (
	"github.com/sibelephant/prisma/db"
)

var (
	//Client is a global Prisma client instance
	Client *db.PrismaClient
)

// Connect initializes the Prisma Client
func Connect() error {
	Client = db.NewClient()
	if err := Client.Connect(); err != nil {
		return err
	}
	return nil
}

// Disconnect closes the Prisma client connection
func Disconnect() {
	if Client != nil {
		Client.Disconnect()
	}
}
