package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	maxDepth = -1
)

var ignoreDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
}

func main() {
	targetPath := parseArgs()

	absPath, err := filepath.Abs(targetPath)
	if err != nil {
		panic(err)
	}

	info, err := os.Stat(absPath)
	if err != nil || !info.IsDir() {
		panic("geçerli bir dizin ver")
	}

	rootName := filepath.Base(absPath)

	tree, err := generateTree(absPath, "", 1)
	if err != nil {
		panic(err)
	}

	outputFile := filepath.Join(absPath, "tree.md")

	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString(rootName + "/\n")
	writer.WriteString(tree)

	writer.Flush()

	fmt.Println("tree.md oluşturuldu:", outputFile)
}

func parseArgs() string {
	path := "."

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--depth=") {
			fmt.Sscanf(strings.TrimPrefix(arg, "--depth="), "%d", &maxDepth)
			continue
		}

		if !strings.HasPrefix(arg, "--") {
			path = arg
		}
	}

	return path
}

func generateTree(path, prefix string, depth int) (string, error) {
	if maxDepth != -1 && depth > maxDepth {
		return "", nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var dirs, files, dotFiles []os.DirEntry

	for _, entry := range entries {
		name := entry.Name()

		if ignoreDirs[name] {
			continue
		}

		if entry.IsDir() {
			dirs = append(dirs, entry)
			continue
		}

		if strings.HasPrefix(name, ".") {
			dotFiles = append(dotFiles, entry)
		} else {
			files = append(files, entry)
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	sort.Slice(dotFiles, func(i, j int) bool {
		return dotFiles[i].Name() < dotFiles[j].Name()
	})

	var ordered []os.DirEntry

	if depth == 1 {
		// ROOT: klasörler → dosyalar → dot dosyalar
		ordered = append(ordered, dirs...)
		ordered = append(ordered, files...)
		ordered = append(ordered, dotFiles...)
	} else {
		// ALT DİZİNLER: eski davranış
		ordered = append(ordered, dirs...)
		ordered = append(ordered, files...)
		ordered = append(ordered, dotFiles...)
	}

	var result strings.Builder

	for i, entry := range ordered {
		isLast := i == len(ordered)-1

		connector := "├── "
		nextPrefix := prefix + "│   "

		if isLast {
			connector = "└── "
			nextPrefix = prefix + "    "
		}

		result.WriteString(prefix + connector + entry.Name())

		if entry.IsDir() {
			result.WriteString("/\n")
			sub, err := generateTree(
				filepath.Join(path, entry.Name()),
				nextPrefix,
				depth+1,
			)
			if err != nil {
				return "", err
			}
			result.WriteString(sub)
		} else {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func shouldIgnore(name string) bool {
	if ignoreDirs[name] {
		return true
	}

	return false
}
