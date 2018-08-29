/*

Package goidoit is an https://www.i-doit.com/ API client implementation written in https://golang.org

Install using go get

	go get github.com/cseeger-epages/i-doit-go-api

Initialize your API object

	a, err := goidoit.NewLogin("https://example.com/src/jsonrpc.php", "yourapikey", "username", "password")
	if err != nil {
		log.Fatal(err)
	}

and do stuff like get your idoit version or search stuff etc.

	v, err := a.Version()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v.Result[0]["version"])

	s, err := a.Search("test")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range s.Result {
		fmt.Println(v)
	}

for Debuging use

	goidoit.Debug(true)

and to disable TLS Certificate Verification (if you use self signed certs for i-doit)

	goidoit.SkipTLSVerify(true)

*/
package goidoit
