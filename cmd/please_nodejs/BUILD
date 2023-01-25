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
        "//pkg/npm",
        "//pkg/npm/pnpm",
        "//pkg/please",
        "//pkg/please/nodejs",
        "//third_party/go/github.com/mattn/go-isatty",
        "//third_party/go/github.com/rs/zerolog",
        "//third_party/go/github.com/peterbourgon/mergemap",
        "//third_party/go/github.com/urfave/cli/v2",
    ],
)

# subinclude("//build/defs:buildkit")

# _buildkit_tool()
