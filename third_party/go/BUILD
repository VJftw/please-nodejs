subinclude("///go//build_defs:go")

subinclude("//:architectures")

go_toolchain(
    name = "toolchain",
    architectures = ARCHITECTURES,
    version = "1.19.2",
    visibility = ["PUBLIC"],
)
