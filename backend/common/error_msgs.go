package common

var ERR_COMMUNITY_NOT_FOUND APIMessage = APIMessage{Message: "No record found for given community name"}
var ERR_INCORRECT_CREDENTIALS APIMessage = APIMessage{Message: "Incorrect Credentials"}
var ERR_USERNAME_ALREADY_EXISTS APIMessage = APIMessage{Message: "Username Already Exists"}
var ERR_COMMUNITY_ALREADY_EXISTS APIMessage = APIMessage{Message: "Community with that name already exists"}
