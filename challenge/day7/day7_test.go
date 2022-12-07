package day7

import (
	"fmt"
	"testing"
)

var lsLines1 = []string{
	"$ ls",
	"dir a",
	"1000 b.txt",
	"1000 c.dat",
}
var cdLine1 = []string{"$ cd a"}
var lsLines2 = []string{
	"$ ls",
	"1000 e.txt",
	"dir d",
}
var cdLine2 = []string{"$ cd d"}
var lsLines3 = []string{
	"$ ls",
	"1000 z.txt",
}

var fullExCommand = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

// func TestAddFile(t *testing.T) {
// 	fname := "testfile.txt"
// 	dir := NewDirectory("testdir", nil)
// 	dir.addFile(fname, 123)
// 	fmt.Println(dir.Files)
// }

// func TestParseLS(t *testing.T) {
// 	sh := InitializeShell()
// 	sh.ParseCommand(lsLines1)
// 	fmt.Println(sh.CurrDir.FileNames())
// }

func TestParseCDRoot(t *testing.T) {
	line := []string{"$ cd /"}
	sh := InitializeShell()
	sh.ParseCommand(line)
	// sh.cd("a")
	fmt.Println(sh.CurrDir.Name)
}

// func TestParseCDNamed(t *testing.T) {
// 	line := []string{"$ cd a"}
// 	sh := InitializeShell()
// 	NewDirectory("a", sh.CurrDir)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line)
// 	// sh.cd("a")
// 	fmt.Println(sh.CurrDir.Name)
// }

// func TestParseCDPrev(t *testing.T) {
// 	line1 := []string{"$ cd a"}
// 	line2 := []string{"$ cd .."}
// 	line3 := []string{"$ cd .."}
// 	sh := InitializeShell()
// 	NewDirectory("a", sh.CurrDir)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line1)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line2)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line3)
// 	fmt.Println(sh.CurrDir.Name)
// }

// func TestParseCDMulti(t *testing.T) {
// 	line1 := []string{"$ cd a"}
// 	line2 := []string{"$ cd b"}
// 	line3 := []string{"$ cd ../../a"}
// 	sh := InitializeShell()
// 	aDir := NewDirectory("a", sh.CurrDir)
// 	NewDirectory("b", aDir)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line1)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line2)
// 	fmt.Println(sh.CurrDir.Name)
// 	sh.ParseCommand(line3)
// 	fmt.Println(sh.CurrDir.Name)
// }

// func TestDirSize(t *testing.T) {
// 	sh := InitializeShell()
// 	sh.ParseCommand(lsLines1)

// 	sh.ParseCommand(cdLine1)
// 	sh.ParseCommand(lsLines2)
// 	fmt.Println(sh.RootDir.size())
// 	fmt.Println(sh.CurrDir.size())
// }

func TestMaxSize(t *testing.T) {
	sh := InitializeShell()
	sh.ParseCommand(lsLines1)
	sh.ParseCommand(cdLine1)
	sh.ParseCommand(lsLines2)
	sh.ParseCommand(cdLine2)
	sh.ParseCommand(lsLines3)

	sh.cd("/")
	dirs := sh.GetDirsWithMaxSize(2000)
	for _, dir := range dirs {
		fmt.Println(dir.Name)
		fmt.Println(dir.size())
	}
}

// func TestInputParser(t *testing.T) {
// 	commandGroups := InputParser(fullExCommand)
// 	for _, cg := range commandGroups {
// 		fmt.Println(cg)
// 	}
// }

func TestExampleSize(t *testing.T) {
	commandGroups := InputParser(fullExCommand)
	// for _, cg := range commandGroups {
	// 	fmt.Println(cg)
	// }
	sh := InitializeShell()
	for _, cg := range commandGroups {
		fmt.Println(cg)
		sh.ParseCommand(cg)
	}

	fmt.Println(sh.RootDir.size())

	sh.cd("/")
	dirs := sh.GetDirsWithMaxSize(100000)
	fileSum := 0
	for _, dir := range dirs {
		fmt.Println(dir.Name)
		fmt.Println(dir.size())
		fileSum += dir.size()
	}
	fmt.Println(fileSum)

}
