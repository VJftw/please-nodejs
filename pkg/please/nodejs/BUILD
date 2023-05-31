subinclude("///go//build_defs:go")

go_library(
    name = "nodejs",
    srcs = [
        "npmpackage.go",
        "npmpackageresolver.go",
    ],
    visibility = ["//cmd/..."],
    deps = [
        "///third_party/go/github.com_rs_zerolog//log",
        "//pkg/npm",
        "//pkg/please",
    ],
)
