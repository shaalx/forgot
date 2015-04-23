package forgot

import (
	"fmt"
	// "github.com/shaalx/sstruct"
	"io/ioutil"
)

func niil() {
	fmt.Println("...")
}

func ListDirs(dir string, suf string) {
	finfos, err := ioutil.ReadDir(dir)
	if nil != err {
		fmt.Println(err)
		return
	}
	for _, info := range finfos {
		if info.IsDir() {
			fmt.Println("---------------------------------/", info.Name())
			if ".git" == info.Name() {
				continue
			}
			ListDirs(suf+"/"+info.Name(), suf+"/"+info.Name())
		}
		fmt.Println(info.Name())
	}
}

/*
func TestPath(t *testing.T) {
	// os.Chdir("./pkg3/httplib")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cwd)
	dir := filepath.Dir(cwd)
	t.Log(dir)
	filepath.Walk(cwd, walk)
}

// down into the sub dirs
func walk(filepath string, info os.FileInfo, err error) error {
	if info.IsDir() {
		if ".git" == info.Name() {
			return nil
		}
		// fmt.Println(info.Name(), "---------------------")
	}
	// fmt.Println(filepath, info.Name())
	return nil
}*/
