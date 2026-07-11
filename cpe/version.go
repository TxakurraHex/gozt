package cpe

import (
	"errors"
	"fmt"
)

type Version struct {
	raw string
}

func NewVersion(raw string) (Version, error) {
	if err := validateVersion(raw); err != nil {
		return Version{}, err
	}
	return Version{raw: raw}, nil
}

// Returns the original, unescaped string, used t ocompare against
// NVD's versionStart/EndIncluding/Excluding fields.
func (v Version) Raw() string { return v.raw }

func (v Version) String() string {
	return formatComponent(v.raw)
}

func validateVersion(s string) error {
	switch s {
	case "":
		return errors.New("version cannot be empty")
	case "-":
		return errors.New(`"-" is reserved for the NA logical value`)
	case "*":
		return errors.New(`"*" is reserved for the ANY logical value`)
	}
	for _, r := range s {
		if r < 0x20 || r > 0x7e {
			return fmt.Errorf("non-printable-ASCII character %q not allowed", r)
		}
		if r == ' ' {
			return errors.New("whitespace is not allowed; CPE requires underscores instead")
		}
	}
	return nil
}