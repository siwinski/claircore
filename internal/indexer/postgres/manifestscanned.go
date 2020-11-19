package postgres

import (
	"context"
	"fmt"

	"github.com/quay/claircore"
	"github.com/quay/claircore/internal/indexer"
)

// ManifestScanned determines if a manifest has been scanned by ALL the provided
// scanners.
func (s *store) ManifestScanned(ctx context.Context, hash claircore.Digest, vs indexer.VersionedScanners) (bool, error) {
	const (
		selectScanned = `
		SELECT scanner_id
		FROM scanned_manifest
				 JOIN manifest ON scanned_manifest.manifest_id = manifest.id
		WHERE manifest.hash = $1;
		`
	)

	// get the ids of the scanners we are testing for.
	expectedIDs, err := s.selectScanners(ctx, vs)
	if err != nil {
		return false, err
	}

	// get a map of the found ids which have scanned this package
	var foundIDs = map[int64]struct{}{}
	rows, err := s.pool.Query(ctx, selectScanned, hash)
	if err != nil {
		return false, fmt.Errorf("store:manifestScanned failed to select scanner IDs for manifest: %v", err)
	}
	defer rows.Close()
	var t int64
	for rows.Next() {
		if err := rows.Scan(&t); err != nil {
			return false, fmt.Errorf("store:manifestScanned failed to select scanner IDs for manifest: %v", err)
		}
		foundIDs[t] = struct{}{}
	}
	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("store:manifestScanned failed to select scanner IDs for manifest: %v", err)
	}

	// compare the expectedIDs array with our foundIDs. if we get a lookup
	// miss we can say the manifest has not been scanned by all the layers provided
	for _, id := range expectedIDs {
		if _, ok := foundIDs[id]; !ok {
			return false, nil
		}
	}

	return true, nil
}
