package utility

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"gitlab.com/wit-id/project-latihan/common/constants"
	"gitlab.com/wit-id/project-latihan/common/httpservice"

	"github.com/lib/pq"
)

func ParseSqlError(err error) string {
	switch e := err.(type) {
	case *pq.Error:
		switch e.Code {
		case "23502":
			//not-null constraint violation
			log.Println("23502")
			b, _ := json.Marshal(e)
			log.Println(string(b))
			return fmt.Sprint("detail", e.Detail, "message", e.Message)
		case "23503":
			//foreign key constraint violation
			replacer := strings.NewReplacer("(", "", ")", "", "=", " ", "Key ", "")
			errMsg := replacer.Replace(e.Detail)
			if strings.Contains(errMsg, "is not present in table") {
				splitted := strings.Split(errMsg, "is not present in table")
				errMsg = fmt.Sprint(splitted[0], " ", "is not exists")
			}

			return fmt.Sprint(errMsg, "|", errMsg)
		case "23505":
			//unique constraint violation
			replacer := strings.NewReplacer("(", "", ")", "", "=", " ", "Key ", "")
			return fmt.Sprint(replacer.Replace(e.Detail), "|", replacer.Replace(e.Detail))
		case "23514":
			// check constraint violation
			log.Println("23514")
			b, _ := json.Marshal(e)
			log.Println(string(b))
			return fmt.Sprint("detail", e.Detail, "message", e.Message)
		case "22P02":
			log.Println("22P02")
			b, _ := json.Marshal(e)
			log.Println(string(b))
			return fmt.Sprint("invalid input|invalid input")
		default:
			log.Println(e.Code)
			b, _ := json.Marshal(e)
			log.Println(string(b))
			return fmt.Sprint("detail", e.Detail, "message", e.Message)
		}
	}

	return err.Error()
}

// ParseError Parsing error (look up to postgresql error). Returning error.
func ParseError(err error) error {
	switch e := err.(type) {
	case *pq.Error:
		switch e.Code {
		case "23505":
			if e.Table == constants.TableEmployee {
				switch e.Constraint {
				case constants.UniqueNIKConstraint:
					return httpservice.ErrNIKAlreadyExist
				case constants.UniqueIDCardConstraint:
					return httpservice.ErrIDCardAlreadyExist
				case constants.UniqueNPWPConstraint:
					return httpservice.ErrNPWPAlreadyExist
				case constants.UniqueEmailConstraint:
					return httpservice.ErrEmailAlreadyExist
				case constants.UniquePhoneNumberConstraint:
					return httpservice.ErrPhoneNumberAlreadyExist
				}
			}
		}
	}

	return httpservice.ErrInternalServerError
}
