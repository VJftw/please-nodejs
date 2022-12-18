subinclude("///go//build_defs:go")

go_library(
    name = "npm",
    srcs = [
        "client.go",
        "packagejson.go",
        "packagemetadata.go",
        "packageversiondata.go",
    ],
    visibility = [
        "//cmd/...",
        "//pkg/...",
    ],
    deps = [
        "//third_party/go/github.com/Masterminds/semver/v3",
        "//third_party/go/github.com/rs/zerolog",
    ],
)
