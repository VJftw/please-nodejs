subinclude("///go//build_defs:go")

go_module(
    name = "zerolog",
    install = [
        ".",
        "internal/json",
        "log",
    ],
    module = "github.com/rs/zerolog",
    version = "v1.28.0",
    visibility = ["PUBLIC"],
    deps = ["//third_party/go/github.com/mattn/go-colorable"],
)
