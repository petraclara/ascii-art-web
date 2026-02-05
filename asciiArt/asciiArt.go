package asciiArt

import(
	"strings"
	"os"

)

func AsciiArt(input string,banner string) (string, error){
	//read banner file

	bannerPath := "banners/" + banner

	
	data , err := os.ReadFile(bannerPath)
	if err != nil{
		return "",err
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")

var result strings.Builder
	//split banner file by newline
	lines := strings.Split(string(data), "\n")

	//split input file by newline
	Inputlines := strings.Split(input, "\\n")
for _, textline := range Inputlines{
		if textline == ""{
			result.WriteString("")
			continue
		}
		for row := 0; row < 8; row++{
			for _, char := range textline{
				asciiIndex := int(char)-32
				bannerLine := asciiIndex*9 + row
				result.WriteString(lines[bannerLine])
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}