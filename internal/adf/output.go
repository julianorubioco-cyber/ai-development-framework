package adf

import (
	"encoding/json"
	"fmt"
	"io"
)

func PrintJSON(writer io.Writer, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(writer, string(data))
	return err
}
