# please-nodejs

NodeJS integration w/ the Please build system.

This includes support for the following:

* `nodejs_toolchain`: Easy management of multiple versions of NodeJS.
* `nodejs_npm_package`: Incremental NPM packages from a NPM package registry.
* `nodejs_esbuild_bundle_browser`: Builds a browser bundle for the given srcs via ESBuild.
* `nodejs_esbuild_bundle_binary`: Builds a binary bundle for the given srcs via ESBuild and `vercel/pkg`.
* Advanced: `nodejs_esbuild_toolchain`: Easy management of multiple versions of ESBuild.
* Advanced: `nodejs_package_json`: Generates a `package.json` from `nodejs_npm_package`s.
* Advanced: `nodejs_node_modules`: Generates a `node_modules` from `nodejs_package_json`.
* Advanced: `nodejs_vercel_pkg_binaries`: Downloads Node binaries used by `vercel/pkg` to generate executables.

## `nodejs_toolchain`

This build rule allows you to specify a NodeJS version to download and re-use in most other `nodejs_*` rules.

## `nodejs_npm_package`

This build rule allows you to specify an NPM package to download. You must define all of the packages deps so that we can incrementally download packages. You can use our helper to generate the build rules into your repository, e.g.

```bash
# view help
$ ./pleasew run ///nodejs//third_party/binary:please_nodejs -- install --help

# generate rules for the latest version of express and its dependencies into the //third_party/nodejs Please package.
$ ./pleasew run ///nodejs//third_party/binary:please_nodejs -- install express

# generate rules for a specific version of express and its dependencies into the //third_party/nodejs Please package.
$ ./pleasew run ///nodejs//third_party/binary:please_nodejs -- install express@4.18.0

# generate rules in a structured directory layout under //third_party/nodejs/... for the latest version of express and its dependencies.
$ ./pleasew run ///nodejs//third_party/binary:please_nodejs -- install -s express

# generate rules in under a user-provided Please package for the latest version of express and its dependencies.
$ ./pleasew run ///nodejs//third_party/binary:please_nodejs -- install --pkg_prefix "//test/esbuild/react-dom/third_party/nodejs" express
```

I highly recommend creating an alias in your `.plzconfig` for this like so:
```ini
[alias "npm-install"]
desc = Runs please_nodejs install to install new dependencies into the repo
cmd = run ///nodejs//third_party/binary:please_nodejs -- -s
```


## `nodejs_esbuild_bundle_browser`

This build rule allows you to build a browser bundle via ESBuild. For example, a React web-application may be built via this rule. See `test/esbuild/react-typescript` for an example.

## `nodejs_esbuild_bundle_binary`

This build rule allows you to build a binary bundle via ESBuild and `vercel/pkg`. For example, a React server application may be built via this rule. See `test/esbuild/react-dom` for an example.

## Usage


### Please Plugin

```ini
; .plzconfig
; Support the non *-rules repo name format of Please plugins.
PluginRepo = ["https://github.com/{owner}/{plugin}/archive/{revision}.zip"]
[Plugin "nodejs"]
Target = //third_party/plugins:nodejs
ToolVersion = "v0.0.1" ; Skipping ToolVersion will build the Tool from source.

[Build]
; This links package.json and node_modules to the repo so that IDEs give us type
; hinting. Make sure to add `package.json` and `node_modules/` to your
; `.gitignore`.
LinkGeneratedSources = true
```

```bash
# .gitignore
# These are generated via Please.
package.json
node_modules/
```

```python
# //third_party/plugins/BUILD
plugin_repo(
    name = "nodejs",
    owner = "VJftw",
    plugin = "please-nodejs",
    revision = "v0.0.1",
)
```
