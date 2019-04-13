package mytestingreadingfile

//mytestingreadingfile is a package that testing for golang strings api and some file operations.
//The errorMessage.txt is the message from compiling the go-ipfs.

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//List is self-defined struct, and I want it similar to the List structure in Java
type List struct {
	name string
}

var mylist []List // an empty list

//Run is the main function for running the whole mytestingreadingfile package.
func Run() {
	ReadFile()
	fmt.Println("test main")

	str := "cannot find package \"github.com/whyrusleeping/tar-utils\" in any of:"
	myMap := findQuotationIndex(str)

	fmt.Println("myfunc", str[myMap[0]+1:myMap[1]])
	//印第21個到54  不包含55
	teststr := str[21:55]
	fmt.Println("teststr", teststr)

	os.Create("MyTestingReadFile/a.txt")

	//第一個是index 第二個是element
	for _, s := range mylist {
		RealMap := findQuotationIndex(s.name)
		fmt.Println(s.name[RealMap[0]+1 : RealMap[1]])
		appendToFile("MyTestingReadFile/a.txt", "go get "+s.name[RealMap[0]+1:RealMap[1]]+"\n")
	}
}

// 寫入檔案（覆蓋原檔案）
func writeToFile(fileName string, content string) {
	err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
	if err != nil {
		fmt.Printf("write err : %v\n", err)
	} else {
		fmt.Println("write success.")
	}
}

// 寫入檔案（追加於文字後）
func appendToFile(fileName string, content string) error {
	// open file only read
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
	}
	defer f.Close()
	return err
}

func findQuotationIndex(mystr string) map[int]int {
	m := make(map[int]int)
	previousIndex := 0
	str := mystr

	//fmt.Println("findQuotationMark...")
	for i := 0; i < 2; i++ {
		if strings.Contains(str, "\"") {
			stri := strings.Index(str, "\"")
			if i == 0 {
				m[i] = stri
				previousIndex = stri
			} else {
				// fmt.Println("len str:", len(str))
				// fmt.Println("len previousIndex", previousIndex)
				m[i] = stri + previousIndex + 1
			}
			// fmt.Println(stri)
			// fmt.Println(len(str))
			str = str[stri+1:]
			// fmt.Println("str", str)
		}

	}

	return m

}

//ReadFile is for reading errorMessage.txt
func ReadFile() {

	// 開檔
	inputFile, Error := os.Open("MyTestingReadFile/errorMessage.txt")

	fmt.Println("reading...")
	// 判斷是否開檔錯誤
	if Error != nil {
		fmt.Println("開檔錯誤!")
		//return
	}
	// 離開時自動執行關檔
	defer inputFile.Close()

	//
	inputReader := bufio.NewReader(inputFile)
	//^^^^^^^^^^^         ^^^^^^^   ^^^^^^^
	// 緩衝輸入物件        建立函數   來源:已開啟檔案

	// 用迴圈讀取檔案內容
	for {
		// 讀取字串直到遇到跳行符號
		//inputString, Error := inputReader.ReadString('\n')
		inputString, Error := inputReader.ReadString(':')
		// 若到檔尾時分發生  io.EOF 錯誤
		// 根據此錯誤 判斷是否離開
		if Error == io.EOF {
			fmt.Println("已讀取到檔尾!!")
			break
		}

		//fmt.Println(inputString)
		//fmt.Println(strings.Index(inputString, "\""))

		if strings.Index(inputString, "\"") != -1 {
			mylist = append(mylist, List{name: inputString})
		}
		//fmt.Println("length", len(channels))
	}
	fmt.Println("the list: ", mylist)
	//fmt.Println(reflect.TypeOf(mylist))
}
