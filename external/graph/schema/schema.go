package schema

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger"
)

var String string

func init() {
	String = Load("github.com/alextanhongpin/go-graphql-template:/external/graph/schema")
}

func Load(dir string) string {
	var schemas []string
	if err := pkger.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".gql" {
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
		panic(err)
	}
	return strings.Join(schemas, "\n")
}
