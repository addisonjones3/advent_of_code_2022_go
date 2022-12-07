package day7

import (
	"strings"
)

type File struct {
	Name      string
	Size      int
	Extension string
}

type Directory struct {
	Name            string
	Files           []*File
	ParentDirectory *Directory
	Directories     map[string]*Directory
}

func NewDirectory(dirName string, parentDir *Directory) *Directory {
	files := make([]*File, 0)
	dirs := make(map[string]*Directory)

	newDir := &Directory{Name: dirName, Files: files, ParentDirectory: parentDir, Directories: dirs}
	if parentDir != nil {
		parentDir.Directories[newDir.Name] = newDir
	}
	return newDir
}

func (d *Directory) addFile(fname string, size int) {
	f := &File{Name: fname, Size: size, Extension: ""}
	i := strings.Index(fname, ".")
	if i > 0 {
		f.Extension = fname[i:]
	}
	d.Files = append(d.Files, f)
}

func (d *Directory) size() int {
	// fmt.Println("On dir", d.Name)
	// fmt.Println(d.Files)
	sizeSum := 0
	if len(d.Files) > 0 {
		for _, f := range d.Files {
			sizeSum += f.Size
		}
	}

	if len(d.Directories) > 0 {
		for _, childDir := range d.Directories {
			sizeSum += childDir.size()
		}
	}

	return sizeSum
}

func (d *Directory) FileNames() []string {
	fnames := make([]string, 0)
	for _, f := range d.Files {
		fnames = append(fnames, f.Name)
	}
	return fnames
}
