# cosmos-repo-go
> A go client for parsing and querying a cosmos database

# Usage

```go
package main

import (
    "log"
    "github.com/mesosphere-incubator/cosmos-repo-go"
)

func main() {

    // Create a new repo from URL
    repo, err := cosmos.NewRepoFromURL("https://universe.mesosphere.com/repo")
    if err != nil {
        log.Fatal("Unable to create a repository: %s", err.Error())
    }

    // Find a package that we know is there
    pkg, err := repo.FindPackageVersion("jenkins", "3.2.4-2.60.2")
    if err != nil {
        log.Fatal("Unable to query repository: %s", err.Error())
    }
    if pkg == nil {
        log.Fatal("Unable to find jenkins: %s", err.Error())
    }
}
```
