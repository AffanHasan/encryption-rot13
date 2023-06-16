package main

import ( 
	"fmt"
	"os"
)

func checkError(error error) {
	if error != nil {
		panic(error)
	}
}

var (
	rot13CharacterMapping = map[byte]byte{'A':'N', 'B':'O', 'C':'P','D':'Q','E':'R','F':'S','G':'T','H':'U','I':'V','J':'W','K':'X','L':'Y','M':'Z','N':'A','O':'B','P':'C','Q':'D','R':'E','S':'F','T':'G','U':'H','V':'I','W':'J','X':'K','Y':'L','Z':'M', 'a':'n', 'b':'o', 'c':'p', 'd':'q', 'e':'r','f':'s','g':'t','h':'u','i':'v','j':'w','k':'x','l':'y','m':'z','n':'a','o':'b','p':'c', 'q':'d','r':'e','s':'f','t':'g','u':'h','v':'i','w':'j','x':'k','y':'l','z':'m'}
)

func rotate13(buffer []byte) {
	for i, c := range buffer {
		if v, flag := rot13CharacterMapping[c]; flag {
			buffer[i] = v
		} else {
		}
	}
}

func main() {
	// fetch list of files in directory
	err := os.Chdir(os.Args[1])
	checkError(err)
	files, err := os.ReadDir(".")
	checkError(err)
	//TODO: Add file count here
	fmt.Printf("\nTotal %d files\n", len(files))
	// create status structs list
        for _, entry := range files {
		if !entry.IsDir() {
			encryptFile(entry.Name())
		}
	}
}

// Take care of error scenarios
func encryptFile(name string) {
	var err error = nil
	// open file
	filePath := fmt.Sprintf("./%s", name)
	file, err := os.Open(filePath)
	checkError(err)
        
	// create tmp
        tmpFilePath := fmt.Sprintf("./%s-rot13", filePath)
	tmpFile, err := os.Create(tmpFilePath)
	checkError(err)
	// read buffer push to channel
	for {
               buffer := make([]byte, 500)
	       currentReadBytesCount, _ := file.Read(buffer)
	       if currentReadBytesCount == 0 { // it means end of file
			break;
	       }
	       rotate13(buffer) // encrypt
               // receive buffer and write to file
               _, writeErr := tmpFile.Write(buffer)
	       checkError(writeErr)
	}
	tmpFile.Sync()
        defer os.Rename(tmpFilePath, filePath)
        defer tmpFile.Close()
        defer os.Remove(filePath)
        defer file.Close()
}
