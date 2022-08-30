# pkg

Go packages by Fox.ONE

# Upgrate Notes

## Separate go mod

Before this change, we put a single `go.mod` file in the root directory, any project that import the library will have a lot of unnecessary dependencies.

Then we delete the global `go.mod` file and put `go.mod` in each sub-package directory, based on this, there are some points that need attention here,

1. Migrate from old way, remove the `github.com/fox-one/pkg` line in your `go.mod` file, then execute `go mod tidy` will load all sub-package you imported.
2. No more tags like `vX.Y.Z` will be created later, according to the [go mod doc](https://go.dev/ref/mod#vcs-version), use `sub-package/vX.Y.Z` instead.
3. When you add a new sub-package, make sure that your sub-package does not depend on `github.com/fox-one/pkg`, otherwise it will lead to dependency conflict for the user.
