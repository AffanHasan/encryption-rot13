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
	filePath := os.Args[1]
	tmpFilePath :=fmt.Sprintf("%s-rot13", filePath)
	var err error = nil

	file, err := os.Open(filePath)
	
        checkError(err)
	tmpFile, err := os.Create(tmpFilePath)
	checkError(err)

	
	channel := make(chan []byte)

	go func () {
           var currentReadBytesCount int = 0
	   var bufferSize int = 500
	   for {
                readBuffer := make([]byte, bufferSize)
		currentReadBytesCount, err = file.Read(readBuffer)
		if err != nil { // it means end of file
			fmt.Println("Reached EOF")
			close(channel)
			break;
		}

		if len(readBuffer) < cap(readBuffer) {
			fmt.Println("slicing")
			rotate13(readBuffer[:currentReadBytesCount + 1])
		} else {
		        rotate13(readBuffer)
	        }
	        channel <- readBuffer
		
          }
       }()

       for c := range channel {
	       _, writeErr := tmpFile.Write(c)
	       checkError(writeErr)
       }
       tmpFile.Sync()
       
       defer os.Rename(tmpFilePath, filePath)
       defer tmpFile.Close()
       defer os.Remove(filePath)
       defer file.Close()
       fmt.Println("Encryption Completed")
}
