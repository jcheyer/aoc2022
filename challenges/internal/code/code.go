package code

import (
	"fmt"
	"strconv"
	"strings"
)

type numVal int

type number struct {
	value             numVal
	moved             bool
	left, right       *number
	orgLeft, orgRight *number
}

func (nv numVal) String() string {
	return fmt.Sprintf("(%d)", nv)
}

type Circle struct {
	start *number
	end   *number
}

func New(data []string) *Circle {
	c := &Circle{}
	var prev *number
	for _, line := range data {
		v, _ := strconv.Atoi(line)
		n := &number{
			value: numVal(v),
			moved: false,
		}
		if prev == nil {
			c.start = n
		} else {
			n.left = prev
			n.orgLeft = prev
			prev.right = n
			prev.orgRight = n
		}
		prev = n
	}
	c.end = prev.right
	return c
}

func (c *Circle) String() string {
	var s strings.Builder
	s.WriteString("Circle:\n")
	s.WriteString(c.list())

	return s.String()
}

func (c *Circle) list() string {
	var s strings.Builder
	n := c.start
	for {
		s.WriteString("Org:\n" + n.StringOrg() + "\n")
		s.WriteString("Current:\n" + n.StringCurrent() + "\n")
		if n == c.end {
			break
		}
		n = n.right
		if n == nil {
			break
		}
	}
	return s.String()
}

func (c *Circle) move(n *number) {
	if n.value == 0 {
		return
	}
	if n == c.start {
		c.start = c.start.right
	}
	if n.value > 0 {

	} else {

	}

}

func (n *number) StringCurrent() string {
	var s strings.Builder
	//s.WriteString("-----\n")
	if n.left != nil {
		s.WriteString(n.left.value.String())
	} else {
		s.WriteString("nil")
	}
	s.WriteString("<-" + n.value.String() + "->")
	if n.right != nil {
		s.WriteString(n.right.value.String())
	} else {
		s.WriteString("nil")
	}
	s.WriteString("\n")
	return s.String()
}

func (n *number) StringOrg() string {
	var s strings.Builder
	//s.WriteString("-----\n")
	if n.orgLeft != nil {
		s.WriteString(n.orgLeft.value.String())
	} else {
		s.WriteString("nil")
	}
	s.WriteString("<-" + n.value.String() + "->")
	if n.orgRight != nil {
		s.WriteString(n.orgRight.value.String())
	} else {
		s.WriteString("nil")
	}
	s.WriteString("\n")
	return s.String()
}
