# goserverapi
Server Api build programing language Go Lang
# code learn 

w.Header().Set("Content-Type", "application/json")

w.WriteHeader(http.StatusOK)

func GetConfigServer() *ServerConfig {
	dataConfig, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Print(err)
	}

	obj := new(ServerConfig)

	err = json.Unmarshal(dataConfig, &obj)

	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

# handle response status code controller 
# load file i18n save lang  
# handle insert
# handle update
# handle delete