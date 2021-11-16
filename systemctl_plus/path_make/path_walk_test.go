package path_make

import (
	"fmt"
	"log"
	"path"
	"testing"
)

func TestLoadPath(t *testing.T) {
	
	files, _ := FilePathWalk("./")
	
	for _, v := range files {
		
		if path.Ext(v) == ".log" {
			log.Println(fmt.Sprintf("add file : %v", v))
		}
		
	}
}
