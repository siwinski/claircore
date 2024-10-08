package indexer

import (
	"context"

	"github.com/quay/claircore"
)

// PackageScanner provides an interface for unique identification or a PackageScanner
// and a Scan method for extracting installed packages from an individual container layer
type PackageScanner interface {
	VersionedScanner
	// Scan performs a package scan on the given layer and returns all
	// the found packages
	Scan(context.Context, *claircore.Layer) ([]*claircore.Package, error)
}

type mockPackageScanner struct {
	name    string
	version string
	kind    string
}

func (p *mockPackageScanner) Scan(context.Context, *claircore.Layer) ([]*claircore.Package, error) {
	return nil, nil
}

func (p *mockPackageScanner) Name() string {
	return p.name
}

func (p *mockPackageScanner) Version() string {
	return p.version
}

func (p *mockPackageScanner) Kind() string {
	return p.kind
}

func NewPackageScannerMock(name, version, kind string) PackageScanner {
	return &mockPackageScanner{
		name:    name,
		version: version,
		kind:    kind,
	}
}

// DefaultRepoScanner provides a DefaultRepository method to allow Package Scanners
// to define what the default repository should be if the Scanner returned packages.
// This is useful to avoid having a dedicated RepositoryScanner for an ecosystem that
// largely duplicates the work of the PackageScanner and would always return the same
// pre-defined repository.
type DefaultRepoScanner interface {
	DefaultRepository(context.Context) *claircore.Repository
}
