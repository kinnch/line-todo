package todo

import (
	"errors"
	"fmt"
	"github.com/kinnch/line-todo/models"
	"github.com/kinnch/line-todo/utilities/sysVar"
	"strings"
	"time"
)

func TodoController(user_id string, message string) string {
	return fmt.Sprintf("we recieve your message: \"%s\" from user: %s \n", message, user_id)
}
func decodeMessage(user_id string, message string) (models.Todo, error) {
	var todo models.Todo
	splted := strings.Split(message, ":")
	// we default time if not specified
	if len(splted) < 2 {
		return todo, errors.New("invalid format message")
	}
	if splted[0] == "" {
		return todo, errors.New("task provided is invalid")
	}
	dtm, err := getDateTimeFromString(splted[1], splted[2])
	if err != nil {
		return todo, err
	}
	todo = models.Todo{
		Task:   splted[0],
		Time:   dtm,
		UserID: user_id,
	}
	return todo, nil
}

func getDateTimeFromString(dateStr string, timeStr string) (time.Time, error) {
	var date string
	var t string
	if dateStr == "" {
		return time.Time{}, errors.New("date is required")
	} else if strings.ToUpper(dateStr) == "TODAY" {
		date = time.Now().In(sysVar.Location()).Format("2006-01-02")
	} else if strings.ToUpper(dateStr) == "TOMORROW" {
		date = time.Now().AddDate(0, 0, 1).In(sysVar.Location()).Format("2006-01-02")
	} else {
		d, err := time.ParseInLocation("2/1/06", dateStr, sysVar.Location())
		if err != nil {
			return time.Time{}, err
		}
		date = d.Format("2006-01-02")
	}

	if timeStr == "" {
		t = time.Now().In(sysVar.Location()).Format("15:04:05Z07:00")
	} else {
		tmp, err := time.ParseInLocation("15:04", timeStr, sysVar.Location())
		if err != nil {
			return time.Time{}, err
		}
		t = tmp.Format("15:04:05+07:00")
	}

	dtm, err := time.Parse(time.RFC3339, date+"T"+t)
	if err != nil {
		return time.Time{}, err
	}
	return dtm, nil

}
