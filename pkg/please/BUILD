subinclude("///go//build_defs:go")

go_library(
    name = "please",
    srcs = [
        "buildfile.go",
        "buildfilemanager.go",
        "please.go",
        "rule.go",
        "target.go",
    ],
    visibility = [
        "//cmd/...",
        "//pkg/...",
    ],
    deps = [
        "//third_party/go/github.com/bazelbuild/buildtools",
        "//third_party/go/github.com/rs/zerolog",
    ],
)

go_test(
    name = "please_test",
    srcs = [
        # "buildfile_test.go",
        "target_test.go",
    ],
    external = True,
    deps = [
        ":please",
        "//third_party/go/github.com/stretchr/testify",
    ],
)
