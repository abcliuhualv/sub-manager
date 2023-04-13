package linkutil

import "net/url"

func ParseVlessLink(link string) (map[string]string, error) {
	parsed, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	uuid := parsed.User.Username()
	address := parsed.Hostname()
	port := parsed.Port()
	queryParams := parsed.RawQuery
	remark := parsed.Fragment

	paramMap := map[string]string{
		"uuid":        uuid,
		"address":     address,
		"port":        port,
		"queryParams": queryParams,
		"remark":      remark,
	}

	return paramMap, nil
}

func BuildVlessLink(paramMap map[string]string) (string, error) {
	parsedUrl := url.URL{
		Scheme:   "vless",
		User:     url.User(paramMap["uuid"]),
		Host:     paramMap["address"] + ":" + paramMap["port"],
		RawQuery: paramMap["queryParams"],
		Fragment: paramMap["remark"],
	}
	return parsedUrl.String(), nil
}
