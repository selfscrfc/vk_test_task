package api_models

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

type CreateActorParams struct {
	ActorId string    `json:"actor_id"`
	Name    string    `json:"name"`
	Sex     int       `json:"sex"`
	Birth   time.Time `json:"birth"`
}

type ActorAndFilms struct {
	ActorId string         `json:"actor_id"`
	Name    string         `json:"name"`
	Sex     int            `json:"sex"`
	Birth   time.Time      `json:"birth"`
	Films   sql.NullString `json:"films"` //обычный массив не подходит, т.к. запрос возвращает строку. Строка не подходит т.к. postgres экранирует кавычки у строк с пробелами, и не экранирует ничего у строек без пробелов.
}

func (a ActorAndFilms) MarshallJSON() (map[string]interface{}, error) {
	trimmed := strings.TrimPrefix(a.Films.String, "{")
	trimmed = strings.TrimSuffix(trimmed, "}")

	result := strings.Split(trimmed, ",")

	jsonMap := make(map[string]interface{})

	jsonMap["actor_id"] = a.ActorId
	jsonMap["name"] = a.Name
	jsonMap["sex"] = a.Sex
	jsonMap["birth"] = a.Birth

	if len(result) > 0 && result[0] == "NULL" {
		jsonMap["films"] = []string{}
	} else {
		jsonMap["films"] = result
	}

	return jsonMap, nil
}

type GetActorsResponse struct {
	Response []ActorAndFilms `json:"response"`
}

func (r GetActorsResponse) MarshallJSON() ([]byte, error) {
	jsonMap := make(map[string]interface{})

	var result []map[string]interface{}

	for _, v := range r.Response {
		jsonSingle, err := v.MarshallJSON()
		if err != nil {
			return nil, err
		}
		result = append(result, jsonSingle)
	}

	jsonMap["response"] = result

	return json.Marshal(jsonMap)
}

type UpdateActorParams struct {
	ActorId string    `json:"actor_id"`
	Name    string    `json:"name"`
	Sex     int       `json:"sex"`
	Birth   time.Time `json:"birth"`
}

type DeleteActorParams struct {
	ActorId string `json:"actor_id"`
}
