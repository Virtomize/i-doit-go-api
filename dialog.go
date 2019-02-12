package goidoit

// GetDialog returns the dialog+ defined values
func (a *API) GetDialog(category, property string) (GenericResponse, error) {

	var CustomStruct interface{}
	CustomStruct = struct {
		Category string `json:"category"`
		Property string `json:"property"`
	}{category, property}

	data, err := a.Request("cmdb.dialog.read", CustomStruct)
	if err != nil {
		return GenericResponse{}, err
	}

	return TypeAssertResult(data)
}
