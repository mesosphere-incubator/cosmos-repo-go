package cosmos

import (
	"testing"
)

/**
 * Test the repository parsing
 */
func TestRepo(t *testing.T) {
	const REPO_STUB = `{
	  "packages": [
	    {
	      "packagingVersion": "2.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.0"
	    },
	    {
	      "packagingVersion": "3.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.1"
	    },
	    {
	      "packagingVersion": "4.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.2"
	    },
	    {
	      "packagingVersion": "5.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.3"
	    }
	  ]
	}`

	repo, err := NewRepoFromString(REPO_STUB)
	if err != nil {
		t.Errorf("Error parsing stub: %s", err.Error())
		return
	}

	// Find a package
	pkg, err := repo.FindPackageVersion("Foo", "1.1")
	if err != nil {
		t.Errorf("Error locating package: %s", err.Error())
		return
	}
	if pkg == nil {
		t.Errorf("Encountered invalid package")
		return
	}

	// Verify that's the correct package
	if pkg.GetName() != "Foo" {
		t.Errorf("Found invalid package")
		return
	}
	if pkg.GetVersion() != "1.1" {
		t.Errorf("Found invalid package")
		return
	}

	// Get all packages
	pkgs, err := repo.FindAllPackageVersions("Foo")
	if err != nil {
		t.Errorf("Error locating package versions: %s", err.Error())
		return
	}
	if len(pkgs) != 4 {
		t.Errorf("Found wrong number of packages")
		return
	}

	if pkgs[0].GetVersion() != "1.0" {
		t.Errorf("Invalid package #0 version")
	}
	if pkgs[1].GetVersion() != "1.1" {
		t.Errorf("Invalid package #1 version")
	}
	if pkgs[2].GetVersion() != "1.2" {
		t.Errorf("Invalid package #2 version")
	}
	if pkgs[3].GetVersion() != "1.3" {
		t.Errorf("Invalid package #3 version")
	}

	// Find latest
	pkg, err = repo.FindLatestPackageVersion("Foo")
	if err != nil {
		t.Errorf("Error locating latest version: %s", err.Error())
		return
	}
	if pkg.GetVersion() != "1.3" {
		t.Errorf("Found wrong version: %s", pkg.GetVersion())
		return
	}

}

/**
 * Test against mesosphere repo
 */
func TestUniverse(t *testing.T) {
	repo, err := NewRepoFromURL("https://universe.mesosphere.com/repo")
	if err != nil {
		t.Errorf("Error parsing stub: %s", err.Error())
		return
	}

	// Find a package that we know is there
	pkg, err := repo.FindPackageVersion("jenkins", "3.2.4-2.60.2")
	if err != nil {
		t.Errorf("Error locating package: %s", err.Error())
		return
	}
	if pkg == nil {
		t.Errorf("Encountered invalid package")
		return
	}

	// Print package description
	if pkg.GetDescription() == "" {
		t.Errorf("Encountered empty description")
	}
}
