package handlers

import (
	"superapp/config"
	st "superapp/internal/storage"

	)



var Storage *st.Storage
var Config *config.Cfg
//var MongoCollections=map[string]interface{}{}

func PreSettings(settings map[string]interface{}){

	Storage = settings["storage"].(*st.Storage)
	// MongoCollections["test_coll"] = Storage.Mongo.GetColl("test", "test_coll")
	Config = settings["config"].(*config.Cfg)
	//InitSettings()
	
}
