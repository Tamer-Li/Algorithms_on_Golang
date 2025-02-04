package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	buf := make([]byte, 10*1024*1024)
	in.Buffer(buf, 10*1024*1024)

	t := scanInt(in)

	for i := 0; i < t; i++ {
		n := scanInt(in)
		strJSON := scanJSON(in, n)
		folderJSON := parseJSON(strJSON)
		fmt.Fprintf(out, "%d\n", antiVirus(folderJSON))
	}
}

func scanInt(scan *bufio.Scanner) int {
	scan.Scan()
	num, _ := strconv.Atoi(scan.Text())
	return num
}

func scanJSON(scan *bufio.Scanner, counts int) string {
	var str strings.Builder
	str.Grow(counts * 100)

	for i := 0; i < counts; i++ {
		scan.Scan()
		str.WriteString(scan.Text())
	}

	return str.String()
}

func parseJSON(strJSON string) Folder {
	var parse Folder
	json.Unmarshal([]byte(strJSON), &parse)
	return parse
}

func antiVirus(folder Folder) int {
	var virus bool
	var countVirus int

	for _, file := range folder.Files {
		if strings.HasSuffix(file, ".hack") {
			countVirus += len(folder.Files)
			virus = true
			break
		}
	}

	if virus {
		for _, subFolder := range folder.Folders {
			countVirus += findVirus(subFolder)
		}
	} else {
		for _, subFolder := range folder.Folders {
			countVirus += antiVirus(subFolder)
		}
	}
	return countVirus
}

func findVirus(folder Folder) int {
	var count int

	count += len(folder.Files)

	for _, subFolder := range folder.Folders {
		count += findVirus(subFolder)
	}

	return count
}
