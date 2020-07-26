/*
 * oas3_go_flutter_exam Employee
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// A EmployeeApiController binds http requests to an api service and writes the service results to the http response
type EmployeeApiController struct {
	service EmployeeApiServicer
}

// NewEmployeeApiController creates a default api controller
func NewEmployeeApiController(s EmployeeApiServicer) Router {
	return &EmployeeApiController{service: s}
}

// Routes returns all of the api route for the EmployeeApiController
func (c *EmployeeApiController) Routes() Routes {
	return Routes{
		{
			"GetEmployee",
			strings.ToUpper("Get"),
			"/api/employee",
			c.GetEmployee,
		},
	}
}

// GetEmployee - get employee information
func (c *EmployeeApiController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	/*
		query := r.URL.Query()
		employeeId := query.Get("employeeId")
		result, err := c.service.GetEmployee(employeeId)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		EncodeJSONResponse(result, nil, w)
	*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	employeeID := r.URL.Query()["employeeId"]

	employee := Employee{
		Id:        employeeID[0],
		FirstName: "ryuichi " + employeeID[0],
		LastName:  "daigo " + employeeID[0],
		Salary:    12000000}
	json.NewEncoder(w).Encode(employee)
}

/*
// GetEmployee - get employee information
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	employeeID := r.URL.Query()["employeeId"]

	employee := Employee{
		Id:        employeeID[0],
		FirstName: "ryuichi " + employeeID[0],
		LastName:  "daigo " + employeeID[0],
		Salary:    12000000}
	json.NewEncoder(w).Encode(employee)
}
*/
