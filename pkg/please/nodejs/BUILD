subinclude("///go//build_defs:go")

go_library(
    name = "nodejs",
    srcs = [
        "npmpackage.go",
        "npmpackageresolver.go",
    ],
    visibility = ["//cmd/..."],
    deps = [
        "//pkg/npm",
        "//pkg/please",
        "//third_party/go/github.com/bazelbuild/buildtools",
        "//third_party/go/github.com/rs/zerolog",
    ],
)
