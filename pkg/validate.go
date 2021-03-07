package pkg

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/urfave/cli/v2"
)

// ValidateBoardSize validates the argument passed to the --size flag
func ValidateBoardSize(c *cli.Context) (BoardSize, error) {
	s := c.String("size")
	bs := BoardSize{}
	rr, _ := regexp.Compile("[^x]*")
	cr, _ := regexp.Compile("[^x]*$")

	match, err := regexp.MatchString("\\d{2}[x]\\d{2}", s)
	if err != nil {
		return bs, err
	}

	if (match == false) || (len(s) > 5) {
		return bs, errors.New("invalid board size flag")
	}

	rows, err := strconv.Atoi(rr.FindString(s))
	if err != nil {
		return bs, err
	}

	columns, err := strconv.Atoi(cr.FindString(s))
	if err != nil {
		return bs, err
	}

	bs.Rows = rows
	bs.Columns = columns

	return bs, nil
}
