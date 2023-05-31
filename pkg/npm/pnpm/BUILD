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
        "///third_party/go/github.com_Masterminds_semver_v3//:v3",
        "///third_party/go/github.com_mitchellh_go-homedir//:go-homedir",
        "///third_party/go/github.com_rs_zerolog//log",
        "//pkg/npm",
    ],
)
