subinclude("///go//build_defs:go")

go_module(
    name = "protobuf",
    install = [
        "jsonpb",
        "proto",
    ],
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
    visibility = ["PUBLIC"],
    deps = ["//third_party/go/google.golang.org/protobuf"],
)
