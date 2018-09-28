package service

import (
	"net/http"

	"github.com/nytimes/gizmo/web"
)

func (s *JSONService) GetMostPopular(r *http.Request) (int, interface{}, error) {
	resourceType := web.Vars(r)["resourceType"]
	section := web.Vars(r)["section"]
	timeframe := web.GetUInt64Var(r, "timeframe")
	res, err := s.client.GetMostPopular(resourceType, section, uint(timeframe))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, res, nil
}
