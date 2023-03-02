# FileSystemSimulator
To run program open bash console go to directory ../FileSystemSimularot
and type: go run main.go

$ go run main.go
➜  FileSystemSimulator git:(main) ✗ go run main.go
//
  home/
    userLobar/
      documents/
      report.docx (2048 bytes)
    test.txt (1024 bytes)
    someUser/
    someFile.txt (1024 bytes)

Enter command (addFolder, addFile, list, exit): addFolder
Enter full folder path: home/userLobar/documents/pdf's

Enter command (addFolder, addFile, list, exit): addFile
Enter full folder path: home/userLobar/documents/pdf's
Enter file name: OS.pdf
Enter file size: 4096

Enter command (addFolder, addFile, list, exit): list
//
  home/
    userLobar/
      documents/
        pdf's/
        OS.pdf (4096 bytes)
      report.docx (2048 bytes)
    test.txt (1024 bytes)
    someUser/
    someFile.txt (1024 bytes)

Enter command (addFolder, addFile, list, exit): exite
Invalid command

Enter command (addFolder, addFile, list, exit): exit