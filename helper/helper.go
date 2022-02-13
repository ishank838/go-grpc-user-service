package helper

import "fmt"

func GetFormattedHeight(f float32) *string {
	res := fmt.Sprintf("%.1f", f)
	return &res
}
