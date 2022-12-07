package day7

import (
	"bufio"
	"fmt"
	"os"
)

func day7PreBuild(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	shellInput := make([]string, 0)
	for s.Scan() {
		shellInput = append(shellInput, s.Text())
	}
	return shellInput
}

func partA(shellInput []string) int {
	sh := InitializeShell()
	commandGroups := InputParser(shellInput)
	for _, cg := range commandGroups {
		sh.ParseCommand(cg)
	}
	sh.cd("/")
	qualDirs := sh.GetDirsWithMaxSize(100000)
	fmt.Println(len(qualDirs))
	dirSum := 0
	for _, dir := range qualDirs {
		dirSum += dir.size()
	}
	return dirSum
}

func partB(shellInput []string) int {
	sh := InitializeShell()
	commandGroups := InputParser(shellInput)
	for _, cg := range commandGroups {
		sh.ParseCommand(cg)
	}
	sh.cd("/")
	diskAvailable := 70000000
	diskRequired := 30000000
	diskUsed := sh.RootDir.size()
	diskFree := diskAvailable - diskUsed
	diskToFree := diskRequired - diskFree
	qualDirs := sh.GetDirsWithMaxSize(diskAvailable)
	minDirSize := diskRequired
	for _, dir := range qualDirs {
		dirSize := dir.size()
		if dirSize >= diskToFree && minDirSize >= dirSize {
			minDirSize = dir.size()
		}
	}
	return minDirSize
}
