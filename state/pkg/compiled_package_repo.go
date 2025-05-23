package pkg

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	biindex "github.com/cloudfoundry/bosh-cli/v7/index"
	birelpkg "github.com/cloudfoundry/bosh-cli/v7/release/pkg"
)

type CompiledPackageRecord struct {
	BlobID   string
	BlobSHA1 string
}

type CompiledPackageRepo interface {
	Save(birelpkg.Compilable, CompiledPackageRecord) error
	Find(birelpkg.Compilable) (CompiledPackageRecord, bool, error)
}

type compiledPackageRepo struct {
	index biindex.Index
}

func NewCompiledPackageRepo(index biindex.Index) CompiledPackageRepo {
	return &compiledPackageRepo{index: index}
}

func (cpr *compiledPackageRepo) Save(pkg birelpkg.Compilable, record CompiledPackageRecord) error {
	err := cpr.index.Save(cpr.pkgKey(pkg), record)

	if err != nil {
		return bosherr.WrapError(err, "Saving compiled package")
	}

	return nil
}

func (cpr *compiledPackageRepo) Find(pkg birelpkg.Compilable) (CompiledPackageRecord, bool, error) {
	var record CompiledPackageRecord

	err := cpr.index.Find(cpr.pkgKey(pkg), &record)
	if err != nil {
		if errors.Is(err, biindex.ErrNotFound) {
			return record, false, nil
		}

		return record, false, bosherr.WrapError(err, "Finding compiled package")
	}

	return record, true, nil
}

type packageToCompiledPackageKey struct {
	PackageName string
	// Fingerprint of a package captures the sorted names of its dependencies
	// (but not the dependencies' fingerprints)
	PackageFingerprint string
	DependencyKey      string
}

func (cpr compiledPackageRepo) pkgKey(pkg birelpkg.Compilable) packageToCompiledPackageKey {
	return packageToCompiledPackageKey{
		PackageName:        pkg.Name(),
		PackageFingerprint: pkg.Fingerprint(),
		DependencyKey:      cpr.convertToDependencyKey(ResolveDependencies(pkg)),
	}
}

func (cpr compiledPackageRepo) convertToDependencyKey(packages []birelpkg.Compilable) string {
	dependencyKeys := []string{}

	for _, pkg := range packages {
		dependencyKeys = append(dependencyKeys, fmt.Sprintf("%s:%s", pkg.Name(), pkg.Fingerprint()))
	}

	sort.Strings(dependencyKeys)

	return strings.Join(dependencyKeys, ",")
}
