//go:build tools

// Package site exists only so that `go mod tidy` can see a use of the landing
// kit.
//
// ⛔ Do not delete this file. A Hugo site has no Go code, so tidy finds nothing
// importing github.com/go-fleettools/landingkit and DELETES the requirement.
// The fleet's Renovate presets set postUpdateOptions: ["gomodTidy"], and the
// grouped dependency branch is auto-merged — so every bump of the kit would
// arrive as a pull request that silently removes it, and the landing would
// quietly fall back to having no shared design at all.
//
// The build tag keeps this file out of every build. It exists only to be seen
// by tidy.
package site

import _ "github.com/go-fleettools/landingkit"
