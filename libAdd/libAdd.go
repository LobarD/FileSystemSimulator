package libAdd

import (
	"fmt"
	"strings"
)

type File struct {
	Name string
	Size string
}

type Folder struct {
	Name     string
	Contents []interface{}
}

func (f *Folder) AddFile(file *File) {
	f.Contents = append(f.Contents, file)
}

func (f *Folder) AddFolder(folder *Folder) {
	f.Contents = append(f.Contents, folder)
}

type FileSystem struct {
	root *Folder
}

func NewFileSystem() *FileSystem { // create root folder ("/")
	return &FileSystem{root: &Folder{Name: "/"}}
}

// AddFolder in simulated file system
func (fs *FileSystem) AddFolder(path string) {
	folders := strings.Split(strings.Trim(path, "/"), "/")
	currentFolder := fs.root
	for _, folderName := range folders {
		var found bool
		for _, item := range currentFolder.Contents {
			if folder, ok := item.(*Folder); ok && folder.Name == folderName {
				currentFolder = folder
				found = true
				break
			}
		}
		if !found {
			newFolder := &Folder{Name: folderName}
			currentFolder.AddFolder(newFolder)
			currentFolder = newFolder
		}
	}
}

// AddFile in simulated file system
func (fs *FileSystem) AddFile(path string, file *File) {
	folders := strings.Split(strings.Trim(path, "/"), "/")
	currentFolder := fs.root
	for _, folderName := range folders[:len(folders)-1] {
		var found bool
		for _, item := range currentFolder.Contents {
			if folder, ok := item.(*Folder); ok && folder.Name == folderName {
				currentFolder = folder
				found = true
				break
			}
		}
		if !found {
			newFolder := &Folder{Name: folderName}
			currentFolder.AddFolder(newFolder)
			currentFolder = newFolder
		}
	}
	currentFolder.AddFile(file)
}

// PrintContents of simulated file system
func (fs *FileSystem) PrintContents(folder *Folder, prefix string) {
	if folder == nil {
		folder = fs.root
	}
	fmt.Println(prefix + folder.Name + "/")
	for _, item := range folder.Contents {
		if folder, ok := item.(*Folder); ok { //if contents is folder
			fs.PrintContents(folder, prefix+"  ")
		} else { //if contents is file
			file := item.(*File)
			fmt.Println(prefix + "  " + file.Name + " (" + file.Size + " bytes)")
		}
	}
}
