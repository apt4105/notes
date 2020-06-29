package fs_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/apt4105/notes/utils"
	"github.com/apt4105/notes/blob"
	"github.com/apt4105/notes/blob/fs"
)

func TestFSGetSetDel(t *testing.T) {
	const (
		opGet = iota
		opSet
		opDel
	)

	type op struct {
		op         int
		name, data string
		err        error
	}

	// one test case is a series of oprations on the file store
	type tcase struct {
		ops []op
	}

	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			// open up a temporary directory to write test files to
			dir, err := ioutil.TempDir("", "fs_test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(dir)

			store := &fs.FS{
				Basepath: dir,
			}

			for _, v := range tc.ops {
				switch v.op {
				case opSet:
					err := store.Set(v.name, strings.NewReader(v.data))
					if !utils.ErrEq(err, v.err) {
						t.Fatalf("unexpected err:\n\t%v\nexpected:\n\t%v", err, v.err)
					}

				case opGet:
					file, err := store.Get(v.name)
					if !utils.ErrEq(err, v.err) {
						t.Fatalf("unexpected err:\n\t%v\nexpected:\n\t%v", err, v.err)
					}
					if err != nil {
						continue
					}

					defer file.Close()

					byt, err := ioutil.ReadAll(file)
					if err != nil {
						t.Fatalf("unexpected err:\n\t%v", err)
					}

					if string(byt) != v.data {
						t.Fatalf("unexpected value:\n\t%v\nexpected\n\t%v",
							string(byt),
							v.data)
					}

				case opDel:
					// TODO
				}
			}

		}
	}

	tcases := map[string]tcase{
		"not found": tcase{
			ops: []op{
				{
					op:   opGet,
					name: "evil-plans.txt",
					data: "1. socks and flip flops",
					err:  blob.ErrNotFound,
				},
			},
		},
		"get and set": tcase{
			ops: []op{
				{op: opSet, name: "evil-plans.txt", data: "1. socks and flip flops"},
				{op: opGet, name: "evil-plans.txt", data: "1. socks and flip flops"},
			},
		},
	}

	for k, v := range tcases {
		t.Run(k, fn(v))
	}
}
