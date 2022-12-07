package directory

import (
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	Size int
	Name string
}

type Directory struct {
	Subdirs map[string]*Directory
	Files   map[string]file
	Parent  *Directory
	Name    string
}

func New(name string) *Directory {
	return &Directory{
		Subdirs: make(map[string]*Directory, 0),
		Files:   make(map[string]file, 0),
		Name:    name,
		Parent:  nil,
	}
}
func NewFile(name string, size int) file {
	return file{
		Size: size,
		Name: name,
	}
}

func (d *Directory) AddDir(sd *Directory) {
	d.Subdirs[sd.Name] = sd
	sd.Parent = d
}

func (d *Directory) AddFile(f file) {
	d.Files[f.Name] = f
}

func (d *Directory) Size() int {
	size := 0
	for _, sd := range d.Subdirs {
		size += sd.Size()
	}
	for _, f := range d.Files {
		size += f.Size
	}

	return size
}

func ParseTerminal(output []string) *Directory {
	var root *Directory
	var wd *Directory

	for _, line := range output {
		//fmt.Printf("Parsing line: %s\nRoot: %+v\nWD: %+v\n", line, root, wd)
		lp := strings.Split(line, " ")
		switch lp[0] {
		case "$":
			switch lp[1] {
			case "cd":
				switch lp[2] {
				case "/":
					if root == nil {
						root = New("/")
					}
					wd = root
					continue
				case "..":
					wd = wd.Parent
				default:
					wd = wd.Subdirs[lp[2]]
				}
			case "ls":
			default:
				fmt.Println("dont know what to do: " + line)
			}
		case "dir":
			wd.AddDir(New(lp[1]))

		default:
			size, err := strconv.Atoi(lp[0])
			if err != nil {
				panic("Parse Error: " + line)
			}
			wd.AddFile(NewFile(lp[1], size))
		}
	}

	return root
}

func (d *Directory) SumSizeOfAllDirsSmallerThan(maxSize int) int {
	size := 0
	if d.Size() <= maxSize {
		size += d.Size()
	}
	//fmt.Printf("Dir %s %d size: %d\n", d.Name, d.Size(), size)
	for _, sd := range d.Subdirs {
		size += sd.SumSizeOfAllDirsSmallerThan(maxSize)
	}

	return size
}

func (d *Directory) DirsGreaterThan(minSize int) []*Directory {
	ret := make([]*Directory, 0)
	if d.Size() > minSize {
		ret = append(ret, d)
	}
	for _, sd := range d.Subdirs {
		ret = append(ret, sd.DirsGreaterThan(minSize)...)
	}
	return ret
}
