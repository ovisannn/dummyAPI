package connection

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	// Setup the mgm default config
	_ = mgm.SetDefaultConfig(nil, "dummyDB", options.Client().ApplyURI("mongodb://localhost:27017"))
}
