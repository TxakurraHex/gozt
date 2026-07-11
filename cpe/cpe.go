package cpe

import (
	"fmt"
	"strings"
)

type CpeEntry struct {
    Part PartType
    Vendor string
    Product string
    Version Version
}

func (e CpeEntry) String() string {
    return fmt.Sprintf(
        "cpe:2.3:%c:%s:%s:%s:*:*:*:*:*:*:*", 
        e.Part.Char(),
        formatComponent(e.Vendor),
        formatComponent(e.Product),
        e.Version.String(),
    )
}

func isCpeSafe(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' || r == '.' || r == '-'
}

func formatComponent(raw string) string {
    if raw == "" {
        return "*"
    }

    var b strings.Builder
    for _, r := range raw {
        if isCpeSafe(r) {
            b.WriteRune(r)
        } else {
            b.WriteByte('\\')
            b.WriteRune(r)
        }
    }
    return b.String()
}