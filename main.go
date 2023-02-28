package main

import (
	"bufio"
	"fmt"
	"os"
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

	for {
		fmt.Print("\nEnter command (addFolder, addFile, list, exit): ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		switch command {
		case "addFolder":
			fmt.Print("Enter full folder path: ")
			folderName, _ := reader.ReadString('\n')
			folderName = strings.TrimSuffix(folderName, "\n")
			fs.addFolder(folderName)

		case "addFile":
			fmt.Print("Enter full folder path: ")
			folderName, _ := reader.ReadString('\n')
			folderName = strings.TrimSuffix(folderName, "\n")

			fmt.Print("Enter file name: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSuffix(fileName, "\n")

			fmt.Print("Enter file size: ")
			fileSize, _ := reader.ReadString('\n')
			fileSizeInt, _ := strconv.Atoi(fileSize)
			fs.addFile(folderName, &File{name: fileName, size: fileSizeInt})

		case "list":
			fs.printContents(nil, "")

		case "exit":
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}
