package chronos

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseUnitToken(cronToken string) (*cronUnit, error) {
	if cronToken == "*" {
		return &cronUnit{}, nil
	}

	if strings.Contains(cronToken, ",") {
		valueTokens := strings.Split(cronToken, ",")

		var values []int
		for _, valueToken := range valueTokens {
			if num, err := strconv.Atoi(valueToken); err == nil {
				values = append(values, num)
			} else {
				return nil, err
			}
		}

		sort.Ints(values)

		return &cronUnit{listed, values}, nil
	} else if strings.Contains(cronToken, "-") {
		valueTokens := strings.Split(cronToken, "-")

		if len(valueTokens) != 2 {
			return nil, fmt.Errorf("unexpected number of upper and lower bounds for range: %d", len(valueTokens))
		}

		var values []int
		for _, valueToken := range valueTokens {
			if num, err := strconv.Atoi(valueToken); err == nil {
				values = append(values, num)
			} else {
				return nil, err
			}
		}

		return &cronUnit{ranged, values}, nil
	} else if strings.Contains(cronToken, "/") {
		valueTokens := strings.Split(cronToken, "/")

		if len(valueTokens) != 2 {
			return nil, fmt.Errorf("unexpected number of values for step value: %d", len(valueTokens))
		}

		if valueTokens[0] != "*" {
			return nil, fmt.Errorf("expected wildcard [*] got [%s]", valueTokens[0])
		}

		if num, err := strconv.Atoi(valueTokens[1]); err == nil {
			return &cronUnit{stepped, []int{num}}, nil
		} else {
			return nil, err
		}
	} else {
		if num, err := strconv.Atoi(cronToken); err == nil {
			return &cronUnit{listed, []int{num}}, nil
		} else {
			return nil, err
		}
	}
}
