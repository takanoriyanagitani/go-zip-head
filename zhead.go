package main

import (
	"archive/zip"
	"errors"
	"log"
	"os"
	"strconv"
)

type Head uint32

func (h Head) ZipToWriter(wtr *zip.Writer) func(*zip.Reader) error {
	return func(rdr *zip.Reader) error {
		var files []*zip.File = rdr.File
		var isz int = len(files)
		var hsz int = int(h)
		var sz int = min(isz, hsz)
		var limited []*zip.File = files[:sz]
		for _, f := range limited {
			e := wtr.Copy(f)
			if nil != e {
				return e
			}
		}
		return wtr.Flush()
	}
}

func (h Head) ZipToStdout(rdr *zip.Reader) error {
	var zw *zip.Writer = zip.NewWriter(os.Stdout)
	e := h.ZipToWriter(zw)(rdr)
	return errors.Join(e, zw.Close())
}

func (h Head) ZipfileToStdout(filename string) error {
	f, e := os.Open(filename)
	if nil != e {
		return e
	}
	defer f.Close()

	s, e := f.Stat()
	if nil != e {
		return e
	}

	var sz int64 = s.Size()

	rdr, e := zip.NewReader(f, sz)
	if nil != e {
		return e
	}

	return h.ZipToStdout(rdr)
}

func headString() string {
	return os.Getenv("ENV_HEAD")
}

func head() uint32 {
	var s string = headString()
	i, e := strconv.Atoi(s)
	switch e {
	case nil:
		return uint32(i)
	default:
		return 10
	}
}

func zipfilename() string {
	return os.Getenv("ENV_INPUT_ZIP_FILENAME")
}

func sub() error {
	var hd Head = Head(head())
	var nm string = zipfilename()
	return hd.ZipfileToStdout(nm)
}

func main() {
	e := sub()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
