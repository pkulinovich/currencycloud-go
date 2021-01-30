package currencycloud_go

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func QueryStruct(content interface{}) (*url.Values, error) {
	marshalContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	var params map[string]interface{}
	err = json.Unmarshal(marshalContent, &params)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for k, v := range params {
		k = strings.ToLower(k)
		var queryVal string
		switch t := v.(type) {
		case string:
			queryVal = t
		case float64:
			queryVal = strconv.FormatFloat(t, 'f', -1, 64)
		case time.Time:
			queryVal = t.Format(time.RFC3339)
		default:
			j, err := json.Marshal(v)
			if err != nil {
				continue
			}
			queryVal = string(j)
		}
		query.Add(k, queryVal)
	}
	return &query, nil
}
