package main

import (
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Folder struct {
	name     string
	contents []interface{}
}

func (f *Folder) addFile(file *File) {
	f.contents = append(f.contents, file)
}

func (f *Folder) addFolder(folder *Folder) {
	f.contents = append(f.contents, folder)
}

type FileSystem struct {
	root *Folder
}

func NewFileSystem() *FileSystem {
	return &FileSystem{root: &Folder{name: "/"}}
}

func (fs *FileSystem) addFolder(path string) {
	folders := strings.Split(strings.Trim(path, "/"), "/")
	currentFolder := fs.root
	for _, folderName := range folders {
		var found bool
		for _, item := range currentFolder.contents {
			if folder, ok := item.(*Folder); ok && folder.name == folderName {
				currentFolder = folder
				found = true
				break
			}
		}
		if !found {
			newFolder := &Folder{name: folderName}
			currentFolder.addFolder(newFolder)
			currentFolder = newFolder
		}
	}
}

func (fs *FileSystem) addFile(path string, file *File) {
	folders := strings.Split(strings.Trim(path, "/"), "/")
	currentFolder := fs.root
	for _, folderName := range folders[:len(folders)-1] {
		var found bool
		for _, item := range currentFolder.contents {
			if folder, ok := item.(*Folder); ok && folder.name == folderName {
				currentFolder = folder
				found = true
				break
			}
		}
		if !found {
			newFolder := &Folder{name: folderName}
			currentFolder.addFolder(newFolder)
			currentFolder = newFolder
		}
	}
	currentFolder.addFile(file)
}

func (fs *FileSystem) printContents(folder *Folder, prefix string) {
	if folder == nil {
		folder = fs.root
	}
	fmt.Println(prefix + folder.name + "/")
	for _, item := range folder.contents {
		if folder, ok := item.(*Folder); ok {
			fs.printContents(folder, prefix+"  ")
		} else {
			file := item.(*File)
			fmt.Println(prefix + "  " + file.name + " (" + strconv.Itoa(file.size) + " bytes)")
		}
	}
}

func main() {
	fs := NewFileSystem()
	fs.addFolder("/home")
	fs.addFolder("/home/UserLobar")
	fs.addFile("/home/UserLobar", &File{name: "test.txt", size: 1024})
	fs.addFolder("/home/UserLobar/documents")
	fs.addFile("/home/UserLobar/documents", &File{name: "report.docx", size: 2048})
	fs.addFolder("/home/someUser")
	fs.addFile("/home/someUser", &File{name: "notes.txt", size: 512})
	fs.addFile("/home/someUser", &File{name: "someFile.txt", size: 1024})
	fs.printContents(nil, "")
}
