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

	// push AppEngine Go Instance config to stdin

	port := int32(8080)
	dc := "/"
	id := "uatoday-instance"
	config := Config{
		AppId:           []byte("test"),
		VersionId:       []byte("test"),
		ApplicationRoot: []byte("."),
		ApiPort:         &port,
		Datacenter:      &dc,
		InstanceId:      &id,
	}

	bytes, err := proto.Marshal(&config)
	if err != nil {
		panic(err)
	}
	output := base64.StdEncoding.EncodeToString(bytes)

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

	os.Stdin, err = os.Open(dir + "/appcfg.proto")
	if err != nil {
		panic(err)
	}
}
