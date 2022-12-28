package modifiers

import (
	"net/http"
)

type response struct {
	Data interface{} `json:"data"`
}

// ExampleProxyResponseHandler is the handler for after a downstream file
// upload request is successfully returned from the API instance
func ExampleProxyResponseHandler(resp *http.Response) error {

	var err error

	// Can optionally change the response here
	// if you like
	//
	// data, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	r.StatusCode = http.StatusInternalServerError
	// 	return err
	// }

	// defer r.Body.Close()

	// logrus.Info("handle response")

	// fmt.Println(string(data))

	// EG:
	// err = json.Unmarshal(data, &reqData)
	// if err != nil {
	// 	r.StatusCode = http.StatusInternalServerError
	// 	data = []byte(err.Error())
	// 	return err
	// }

	// req := requests.NewCreateRequest(&reqData.Data)

	// err = req.HandleRequest()
	// if err != nil {
	// 	r.StatusCode = http.StatusInternalServerError
	// 	data = []byte(err.Error())
	// 	return err
	// }

	// r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return err

}
