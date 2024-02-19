package service

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const DbUrl = "https://awesome-ionutvelicu.turso.io/v2/pipeline"
const DbAuthToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOiIyMDI0LTAxLTIwVDEyOjU5OjUyLjg0ODQ1NTg4NFoiLCJpZCI6ImJiYTFiMTI3LWI3OGUtMTFlZS1iMTg5LWE2MDMzNTNjNjBiMyJ9.j6Vj9yR7qLcKFO8_EzG_Rr74gmUJJIH7Rhx0NqNaIOd4WGJr_tIfuuAFPWHqQIF-xI2RiGoyjuH93oBgCol5AA"

var httpClient = &http.Client{}

type DbStatement struct {
	Sql string `json:"sql"`
}

type DbCommand struct {
	Tp   string       `json:"type"`
	Stmt *DbStatement `json:"stmt"`
}

var CloseCommand = DbCommand{"close", nil}

type DbRow []struct {
	Value string `json:"value"`
}

func DbCall(query string) (string, error) {
	command := DbCommand{
		Tp:   "execute",
		Stmt: &DbStatement{Sql: query},
	}

	commands := []DbCommand{command, CloseCommand}
	data, _ := json.Marshal(map[string][]DbCommand{"requests": commands})

	req, err := http.NewRequest("POST", DbUrl, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+DbAuthToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ExtractRows(content string, key string) ([]DbRow, error) {
	str := ExtractValue(content, key)
	t := make([]DbRow, 10)
	err := json.Unmarshal([]byte(str), &t)
	return t, err
}

func ExtractValue(content string, key string) string {
	result := gjson.Get(content, "results.0.response.result."+key)
	return result.Raw
}
