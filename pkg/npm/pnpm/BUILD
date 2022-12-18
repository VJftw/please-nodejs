subinclude("///go//build_defs:go")

go_library(
    name = "pnpm",
    srcs = [
        "cache.go",
        "depsmanager.go",
    ],
    visibility = [
        "//cmd/...",
    ],
    deps = [
        "//pkg/npm",
        "//third_party/go/github.com/Masterminds/semver/v3",
        "//third_party/go/github.com/mitchellh/go-homedir",
        "//third_party/go/github.com/rs/zerolog",
    ],
)
