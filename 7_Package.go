/*
The packageDemo package display how the doc and test works with
codes inside a package

Run this command to see doc via CMD
    go doc

Run this command to see via well-formatted HTML
    godoc -http=:/3034

*/
package packageDemo

import console "fmt" // alias for imported package

// Demo() will show a printed line and return with "0" which indicates success.
//
// a public function start with capital letter
func Demo() int {
	console.Println("This is a demo") // use alias to call the member function
	return 0
}

// DemoWithFailure() will show a printed line and return wth "-1" which indicates failure
func DemoWithFailure() int {
	console.Println("This is a failed demo") // use alias to call the member function
	return -1
}

// a private function, invisible to outsider
func instance() {
	// Intended blank
}
