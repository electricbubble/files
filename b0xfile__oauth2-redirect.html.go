// Code generaTed by fileb0x at "2019-07-04 12:13:14.479076 +0800 CST m=+0.563417945" from config file "b0x.yaml" DO NOT EDIT.
// modified(2019-07-04 12:08:17.068777257 +0800 CST)
// original path: swagger-ui/dist/oauth2-redirect.html

package swaggerFiles

import (
	_ "embed"
	"os"
)

// FileOauth2RedirectHTML is "/oauth2-redirect.html"
//go:embed oauth2-redirect.html
var FileOauth2RedirectHTML []byte

func init() {

	f, err := FS.OpenFile(CTX, "/oauth2-redirect.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileOauth2RedirectHTML)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
