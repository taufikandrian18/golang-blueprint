package utility

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gitlab.com/wit-id/project-latihan/common/constants"

	"github.com/dustin/go-humanize"
)

// PrettyPrint will transform struct data as json string for nicer log ...
func PrettyPrint(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}

// PrettyPrintWithoutIndent will transform struct data as json string for nicer log ...
func PrettyPrintWithoutIndent(data interface{}) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}

// FormatPhoneNumber default for prefix ID.
func FormatPhoneNumber(phoneNumber string) (response string) {
	prefix := phoneNumber[0:3]
	if prefix != "+62" {
		phoneNumber = strings.Replace(phoneNumber, "0", "+62", 1)
	}

	response = phoneNumber

	return
}

// FormatRupiah to convert number 10000 to Rp 10.000.
func FormatRupiah(amount int) string {
	// convert to positive if number negative
	amountConv := math.Abs(float64(amount))
	humanizeValue := humanize.Comma(int64(amountConv))
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)

	return "Rp " + stringValue
}

// GenerateOrderNumber will generate number order ex: 000000001.
func GenerateOrderNumber(id, length int) (idString string) {
	idString = strconv.Itoa(id)

	numOfZero := length - len(idString)

	for i := 1; i <= numOfZero; i++ {
		idString = "0" + idString
	}

	return
}

func GenerateInvoices(prefix string, id, length int) (invoices string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)

	invoices = prefix + "/" + timeNow.Format("20060102") + "/" + GenerateOrderNumber(id, length)

	return
}

func RandomNumber(low, hi int) int {
	return low + rand.Intn(hi-low)
}

// GenerateSlug will change title to slug format
func GenerateSlug(title string) string {
	// Convert to lowercase and replace all non-alphanumeric characters with hyphen
	slug := regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(strings.ToLower(title), "-")

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	slug = slug + "-" + strconv.Itoa(RandomNumber(1, 999))

	return slug
}

func GenerateSlugARMS(number int32) string {

	str := fmt.Sprintf("ARMS%05d", number)

	return str
}

func PropertyLevel(levelString string) (levelNumber string) {
	if levelString == constants.PropertyLevelCollection {
		levelNumber = constants.EnumPropertyLevelCollection
	}
	if levelString == constants.PropertyLevelEconomy {
		levelNumber = constants.EnumPropertyLevelEconomy
	}
	if levelString == constants.PropertyLevelMidScale {
		levelNumber = constants.EnumPropertyLevelMidScale
	}
	if levelString == constants.PropertyLevelLifestyle {
		levelNumber = constants.EnumPropertyLevelLifestyle
	}
	if levelString == constants.PropertyLevelUpScale {
		levelNumber = constants.EnumPropertyLevelUpScale
	}
	if levelString == constants.PropertyLevelFoodBeverages {
		levelNumber = constants.EnumPropertyLevelFoodBeverages
	}

	return
}

// ActivityBodyMessage Create activity log body message
func ActivityBodyMessage(activityType string, mapData map[string]string) string {
	var message string

	switch activityType {
	case constants.ActivityAddBoardTask:
		message = fmt.Sprintf("%s added %s to %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle],
			mapData[constants.ActivityLoggingColumnTitle])
	case constants.ActivityArchiveBoardTask:
		message = fmt.Sprintf("%s archived %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle])
	case constants.ActivityMoveBoardTask:
		message = fmt.Sprintf("%s moved %s from %s to %s", mapData[constants.ActivityLoggingFullName],
			mapData[constants.ActivityLoggingTaskTitle],
			mapData[constants.ActivityLoggingColumnTitle],
			mapData[constants.ActivityLoggingColumnTitle2])
	}

	return message
}
