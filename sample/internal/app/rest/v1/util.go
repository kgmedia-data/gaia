package v1

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func convertSlice[T any, R any](input []T, convertFunc func(T) R) []R {
	var result []R
	for _, item := range input {
		result = append(result, convertFunc(item))
	}
	return result
}

func queryParamToOffsetLimit(c echo.Context, useDefault bool) (int, int) {
	var page, limit int

	err := echo.QueryParamsBinder(c).Int("page", &page).
		Int("limit", &limit).BindError()

	if err != nil {
		return 0, 10
	}

	if page < 1 {
		if useDefault {
			return 0, 10
		} else {
			return 0, 0
		}
	}

	return (page - 1) * limit, limit
}

func getDepartmentId(c echo.Context) int {
	deptId, err := strconv.Atoi(c.Param(DEPT_ID))
	if err != nil {
		return 0
	}

	return deptId
}

func getEmployeeId(c echo.Context) int {
	emplId, err := strconv.Atoi(c.Param(EMPLOYEE_ID))
	if err != nil {
		return 0
	}

	return emplId
}

func paramInt(c echo.Context, key string, defaultValue int) int {
	var value int
	err := echo.QueryParamsBinder(c).Int(key, &value).BindError()
	if err != nil {
		return defaultValue
	}
	return value
}
