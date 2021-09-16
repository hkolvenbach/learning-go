package packageone

import "fmt"

var PackageVar = "PackageVar"

func PrintMe(myVar, blockVar string) {
	fmt.Println("Printing variables myVar, blockVar, PackageVar:", myVar, blockVar, PackageVar)
}
