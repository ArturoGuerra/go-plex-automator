package revars

import "regexp"

var (
    LinuxPath *regexp.Regexp = regexp.MustCompile(`^(?:\/([^\/ ]+))+\/?$`)
    StrInt *regexp.Regexp = regexp.MustCompile(`^[\dA-Fa-f]+$`)
)
