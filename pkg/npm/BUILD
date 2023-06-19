subinclude("///go//build_defs:go")

go_library(
    name = "npm",
    srcs = [
        "client.go",
        "packagejson.go",
        "packagelockjsonpackage.go",
        "packagemetadata.go",
        "packageversiondata.go",
    ],
    visibility = [
        "//cmd/...",
        "//pkg/...",
    ],
    deps = [
        "///third_party/go/github.com_Masterminds_semver_v3//:v3",
    ],
)
