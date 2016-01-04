// "main" for standalone running file like this file
package main

import "fmt"

// main.main() will be called first when a go program is running
func main() {

    // package "fmt" support I/O with all UTF-8 character
    // Related info: http://dominik.honnef.co/posts/2012/04/dealing_with_encodings_in_go/
    fmt.Printf("Hello, world; 你好, 世界\n")
}
