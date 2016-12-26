

package util


import (
	"errors"
	"fmt"
  "reflect"
	"encoding/json"
)

/*
 Convert to JSON String from []Row (Rocks DB)
*/
func GetJSONString(stub *shim.ChaincodeStub, row Row, dataObject interface{}) (string, error) {

  structType := reflect.TypeOf(dataObject)

  numField := structType.NumField()
  columnLegth := len(row.Columns)

  var jsonRespBuffer bytes.Buffer

  if (numField != columnLegth) {
    err := errors.New("Invalid Type. Can not convert")
    return "", err
  }


  for i := 0; i < numField; i++ {
    structField := structType.Field(i)
    jsonTag := structField.Tag.Get("json")

    if i == 0 {
      jsonRespBuffer.WriteString("{")
    }

    jsonRespBuffer.WriteString("\"" + jsonTag + "\"" + ":" + "\"" + row.Columns[i].GetString_() + "\"")

    if i != numField-1 {
      jsonRespBuffer.WriteString(",")
    } else {
      jsonRespBuffer.WriteString("}")
    }
  }

	return jsonRespBuffer.String(), nil
}
