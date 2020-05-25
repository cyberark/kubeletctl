package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func IsNotArgsEmpty(args []string) bool {
	return args != nil && len(args) > 0
}

func AreNamespacePodContainerAndPodUIDFlagsSet(namespace string, pod string, contianer string, uid string) bool {
	return AreNamespacePodAndContainerFlagsSet(namespace, pod, contianer) && uid != ""
}

func AreNamespacePodAndContainerFlagsSet(namespace string, pod string, contianer string) bool {
	return namespace != "" && pod != "" && contianer != ""
}

// Taken from kubetok
func getDecodedJwtToken(tokenString string) *jwt.MapClaims {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})

	return &claims
}

// Taken from kubetok
func PrintDecodedToken(tokenString string) {
	claims := getDecodedJwtToken(tokenString)

	twOuter := table.NewWriter()
	twOuter.AppendHeader(table.Row{"Key", "Value"})
	twOuter.SetAlignHeader([]text.Align{text.AlignCenter, text.AlignCenter})

	for key, val := range *claims {
		row := make(table.Row, 2)
		row[0] = key
		row[1] = val
		twOuter.AppendRow(row)
	}
	twOuter.SetAlign([]text.Align{text.AlignCenter, text.AlignCenter})
	//twOuter.SetStyle(table.StyleLight)
	twOuter.SetStyle(table.StyleRounded)
	twOuter.Style().Title.Align = text.AlignCenter
	twOuter.SetTitle("Decoded JWT token")
	twOuter.Style().Options.SeparateRows = true
	fmt.Println(twOuter.Render())
}
