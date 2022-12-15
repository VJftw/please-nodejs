subinclude("//build/defs:nodejs")

nodejs_toolchain(
    name = "nodejs",
    version = "18.12.0",
)

nodejs_esbuild_toolchain(
    name = "esbuild",
    version = "0.15.16",
)

nodejs_vercel_pkg_binaries(
    name = "pkg_binaries",
    extra_os_architectures = ["macos-amd64"],
    node_version = "v18.5.0",
    version = "v3.4",
    visibility = ["PUBLIC"],
)
