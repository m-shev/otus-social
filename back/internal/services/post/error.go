package post

import "errors"

var (
	ErrorSendMessageToQueue = errors.New("got an error when trying to send a message in the queue")
)
