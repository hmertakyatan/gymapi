package helpers

import (
	"log"

	"github.com/google/uuid"
)

func StringToUUID(ID string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatal("Got an error when parsing UUID. ERROR: ", err)
		return uuid.Nil, err
	}
	return parsedID, nil
}
