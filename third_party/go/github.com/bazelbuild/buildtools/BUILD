subinclude("///go//build_defs:go")

go_module(
    name = "buildtools",
    install = [
        "api_proto",
        "build",
        "build_proto",
        "edit",
        "file",
        "labels",
        "lang",
        "tables",
        "wspace",
    ],
    module = "github.com/bazelbuild/buildtools",
    version = "v0.0.0-20221110131218-762712d8ce3f",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go/github.com/golang/protobuf",
        "//third_party/go/google.golang.org/protobuf",
    ],
)
