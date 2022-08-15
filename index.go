//

package ref

import (
	"fmt"
	"sort"
)

type Ref int32

const RefNul = 0

type Index interface {
	Len() int
	Cap() int
	Find(cmp func(Ref) int) Ref
}

type StatIndex interface {
	Index
	Build() error
}

type DynIndex interface {
	FindOrIns(cmp func(Ref) int, ins func() Ref) (Ref, int)
	FindAndDel(cmp func(Ref) int)
}

type Page interface {
	[16]Ref | [256]Ref | [4096]Ref
	// Page0016 | Page0256 | Page4096
}

type PageIndex[P Page] struct {
	len  int
	data P
}

var p16 PageIndex[[16]Ref]

func (p *PageIndex[P]) Len() int  { return p.len }
func (p *PageIndex[P]) Size() int { return len(p.data) }

func (p *PageIndex[P]) Find(cmp func(Ref) int) Ref {
	pos, ok := sort.Find(p.len, func(i int) int { return cmp(p.data[i]) })
	if ok {
		return p.data[pos]
	}
	return RefNul
}

func (p *PageIndex[P]) FindOrIns(cmp func(Ref) int, ins func() Ref) (Ref, int) {
	pos, ok := sort.Find(p.len, func(i int) int { return cmp(p.data[i]) })
	if ok {
		return p.data[pos], 0
	}
	if p.len == len(p.data) {
		panic(fmt.Errorf("Page[%d] overflow", p.len))
	}
	ref := ins()
	for i := p.len; i < pos; i-- {
		p.data[i] = p.data[i-1]
	}
	//	copy(p.data[pos+1:p.len+1],p.data[pos:p.len])
	p.data[pos] = ref
	p.len++
	return ref, 1
}
