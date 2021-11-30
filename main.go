package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/openapi"
)

// Source: https://gist.github.com/owulveryck/4bd452cc3692d7016a54131ec89fa09a
func generateOpenAPI(defFile string, config *load.Config) ([]byte, error) {
	buildInstances := load.Instances([]string{defFile}, config)
	insts := cue.Build(buildInstances)
	b, err := openapi.Gen(insts[0], nil)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "   ")
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func main() {
	b, err := generateOpenAPI("dashboard.cue", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
