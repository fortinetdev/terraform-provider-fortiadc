package fortiadc

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateConvIPMask2CIDR(oNewIP, oOldIP string) string {
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return ip + "/" + strconv.Itoa(prefixSize)
		}
	}
	return oOldIP
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

var FILTER_TYPE_STRING = 0
var FILTER_TYPE_LIST = 1

type s_filters struct {
	typ        int
	filter     string
	fiter_list []string
}

func parse_filter(filter string) []s_filters {
	filters := make([]s_filters, 0)
	if !strings.HasPrefix(filter, "filter=") {
		return filters
	}
	filter = strings.TrimPrefix(filter, "filter=")
	andSlice := strings.Split(filter, "&")
	for i := 0; i < len(andSlice); i++ {
		orSlice := strings.Split(andSlice[i], ",")
		var tmp_filter s_filters
		if len(orSlice) > 1 {
			tmp_filter.typ = FILTER_TYPE_LIST
			tmp_filter.fiter_list = make([]string, 0)
			for j := 0; j < len(orSlice); j++ {
				tmp_filter.fiter_list = append(tmp_filter.fiter_list, orSlice[j])
			}
		} else {
			tmp_filter.typ = FILTER_TYPE_STRING
			tmp_filter.filter = andSlice[i]
		}
		filters = append(filters, tmp_filter)
	}
	return filters
}

func _check_match(value, oper, patt string) bool {
	fmt.Printf("%v %v %v\n", value, oper, patt)
	switch oper {
	case "==":
		if value == patt {
			return true
		}
	case "=*":
		if strings.EqualFold(value, patt) {
			return true
		}
	case "!=":
		if value != patt {
			return true
		}
	case "!*":
		if !strings.EqualFold(value, patt) {
			return true
		}
	case "=@":
		if strings.Contains(
			strings.ToLower(value),
			strings.ToLower(patt)) {
			return true
		}
	case "!@":
		if !strings.Contains(
			strings.ToLower(value),
			strings.ToLower(patt)) {
			return true
		}
	case "<=":
		a, err1 := strconv.Atoi(value)
		b, err2 := strconv.Atoi(patt)
		if err1 == nil && err2 == nil {
			if a <= b {
				return true
			}
		}

	case "<":
		a, err1 := strconv.Atoi(value)
		b, err2 := strconv.Atoi(patt)
		if err1 == nil && err2 == nil {
			if a < b {
				return true
			}
		}
	case ">=":
		a, err1 := strconv.Atoi(value)
		b, err2 := strconv.Atoi(patt)
		if err1 == nil && err2 == nil {
			if a >= b {
				return true
			}
		}
	case ">":
		a, err1 := strconv.Atoi(value)
		b, err2 := strconv.Atoi(patt)
		if err1 == nil && err2 == nil {
			if a > b {
				return true
			}
		}
	default:
	}
	return false
}

func check_match(filter string, data map[string]interface{}) bool {
	reg := regexp.MustCompile(`([^=*!@><]+)([=*!@><]+)([^=*!@><]+)`)
	match := reg.FindStringSubmatch(filter)
	if match != nil {

		if _, ok := data[match[1]].(string); ok {
			ret := _check_match(data[match[1]].(string), match[2], match[3])
			if ret {
				return true
			}

		}

	}
	return false
}

func run_filter(filters []s_filters, data map[string]interface{}) bool {
	for i := 0; i < len(filters); i++ {
		if filters[i].typ == FILTER_TYPE_STRING {
			if !check_match(filters[i].filter, data) {
				return false
			}
		} else {
			flg := false
			for j := 0; j < len(filters[i].fiter_list); j++ {
				if check_match(filters[i].fiter_list[j], data) {
					flg = true
					break
				}
			}
			if !flg {
				return false
			}
		}
	}
	return true
}

