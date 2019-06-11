package cosmos

import (
	"fmt"
	"testing"
)

/**
 * Test the repository parsing
 */
func TestRepo(t *testing.T) {
	const REPO_STUB = `{
	  "packages": [
	    {
	      "packagingVersion": "4.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.0"
	    },
	    {
	      "packagingVersion": "4.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.1"
	    },
	    {
	      "packagingVersion": "4.0",
	      "name": "Foo",
	      "description": "Some package",
	      "version": "1.2"
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
	if len(pkgs) != 3 {
		t.Errorf("Find wrong number of packages")
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
	fmt.Println(pkg.GetDescription())
}
