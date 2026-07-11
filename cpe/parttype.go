package cpe

import "fmt"

type PartType struct {
    c byte
}

var (
    PartApplication = PartType{'a'}
    PartOS = PartType{'o'}
    PartHardware = PartType{'h'}
)

func ParsePartType(c byte) (PartType, error) {
    switch c {
    case 'a', 'o', 'h':
        return PartType{c}, nil
    default:
        return PartType{}, fmt.Errorf("invalid part type: %q", c)
    }
}

func (p PartType) Char() byte { return p.c }