func escapeFilter(filter string) string {
	var rstSb strings.Builder
	andSlice := strings.Split(filter, "&")

	for i := 0; i < len(andSlice); i++ {
		orSlice := strings.Split(andSlice[i], ",")
		if i > 0 {
			rstSb.WriteString("&")
		}
		rstSb.WriteString("filter=")
		for j := 0; j < len(orSlice); j++ {
			reg := regexp.MustCompile(`([^=*!@><]+)([=*!@><]+)([^=*!@><]+)`)
			match := reg.FindStringSubmatch(orSlice[j])
			if j > 0 {
				rstSb.WriteString(",")
			}
			if match != nil {
				argName := match[1]
				argName = strings.ReplaceAll(argName, "_", "-")
				argName = strings.ReplaceAll(argName, "fosid", "id")
				argName = strings.ReplaceAll(argName, ".", "\\.")
				argName = strings.ReplaceAll(argName, "\\", "\\\\")
				argValue := url.QueryEscape(match[3])
				rstSb.WriteString(argName)
				rstSb.WriteString(match[2])
				rstSb.WriteString(argValue)
			}
		}
	}
	return rstSb.String()
}

func sortStringwithNumber(v string) string {
	i := len(v) - 1
	for ; i >= 0; i-- {
		if '0' > v[i] || v[i] > '9' {
			break
		}
	}
	i++

	b64 := make([]byte, 64/8)
	s64 := v[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}

	return v[:i] + string(b64)
}

func dynamic_sort_subtable(result []map[string]interface{}, fieldname string, d *schema.ResourceData) {
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		vs := v.(string)
		if vs == "true" || vs == "natural" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return sortStringwithNumber(v1) < sortStringwithNumber(v2)
			})
		} else if vs == "alphabetical" {
			sort.Slice(result, func(i, j int) bool {
				v1 := fmt.Sprintf("%v", result[i][fieldname])
				v2 := fmt.Sprintf("%v", result[j][fieldname])

				return v1 < v2
			})
		}
	}
}

func fortiAPIPatch(t interface{}) bool {
	if t == nil {
		return false
	} else if _, ok := t.(string); ok {
		return true
	} else if _, ok := t.(float64); ok {
		return true
	} else if _, ok := t.([]interface{}); ok {
		return true
	}

	return false
}

func isImportTable() bool {
	itable := os.Getenv("FORTIADC_IMPORT_TABLE")
	if itable == "false" {
		return false
	}
	return true
}

func convintflist2i(v interface{}) interface{} {
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return v
		}
		return t[0]
	} else if t, ok := v.(string); ok {
		if t == "" {
			return 0
		} else if iVal, _ := strconv.Atoi(t); ok {
			return iVal
		}
	}
	return v
}

func convintflist2str(v interface{}) interface{} {
	res := ""
	if t, ok := v.([]interface{}); ok {
		if len(t) == 0 {
			return res
		}

		bFirst := true
		for _, v1 := range t {
			if t1, ok := v1.(float64); ok {
				if bFirst == true {
					res += strconv.Itoa(int(t1))
					bFirst = false
				} else {
					res += " "
					res += strconv.Itoa(int(t1))
				}
			}
		}
	}
	return res
}

func i2ss2arrFortiAPIUpgrade(v string, splitv string) bool {
	splitv = strings.ReplaceAll(splitv, "v", "")

	v1, err := version.NewVersion(v)
	if err != nil {
		return false
	}

	v2, err := version.NewVersion(splitv)
	if err != nil {
		return false
	}

	if v1.GreaterThanOrEqual(v2) {
		return true
	}

	return false
}

func intBetweenWithZero(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if (v >= min && v <= max) || (v == 0) {
			return warnings, errors
		}

		errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) or equal to 0, got %d", k, min, max, v))

		return warnings, errors
	}
}

func toCertFormat(v interface{}) interface{} {
	if t, ok := v.(string); ok {
		if t != "" && !strings.HasPrefix(t, "\"") {
			t = strings.TrimRight(t, "\n")
			return "\"" + t + "\""
		}
	}
	return v
}

// parsePkeyMkeyImportID parses the import ID of a resource and returns the pkey and mkey and new state id
func parsePkeyMkeyImportID(id string) (string, string, string, error) {
	idParts := strings.Split(id, ".")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return "", "", "", fmt.Errorf("unexpected format of import ID (%q), expected pkey.mkey", id)
	}

	pkey := idParts[0]
	mkey := idParts[1]
	newID := pkey + "_" + mkey

	return pkey, mkey, newID, nil
}
