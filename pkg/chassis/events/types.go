package events

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	evpb "github.com/circadence-official/galactus/api/gen/go/generic/events/v1"
)

func AggregateTypeAsString(aggregateType evpb.AggregateType) string {
	if aggregateType == 0 {
		return "*"
	}
	v := strings.Split(reflect.TypeOf(aggregateType).String(), ".")[0]
	t := fmt.Sprintf("%s.%s", v, aggregateType.String())
	return t
}

func EventTypeAsString(eventType *evpb.EventType) string {
	if eventType == nil {
		return "*"
	}
	t := reflect.TypeOf(eventType.GetCode()).String()
	t = strings.TrimPrefix(t, "*")
	return t
}

func EventCodeAsString(eventCode interface{}) string {
	if eventCode == nil {
		return "*"
	}
	code := fmt.Sprintf("%v", eventCode)
	return code
}

func ValidateEventCode(eventCode interface{}) error {
	if eventCode == nil {
		return nil
	}
	// convert code to string
	code := fmt.Sprintf("%v", eventCode)
	if !strings.Contains(code, "EVENT_CODE") {
		return errors.New("invalid event code: must contain the substring 'EVENT_CODE'")
	}
	return nil
}
