package forgot

import (
	"github.com/funny/overall"
	"testing"
)

func TestGC(t *testing.T) {
	sumary := overall.GCSummary()
	t.Log(sumary.String())
}
