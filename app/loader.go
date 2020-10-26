package app

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/markbates/pkger"
)

func Loader(dir string) (string, error) {
	var schemas []string
	if err := pkger.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		f, err := pkger.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		schemas = append(schemas, string(b))

		return nil
	}); err != nil {
		return "", err
	}
	return strings.Join(schemas, "\n"), nil
}
