package indexer

import (
	"net/http"
)

// Opts are options to instantiate a indexer
type Opts struct {
	Client        *http.Client
	ScannerConfig struct {
		Package, Dist, Repo map[string]func(interface{}) error
	}
	Store        Store
	LayerScanner *LayerScanner
	FetchArena   FetchArena
	Ecosystems   []*Ecosystem
	Vscnrs       VersionedScanners
}
