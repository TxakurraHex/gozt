package source

import (
	"context"
	"errors"
	"fmt"
	"gozt/cpe"
	"os/exec"
	"strings"
	"time"
)

type DpkgSource struct {
	name string
}

func (s DpkgSource) Name() string { return s.name }
func (s DpkgSource) Collect() ([]cpe.CpeEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	out, err := exec.CommandContext(ctx, "dpkg-query", "-W", "-f=${Package} ${Version}\n").Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil, fmt.Errorf("dpkg-query failed: %s", exitErr.Stderr)
		}
		return nil, fmt.Errorf("dpkg-query failed: %w", err)
	}

	var entries []cpe.CpeEntry
	var errs []error

	for i, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		fields := strings.SplitN(line, " ", 2)
		if len(fields) != 2 {
			errs = append(errs, fmt.Errorf("line %d: malformed entry %q", i+1, line))
			continue
		}
		v, err := cpe.NewVersion(fields[1])
		if err != nil {
			errs = append(errs, fmt.Errorf("line %d (%s): %w", i+1, fields[0], err))
			continue
		}
		entries = append(entries, cpe.CpeEntry{
			Part: cpe.PartApplication,
			Product: fields[0],
			Version: v,
		})
	}

	return entries, errors.Join(errs...)
}