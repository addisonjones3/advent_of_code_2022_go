package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Shell struct {
	RootDir *Directory
	CurrDir *Directory
}

func InitializeShell() *Shell {
	rootDir := NewDirectory("/", nil)
	sh := &Shell{RootDir: rootDir, CurrDir: rootDir}
	return sh
}

func (sh *Shell) ParseCommand(commandLines []string) {
	lineFields := strings.Fields(commandLines[0])
	if lineFields[0] != "$" {
		fmt.Println("Not a valid command")
	}

	switch lineFields[1] {
	case "ls":
		sh.ls(commandLines[1:])
	case "cd":
		sh.cd(lineFields[len(lineFields)-1])
	}
}

func (sh *Shell) ls(dirMemberLines []string) {
	for _, l := range dirMemberLines {
		items := strings.Fields(l)
		fileSize, _ := strconv.Atoi(items[0])
		if items[0] == "dir" {
			dirName := items[1]
			newChildDir := NewDirectory(dirName, sh.CurrDir)
			sh.CurrDir.Directories[newChildDir.Name] = newChildDir
			continue
		}
		if fileSize > 0 {
			sh.CurrDir.addFile(items[1], fileSize)
			continue
		}

	}
}

func (sh *Shell) cd(s string) {
	cdSteps := make([]string, 0)
	if s != "/" {
		cdSteps = strings.Split(s, "/")
	} else {
		cdSteps = append(cdSteps, s)
	}
	for _, dir := range cdSteps {
		switch dir {
		case "/":
			sh.CurrDir = sh.RootDir
		case "..":
			if sh.CurrDir.ParentDirectory != nil {
				sh.CurrDir = sh.CurrDir.ParentDirectory
			}
		default:
			sh.CurrDir = sh.CurrDir.Directories[dir]
		}
	}
}

func (sh *Shell) GetDirsWithMaxSize(maxSize int) []*Directory {
	dirs := make([]*Directory, 0)
	currDirSize := sh.CurrDir.size()
	if currDirSize <= maxSize {
		dirs = append(dirs, sh.CurrDir)
	}
	for _, dir := range sh.CurrDir.Directories {
		sh.CurrDir = dir
		nextDirs := sh.GetDirsWithMaxSize(maxSize)
		dirs = append(dirs, nextDirs...)
	}
	return dirs
}

func InputParser(inputLines []string) [][]string {
	commandOutput := make([][]string, 0)

	currCommandGroup := make([]string, 0)
	for _, inputLine := range inputLines {
		if string(inputLine[0]) == "$" {
			if len(currCommandGroup) > 0 {
				commandOutput = append(commandOutput, currCommandGroup)
				currCommandGroup = []string{inputLine}
			} else {
				currCommandGroup = append(currCommandGroup, inputLine)
			}
		} else {
			currCommandGroup = append(currCommandGroup, inputLine)
		}
	}
	commandOutput = append(commandOutput, currCommandGroup)
	return commandOutput
}
