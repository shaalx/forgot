package forgot

// import (
// 	"go/build"
// 	"os"
// 	"testing"
// )

// func TestBuild(t *testing.T) {
// 	pkg, err := build.Import("github.com/shaalx/forgot", "", 0)
// 	if nil != err {
// 		t.Error(err)
// 	}
// 	t.Logf("%#v\n", pkg)
// }

// func TestBuilds(t *testing.T) {
// 	// os.Chdir("./pkg3/httplib")
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(cwd)
// 	pkg, err := build.ImportDir(cwd, build.AllowBinary)
// 	if nil != err {
// 		t.Error(err)
// 	}
// 	t.Logf("%#v\n", pkg)
// }

// func TestListDirs(t *testing.T) {
// 	pwd, _ := os.Getwd()
// 	ListDirs(pwd, "./")
// 	// ListDirs("./")
// }
