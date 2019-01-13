package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_decodeMessage_invalid_format(t *testing.T) {
	todo, err := decodeMessage("Buy milk")
	assert.Zero(t, todo)
	assert.NotNil(t, err)
}

func Test_getDateTimeFromString_convert_dateStr__success(t *testing.T) {
	dtm, err := getDateTimeFromString("TodAy", "")
	assert.Nil(t, err)
	duration, _ := time.ParseDuration("1s")
	assert.WithinDuration(t, time.Now(), dtm, duration)

	dtm, err = getDateTimeFromString("Tomorrow", "")
	assert.Nil(t, err)
	assert.WithinDuration(t, time.Now().AddDate(0, 0, 1), dtm, duration)

	dtm, err = getDateTimeFromString("2/5/18", "")
	assert.Nil(t, err)
	expDtm, err := time.Parse(time.RFC3339, "2018-05-02T"+time.Now().Format("15:04:05Z07:00"))
	assert.Nil(t, err)
	assert.WithinDuration(t, dtm, expDtm, duration)

}

func Test_getDateTimeFromString_convert_dateStr__fail(t *testing.T) {
	_, err := getDateTimeFromString("", "")
	assert.NotNil(t, err)

	_, err = getDateTimeFromString("12/31/18", "")
	assert.NotNil(t, err)
}

func Test_getDateTimeFromString_convert_timeStr_success(t *testing.T) {
	dtm, err := getDateTimeFromString("2/5/18", "15:04")
	assert.Nil(t, err)
	duration, _ := time.ParseDuration("1s")
	expDtm, err := time.Parse(time.RFC3339, "2018-05-02T15:04:00+07:00")
	assert.Nil(t, err)
	assert.WithinDuration(t, dtm, expDtm, duration)
}

func Test_getDateTimeFromString_convert_timeStr_fail(t *testing.T) {
	_, err := getDateTimeFromString("2/5/18", "30:04")
	assert.NotNil(t, err)
}
