package goidoit

// GetReport returns the report given by its id
func (a *API) GetReport(RepID int) (GenericResponse, error) {

	CustomStruct := struct {
		ID int `json:"id"`
	}{RepID}

	data, err := a.Request("cmdb.reports.read", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}
