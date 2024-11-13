package gocollect

import (
	"fmt"
	"net/http"
	"os"
	"runtime/coverage"
)

func init() {
	gocoverDir := os.Getenv("GOCOVERDIR")
	if gocoverDir == "" {
		return
	}

	gocollectPath := os.Getenv("GOCOLLECT_PATH")
	if gocollectPath == "" {
		gocollectPath = "/collect"
	}

	http.Handle(gocollectPath, Collect(gocoverDir))
}

func Collect(dir string) http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
		logErr(coverage.WriteMetaDir(dir))
		logErr(coverage.WriteCountersDir(dir))
	}
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
