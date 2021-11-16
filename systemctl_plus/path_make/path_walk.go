package path_make

import (
	
	//"github.com/hewenyu/toolspackage/klog"
	"os"
	"path/filepath"
)

func FilePathWalk(fp string) ([]string, error) {
	
	var files = make([]string, 0)
	
	err := filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		
		dirPath, _ := filepath.Abs(path)
		files = append(files, dirPath)
		
		return nil
	})
	
	return files, err
}
