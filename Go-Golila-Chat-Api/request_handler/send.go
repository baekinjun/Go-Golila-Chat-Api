package request_handler

import (
	"fmt"
	utils "wooah/utils"
)

// const API_SERVER_URL = "http://localhost:1012"
const API_SERVER_URL = "http://dev.wooah-market.com"

func SendToAPIServer(accessToken string, roomId string, message string) (result map[string]interface{}, err error) {
	url := API_SERVER_URL + "/api/chat/" + roomId + "/"
	data := map[string]string{
		"content": message,
	}
	result, err = utils.HttpRequestPost(accessToken, url, data)
	fmt.Println(result)
	fmt.Println(err)
	return
}
