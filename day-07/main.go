package main

import (
	"fmt"
	"github.com/nielsvangijzen/aoc2022/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Parser struct {
	input   []string
	current string
	index   int

	Root       *Directory
	currentDir *Directory
}

func NewParser(input string) *Parser {
	lines := strings.Split(input, "\n")
	root := &Directory{
		Name:        "/",
		Files:       map[string]*File{},
		Directories: map[string]*Directory{},
	}

	// Because the first line is always cd / we just create the root dir and process
	// from then on!
	return &Parser{
		input:      lines[1:],
		Root:       root,
		currentDir: root,
		index:      -1,
	}
}

func (p *Parser) Next() bool {
	p.index += 1
	if p.index >= len(p.input) {
		return false
	}

	p.current = p.input[p.index]
	return true
}

func (p *Parser) Parse() (*Directory, error) {
	for p.Next() {
		line := p.current

		// We can determine the action that needs to be taken by the first 4 characters.
		// Technically we can do it with the first 3 but that wouldn't be as clear in the
		// code.
		prefix := line[:4]
		switch prefix {
		case "$ cd":
			p.runCD()
		case "$ ls":
			// The ls function will process the output of the ls command. This advances the
			// internal counter and lines until the next dollar sign is found.
			p.runLS()
		default:
			return nil, fmt.Errorf("parser got line: [%s], this is either "+
				"garbage data or should be handled by a subcommand", line)
		}
	}

	return p.Root, nil
}

func (p *Parser) runCD() {
	parts := strings.Split(p.current, " ")
	dir := parts[len(parts)-1]

	// We have 2 'edge' cases (/ and ..). Those have special meaning.
	switch dir {
	case "..":
		// We should go to the parent here, so we look at the address of the current
		// directory's parent. We change the current directory in the parser to that
		// address.
		if p.currentDir.parent == nil {
			return
		}
		p.currentDir = p.currentDir.parent
		return
	case "/":
		// A single slash means we want to CD to root. We could just traverse up the tree
		// until we hit a nil pointer indicating root. But keeping a pointer to the root
		// dir around is good enough for now this use case.
		p.currentDir = p.Root
	default:
		// To be clear, we assume that the directory exists because we would have made it
		// in the LS stage of the output. Normally we'd check if the directory exists
		// prior to entering it.
		p.currentDir = p.currentDir.Directories[dir]
	}
}

func (p *Parser) runLS() {
	for p.Next() {
		switch p.current[0] {
		case '$':
			// In LS, if the line starts with a dollar sign, it means we've hit another
			// command. LS doesn't process this, so we revert the counter by one and give
			// control back to the parser.
			p.index -= 1
			return
		case 'd':
			// We've found a new directory :D. We split the command on spaces and take the
			// last part. We don't have to worry about multiple dirs in the path as they
			// don't seem to exist in the input.
			parts := strings.Split(p.current, " ")
			p.currentDir.createDirectory(parts[1])
		default:
			// We've found a new file :D. We let the addFile command deal with
			// this and continue on!
			err := p.addFile()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (p *Parser) addFile() error {
	parts := strings.Split(p.current, " ")
	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	// We keep two copies around of the filename. One in the Files map and one in the
	// file itself. In a real scenario you might not want to do this because it uses
	// more RAM. But for now we go for convenience over efficiency.
	p.currentDir.Files[parts[1]] = &File{
		Name: parts[1],
		Size: size,
	}

	return nil
}

type Directory struct {
	Name        string
	Files       map[string]*File
	Directories map[string]*Directory

	parent *Directory
}

func (d *Directory) createDirectory(name string) *Directory {
	newDir := &Directory{
		Name:        name,
		Files:       map[string]*File{},
		Directories: map[string]*Directory{},
		parent:      d,
	}

	d.Directories[name] = newDir
	return newDir
}

// Size this is a simple function that walks the current directory recursively
// and counts up the sizes of its own files and the sizes of its child
// directories.
func (d *Directory) Size() (size int) {
	for _, file := range d.Files {
		size += file.Size
	}

	for _, directory := range d.Directories {
		size += directory.Size()
	}

	return
}

// FindDirectoryWithMaxSize is a function that finds all directories with the
// maximum size of maxSize and adds them to the sum of those directories (this is
// the magic for part one)
func FindDirectoryWithMaxSize(d *Directory, maxSize int) int {
	sum := 0

	if d.Size() <= maxSize {
		sum += d.Size()
	}

	for _, directory := range d.Directories {
		sum += FindDirectoryWithMaxSize(directory, maxSize)
	}

	return sum
}

// FindAllSizes walks the directory tree and calculates the file size for each
// directory it sees. The tree is turned into a flat slice here.
func FindAllSizes(d *Directory) []int {
	var sizes []int

	sizes = append(sizes, d.Size())

	for _, directory := range d.Directories {
		// We merge the two slices by using ... after the FindAllSlices
		sizes = append(sizes, FindAllSizes(directory)...)
	}

	return sizes
}

type File struct {
	Name string
	Size int
}

func partOne() {
	input := util.MustInputString("day-07/input.txt")
	parser := NewParser(input)

	dir, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("D07P01: %d\n", FindDirectoryWithMaxSize(dir, 100000))
}

func partTwo() {
	input := util.MustInputString("day-07/input.txt")
	parser := NewParser(input)

	// We parse the tree like normal
	dir, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Now we calculate the total size of the directories:
	totalSize := dir.Size()
	totalDiskSpace := 70000000
	freeSpace := totalDiskSpace - totalSize
	spaceRequired := 30000000 - freeSpace

	// Now we turn the tree into a flat structure by walking it recursively
	// We only care about the size and not the directory that matches it
	// giving us an easy way out in this case.
	allSizes := FindAllSizes(dir)

	// Now we sort it from smallest to largest.
	sort.Ints(allSizes)

	// We can now loop over the slice to find the very first size
	// that leaves us enough room for the update!
	for _, size := range allSizes {
		if size >= spaceRequired {
			fmt.Printf("D07P02: %d\n", size)
			break
		}
	}
}

func main() {
	partOne()
	partTwo()
}
