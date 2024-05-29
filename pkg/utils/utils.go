package utils

import (
	"fmt"

	"strings"

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
		if key == "kubernetes.io" || key == "aud" {
			if nestedMap, ok := val.(map[string]interface{}); ok {
				valStr := stringifyMap(nestedMap, 0)
				row[1] = valStr
			} else if lineSlice, ok := val.([]interface{}); ok {
				lineStr := ""
				for _, aud := range lineSlice {
					lineStr += fmt.Sprintf("%v\n", aud)
				}
				row[1] = lineStr
			} else {
				row[1] = fmt.Sprintf("%v", val)
			}
		} else {
			row[1] = fmt.Sprintf("%v", val)
		}
		twOuter.AppendRow(row)
	}
	//twOuter.SetAlign([]text.Align{text.AlignCenter, text.AlignCenter})
	twOuter.SetAlign([]text.Align{text.AlignLeft, text.AlignLeft})
	//twOuter.SetStyle(table.StyleLight)
	twOuter.SetStyle(table.StyleRounded)
	twOuter.Style().Title.Align = text.AlignCenter
	twOuter.SetTitle("Decoded JWT token")
	twOuter.Style().Options.SeparateRows = true
	fmt.Println(twOuter.Render())

}
func stringifyMap(m map[string]interface{}, indent int) string {
	var result strings.Builder
	indentation := strings.Repeat("  ", indent)
	for key, val := range m {
		switch v := val.(type) {
		case map[string]interface{}:
			result.WriteString(fmt.Sprintf("%s%s:\n%s", indentation, key, stringifyMap(v, indent+1)))
		default:
			result.WriteString(fmt.Sprintf("%s%s: %v\n", indentation, key, v))
		}
	}
	return result.String()
}
