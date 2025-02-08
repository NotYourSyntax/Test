package main

import "fmt"

func main() {
	asciiArt := `
      /\     /\
     {  ` + "`" + `---'  }
     {  O   O  }
     ~~>  V  <~~
      \ \|/ /
       ` + `-----'____
       /     \    \_  \
      {       }\  )  \ \
      |  \_/  ) /   )  |
       \__/  /(_/   /  |
         (__/      |__/
	`

	fmt.Println(asciiArt)
}
