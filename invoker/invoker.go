package invoker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

type Invoker struct {
	Dir     string
	Scripts map[string]os.FileInfo `json:",omitempty"`
}

func New(abspath string) (*Invoker, error) {
	all, err := ioutil.ReadDir(abspath)
	if err != nil {
		return nil, err
	}

	files := make(map[string]os.FileInfo)
	for i := range all {
		if !all[i].IsDir() {
			continue
		}

		path := fmt.Sprintf("%s/%s", abspath, all[i].Name())
		suball, err := ioutil.ReadDir(path)
		if err != nil {
			continue
		}

		rgx := regexp.MustCompile(fmt.Sprintf(`^(%s)\.(\w+)`, all[i].Name()))
		for j := range suball {
			if suball[j].IsDir() {
				continue
			}
			res := rgx.FindStringSubmatch(suball[j].Name())
			if len(res) == 3 {
				files[res[1]] = suball[j]
				break
			}
		}
	}

	if len(files) == 0 {
		return nil, errors.New("new invoker: no scripts found in specified directory")
	}

	return &Invoker{
		Dir:     abspath,
		Scripts: files,
	}, nil
}

func (i *Invoker) Run(script string, args ...string) ([]byte, error) {
	fi, ok := i.Scripts[script]
	if !ok {
		return nil, fmt.Errorf("invoke error: invalid script name %s", script)
	}

	run := fmt.Sprintf("%s/%s/%s", i.Dir, script, fi.Name())

	out, err := exec.Command(run, args...).Output()
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, errors.New("invoke error: no result from script")
	}

	return out, nil
}
