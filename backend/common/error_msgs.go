package common

var ERR_COMMUNITY_NOT_FOUND APIMessage = APIMessage{Message: "No record found for given community name"}
var ERR_COMMENTS_NOT_FOUND APIMessage = APIMessage{Message: "No record found for given post id"}
var ERR_COMMENT_NOT_VOTABLE APIMessage = APIMessage{Message: "Comment is not votable"}
var ERR_INCORRECT_CREDENTIALS APIMessage = APIMessage{Message: "Incorrect Credentials"}
var ERR_USERNAME_ALREADY_EXISTS APIMessage = APIMessage{Message: "Username Already Exists"}
var ERR_COMMUNITY_ALREADY_EXISTS APIMessage = APIMessage{Message: "Community with that name already exists"}
var ERR_PARENT_COMMENT_IS_DELETED APIMessage = APIMessage{Message: "Parent comment is deleted"}
