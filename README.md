# FileSystemSimulator

➜  FileSystemSimulator git:(main) ✗ go run main.go
//
  home/
    UserLobar/
      documents/
      report.docx (2048 bytes)
    test.txt (1024 bytes)
    someUser/
    notes.txt (512 bytes)
    someFile.txt (1024 bytes)

Enter command (addFolder, addFile, list, exit): addFolder
Enter full folder path: /home/UserLobar/images

Enter command (addFolder, addFile, list, exit): addFile
Enter full folder path: /home/UserLobar/images
Enter file name: emptyFile.jpg
Enter file size: 0

Enter command (addFolder, addFile, list, exit): list
//
  home/
    UserLobar/
      documents/
      report.docx (2048 bytes)
      images/
      emptyFile.jpg (0 bytes)
    test.txt (1024 bytes)
    someUser/
    notes.txt (512 bytes)
    someFile.txt (1024 bytes)

Enter command (addFolder, addFile, list, exit): exit