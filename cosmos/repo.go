package cosmos

import (
	"fmt"
	"strings"
)

type CosmosRepositoryImpl struct {
	Packages []CosmosPackage
}

func (r *CosmosRepositoryImpl) FindAllPackageVersions(name string) ([]CosmosPackage, error) {
	var lst []CosmosPackage = nil
	for _, pkg := range r.Packages {
		if pkg.GetName() == name {
			lst = append(lst, pkg)
		}
	}
	return lst, nil
}

func (r *CosmosRepositoryImpl) FindPackageVersion(name string, version string) (CosmosPackage, error) {
	for _, pkg := range r.Packages {
		if pkg.GetName() == name && pkg.GetVersion() == version {
			return pkg, nil
		}
	}
	return nil, fmt.Errorf("Package not found")
}

/**
 * Parse the given repository buffer into a cosmos repository structure
 */
func parseRepo(raw map[string]interface{}) (CosmosRepository, error) {
	// Parse packages
	packages, ok := raw["packages"]
	if !ok {
		return nil, fmt.Errorf("Mal-formatted repository JSON: missing packages key")
	}
	packagesList, ok := packages.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Mal-formatted repository JSON: packages is not an array")
	}

	// Process packages and store them in the repository
	repo := &CosmosRepositoryImpl{}
	for idx, pkg := range packagesList {
		pkgMap, ok := pkg.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Invalid package entry #%d", idx)
		}

		pkgVersion, ok := pkgMap["packagingVersion"]
		if !ok {
			return nil, fmt.Errorf("Missing `packagingVersion` field on package #%d", idx)
		}

		pkgVersionString, ok := pkgVersion.(string)
		if !ok {
			return nil, fmt.Errorf("Field `packagingVersion` is not string in package #%d", idx)
		}

		// Parse package according to the version string
		if strings.HasPrefix(pkgVersionString, "5.") {
			packageInst, err := parseV50Package(pkgMap)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse package #%d as v5.x package: %s", idx, err.Error())
			}
			repo.Packages = append(repo.Packages, packageInst)

		} else if strings.HasPrefix(pkgVersionString, "4.") {
			packageInst, err := parseV40Package(pkgMap)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse package #%d as v4.x package: %s", idx, err.Error())
			}
			repo.Packages = append(repo.Packages, packageInst)

		} else if strings.HasPrefix(pkgVersionString, "3.") {
			packageInst, err := parseV30Package(pkgMap)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse package #%d as v3.x package: %s", idx, err.Error())
			}
			repo.Packages = append(repo.Packages, packageInst)

		} else if strings.HasPrefix(pkgVersionString, "2.") {
			packageInst, err := parseV20Package(pkgMap)
			if err != nil {
				return nil, fmt.Errorf("Unable to parse package #%d as v2.x package: %s", idx, err.Error())
			}
			repo.Packages = append(repo.Packages, packageInst)

		} else {
			return nil, fmt.Errorf("Invalid `packagingVersion` value '%s' on package #%d", pkgVersionString, idx)
		}
	}

	return repo, nil
}
