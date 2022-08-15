//

package ref

import (
	"fmt"
	"testing"
)

func TestIndex(t *testing.T) {
	fmt.Println("TestIndex:")

	fmt.Println(p16.FindOrIns(cmp, ins))
	fmt.Println(p16.FindOrIns(cmp, ins))
	fmt.Println(p16.FindOrIns(cmp, ins))
	fmt.Println(p16, p16.Len(), p16.Size())
}

var tmp Ref = 0

func ins() Ref      { tmp++; return tmp }
func cmp(r Ref) int { return 1 }
