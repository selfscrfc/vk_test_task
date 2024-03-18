package api_models

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

type CreateFilmParams struct {
	FilmId      string    `json:"film_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rate        int       `json:"rate"`
	Actors      []string  `json:"actors"`
}

type UpdateFilmParams struct {
	FilmId      string    `json:"film_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rate        int       `json:"rate"`
	Actors      []string  `json:"actors"`
}

type GetFilmsParams struct {
	SortBy      int `json:"sort_by"`
	IsAscending int `json:"is_ascending"`
}

type FilmAndActors struct {
	FilmId      string         `json:"film_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Rate        int            `json:"rate"`
	ReleaseDate string         `json:"release_date"`
	Actors      sql.NullString `json:"actors"` //обычный массив не подходит, т.к. запрос возвращает строку. Строка не подходит т.к. postgres экранирует кавычки у строк с пробелами, и не экранирует ничего у строек без пробелов.
}

func (a FilmAndActors) MarshallJSON() (map[string]interface{}, error) {
	trimmed := strings.TrimPrefix(a.Actors.String, "{")
	trimmed = strings.TrimSuffix(trimmed, "}")

	result := strings.Split(trimmed, ",")

	jsonMap := make(map[string]interface{})

	jsonMap["film_id"] = a.FilmId
	jsonMap["name"] = a.Name
	jsonMap["description"] = a.Description
	jsonMap["rate"] = a.Rate
	jsonMap["release_date"] = a.ReleaseDate

	if len(result) > 0 && result[0] == "NULL" {
		jsonMap["actors"] = []string{}
	} else {
		jsonMap["actors"] = result
	}

	return jsonMap, nil
}

func (r GetFilmsResponse) MarshallJSON() ([]byte, error) {
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

type GetFilmsResponse struct {
	Response []FilmAndActors `json:"response"`
}

type DeleteFilmParams struct {
	FilmId string `json:"film_id"`
}

type SearchFilmParams struct {
	Name      string `json:"name"`
	ActorName string `json:"actor_name"`
}

type SearchFilmResponse struct {
	Response []FilmAndActors `json:"response"`
}

func (r SearchFilmResponse) MarshallJSON() ([]byte, error) {
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
