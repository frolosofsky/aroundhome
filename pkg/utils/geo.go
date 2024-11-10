package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/frolosofsky/aroundhome/pkg/model"
)

// parses string "latitude;longitude" into position
func ParsePosition(s string) (res model.Position, err error) {
	parts := strings.Split(s, ";")
	if len(parts) != 2 {
		return res, errors.New("must contain two ;-separated parts")
	}

	res.Latitude, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}

	res.Longitude, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return
	}

	return
}

func PositionToString(p model.Position) string {
	return fmt.Sprintf("%.5f;%.5f", p.Latitude, p.Longitude)
}
