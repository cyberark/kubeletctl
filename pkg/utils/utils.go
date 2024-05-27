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
		if key == "kubernetes.io" {
			if nestedMap, ok := val.(map[string]interface{}); ok {
				valStr := stringifyMap(nestedMap)
				row[1] = valStr
			} else {
				row[1] = fmt.Sprintf("%v", val)
			}
		} else {
			row[1] = fmt.Sprintf("%v", val)
		}
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
func stringifyMap(m map[string]interface{}) string {
	var result string
	for key, val := range m {
		result += fmt.Sprintf("%s: %s\n", key, val)
	}
	return result
}
