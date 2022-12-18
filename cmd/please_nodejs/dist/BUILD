subinclude("//:architectures")

subinclude("//:version")

def cross_compile(arch):
    return genrule(
        name = f"please_nodejs_{arch}",
        srcs = [f"///{arch}//cmd/please_nodejs"],
        outs = [f"please_nodejs_{VERSION}_{arch}"],
        cmd = "mv $SRC $OUT",
    )

filegroup(
    name = "release_files",
    srcs = [cross_compile(arch) for arch in ARCHITECTURES],
    labels = ["hlink:plz-out/package"],
)
