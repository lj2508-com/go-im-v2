package ctrl

import (
	"fmt"
	"go-im-v2/util"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}
func Upload(w http.ResponseWriter, req *http.Request) {
	UploadLocal(w, req)
}

func UploadLocal(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("file")
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	suffix := ".png"
	filename := header.Filename
	tmp := strings.Split(filename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	filetype := req.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	filename = fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	tempfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	_, err = io.Copy(tempfile, file)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	url := "/mnt/" + filename
	util.RespOk(w, url, "")
}
