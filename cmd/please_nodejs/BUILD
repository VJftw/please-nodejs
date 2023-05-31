subinclude("///go//build_defs:go")

go_binary(
    name = "please_nodejs",
    srcs = [
        "install.go",
        "main.go",
        "npmpackagedeps.go",
        "packagejson.go",
    ],
    static = True,
    visibility = ["PUBLIC"],
    deps = [
        "///third_party/go/github.com_mattn_go-isatty//:go-isatty",
        "///third_party/go/github.com_peterbourgon_mergemap//:mergemap",
        "///third_party/go/github.com_rs_zerolog//:zerolog",
        "///third_party/go/github.com_rs_zerolog//log",
        "///third_party/go/github.com_urfave_cli_v2//:v2",
        "//pkg/npm",
        "//pkg/npm/pnpm",
        "//pkg/please",
        "//pkg/please/nodejs",
    ],
)
