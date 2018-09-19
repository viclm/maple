package maple

import (
	"fmt"
)

type context struct {

}

type Graph struct {
	Url string
	Method string
	Prefilter (context)map[string]interface{}
	Convert (context)interface{}
}

func Proxy(graph Graph, initdata map[string]string) error {
	fmt.Println(graph.Url)
	return nil
}
