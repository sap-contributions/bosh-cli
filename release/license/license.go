package license

import (
	crypto2 "github.com/cloudfoundry/bosh-utils/crypto"

	"github.com/cloudfoundry/bosh-cli/v7/crypto"
	"github.com/cloudfoundry/bosh-cli/v7/release/resource"
)

type License struct {
	resource resource.Resource
}

func NewLicense(resource resource.Resource) *License {
	return &License{resource: resource}
}

func (l License) Name() string        { return l.resource.Name() }
func (l License) Fingerprint() string { return l.resource.Fingerprint() }

func (l *License) ArchivePath() string   { return l.resource.ArchivePath() }
func (l *License) ArchiveDigest() string { return l.resource.ArchiveDigest() }

func (l *License) Build(dev, final resource.ArchiveIndex) error { return l.resource.Build(dev, final) }
func (l *License) Finalize(final resource.ArchiveIndex) error   { return l.resource.Finalize(final) }

func (l *License) RehashWithCalculator(calculator crypto.DigestCalculator, archiveFileReader crypto2.ArchiveDigestFilePathReader) (*License, error) {
	newLicenseResource, err := l.resource.RehashWithCalculator(calculator, archiveFileReader)
	return &License{newLicenseResource}, err
}
