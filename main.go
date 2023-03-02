package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LobarD/FileSystemSimulator/libAdd"
)

func main() {
	fs := libAdd.NewFileSystem() //file system
	fs.AddFolder("/home")
	fs.AddFolder("/home/userLobar")
	fs.AddFile("/home/userLobar", &libAdd.File{Name: "test.txt", Size: "1024"})
	fs.AddFolder("/home/userLobar/documents")
	fs.AddFile("/home/userLobar/documents", &libAdd.File{Name: "report.docx", Size: "2048"})
	fs.AddFolder("/home/someUser")
	fs.AddFile("/home/someUser", &libAdd.File{Name: "someFile.txt", Size: "1024"})
	fs.PrintContents(nil, "")

	for {
		fmt.Print("\nEnter command (addFolder, addFile, list, exit): ")
		reader := bufio.NewReader(os.Stdin)         //bufer
		command, _ := reader.ReadString('\n')       // console input
		command = strings.TrimSuffix(command, "\n") //delete last simbol

		switch command {
		case "addFolder":
			fmt.Print("Enter full folder path: ")
			folderName, _ := reader.ReadString('\n') //intput full folder path
			folderName = strings.TrimSuffix(folderName, "\n")
			fs.AddFolder(folderName)

		case "addFile":
			fmt.Print("Enter full folder path: ")
			folderName, _ := reader.ReadString('\n')
			folderName = strings.TrimSuffix(folderName, "\n")

			fmt.Print("Enter file name: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSuffix(fileName, "\n")

			fmt.Print("Enter file size: ")
			fileSize, _ := reader.ReadString('\n')
			fileSize = strings.TrimSuffix(fileSize, "\n")
			fs.AddFile(folderName, &libAdd.File{Name: fileName, Size: fileSize})

		case "list":
			fs.PrintContents(nil, "") //printing current state simulated file sysytem

		case "exit":
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}
