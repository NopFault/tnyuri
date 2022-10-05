package tnyuri

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func modelName[M Model]() string {
	var modelName string = reflect.TypeOf(*new(M)).Name()
	return strings.ToLower(modelName)

}

func By[M Model](by string, val string) []M {
	return RowsBy[M](modelName[M](), by, val)
}

func Update[M Model](data map[string]string, by string, val string) {
	q := "UPDATE " + modelName[M]() + " SET "
	var d = []string{}
	for attr, val := range data {
		d = append(d, attr+"='"+val+"'")
	}
	q += strings.Join(d, ",") + " WHERE " + by + "='" + val + "'"
	fmt.Println(q)
	Exec(q)
}

func Delete[M Model](user string, id int) {
	Exec("DELETE FROM " + modelName[M]() + " WHERE user='" + user + "' AND id='" + strconv.Itoa(id) + "'")
}
