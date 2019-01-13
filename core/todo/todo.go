package todo

import (
	"errors"
	"fmt"
	"github.com/kinnch/line-todo/models"
	"strings"
)



func TodoController(user_id string ,message string) string {
	return fmt.Sprintf("we recieve your message: \"%s\" from user: %s \n", message, user_id)
}
func decodeMessage(message string) (models.Todo , error) {
	splted := strings.Split(message, ":")
	// we default time if not specified
	if len(splted) < 2 {
		return models.Todo{}, errors.New("invalid format message")
	}
	return models.Todo{},nil
}