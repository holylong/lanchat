package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	// "flag"
)

/*read file return line list*/
func ReadFile(xtPath string) []string {
	var dir_list []string
	f, err := os.Open(xtPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}
		//fmt.Println(line)
		dir_list = append(dir_list, line)
	}
	return dir_list
}

/*handle text line*/
func HandleLine(line, lineBegin, lineEnd string) string {
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	line = strings.TrimSpace(line)
	//fmt.Println(line)
	strRet := lineBegin + " " + line + " " + lineEnd
	strRet = strings.Replace(strRet, "\n", "", -1)
	return strRet
}

/*handle text line list*/
func HandleLineList(xtPathList []string, lineBegin, lineEnd string) {
	for i := 0; i < len(xtPathList); i++ {
		xtPathList[i] = HandleLine(xtPathList[i], lineBegin, lineEnd)
	}
}

/*write file to txt*/
func WriteFile(arrfiles []string, xtPath, strExt string, flagtype bool) {
	xtPath += "." + strExt
	f, err := os.Create(xtPath)
	if err != nil {
		fmt.Println("Create File Error")
	}
	f.Close()
	fwrite, ferr := os.OpenFile(xtPath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer fwrite.Close()
	if ferr != nil {
		fmt.Println(ferr.Error())
	} else {
		//fmt.Println(arrfiles[0])
		for i := 0; i < len(arrfiles); i++ {
			_, ferr = fwrite.Write([]byte(arrfiles[i]))
			if ferr != nil {
				fmt.Println("Error!")
			}
		}
	}
}

func MontainFile(xtPath, lineBegin, lineEnd string) []string {
	arrfiles := ReadFile(xtPath)
	HandleLineList(arrfiles, lineBegin, lineEnd)
	return arrfiles
	//WriteFile(arrfiles, xtPath, strExt, flagtype)
}

func HandleMultiFile(xtPath []string) {

}
