package appenginetestinit

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

var (
	SavedHttpTransport = http.DefaultTransport
	SavedHttpClient    = http.DefaultClient
)

// Call this function from your test package init() function
func Use() {
	// nothing here
}

func init() {
	if strings.LastIndex(os.Args[0], "test") != len(os.Args[0])-len("test") {
		// not a test environment
		SavedHttpTransport = nil
		SavedHttpClient = nil
		return
	}

	// Task: push AppEngine Go Instance config to stdin

	// 1. create configuration instance
	port := int32(8080)
	dc := "/"
	id := "test-instance"
	authDomain := "test"
	config := Config{
		AppId:           []byte("test-app"),
		VersionId:       []byte("test"),
		ApplicationRoot: []byte("."),
		ApiPort:         &port,
		Datacenter:      &dc,
		InstanceId:      &id,
		AuthDomain:      &authDomain,
	}

	// 2. serialize configuration
	bytes, err := proto.Marshal(&config)
	if err != nil {
		panic(err)
	}
	output := base64.StdEncoding.EncodeToString(bytes)

	// 3. write configuration to a file
	dir := os.TempDir()
	fo, err := os.Create(dir + "/appcfg.proto")
	if err != nil {
		panic(err)
	}

	n := len(output)
	for n > 0 {
		cnt, err := fo.Write([]byte(output[len(output)-n : len(output)]))
		if err != nil {
			panic(err)
		}
		n -= cnt
	}

	if err = fo.Close(); err != nil {
		panic(err)
	}

	// 4. Point standard input to the configuration file
	os.Stdin, err = os.Open(dir + "/appcfg.proto")
	if err != nil {
		panic(err)
	}
}
