package second_local_package

import "go_playground/hello/local_package"

func PrintFromLocalPackage(text string) string {
	local_package.Print(text)

	return "Printed"
}