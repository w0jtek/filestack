package transformer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Create is responsible for initializing transformer based on a given command
func Create(command string) (transformer Transformer, err error) {

	if command == "exif" {
		transformer = NewExif()
		return
	}

	match, _ := regexp.MatchString(`crop=\d+,\d+,\d+,\d+`, command)
	if match {
		s := strings.Split(command[5:], ",")
		if len(s) != 4 {
			err = fmt.Errorf("Expected four integer parameters for cropping.")
			return
		}
		var cropParams []int
		for _, s := range s {
			param, atoiErr := strconv.Atoi(s)
			if atoiErr != nil {
				err = fmt.Errorf("Crop parameter is invalid")
			}
			cropParams = append(cropParams, param)
		}

		transformer = NewCrop(cropParams[0], cropParams[1], cropParams[2], cropParams[3])
		return
	}

	return nil, fmt.Errorf("Cannot understand command: %s", command)
}
