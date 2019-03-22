package goidoit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// export getID
func TgetID() int {
	return getID()
}

// reset id
func TResetID() {
	id = 0
}

// export debug func
func TdebugPrint(format string, a ...interface{}) (n int, err error) {
	return debugPrint(format, a...)
}

func TgetDebug() bool {
	return debug
}

func idoitAPIStub() *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var req Request
		err = json.Unmarshal(b, &req)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		switch req.Method {
		case "idoit.login":
			resp = `{"jsonrpc":"2.0","result":{"result":true,"userid":"1","name":"username","mail":"","username":"username","session-id":"mysessionid","client-id":"1","client-name":"client"},"id":1}`
		case "test", "idoit.search", "idoit.version", "cmdb.objects.read", "cmdb.objects.create", "cmdb.objects.update", "cmdb.objects.delete", "cmdb.objets.quickpurge", "idoit.reports.read", "idoit.dialog.read", "cmdb.category.read", "cmdb.category.create", "cmdb.category.update", "cmdb.category.delete":
			resp = `{"jsonrpc":"2.0","result":[]}`
		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(resp))
	}))
}
