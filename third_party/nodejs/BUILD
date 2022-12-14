subinclude("//build/defs:nodejs")

nodejs_toolchain(
    name = "nodejs",
    version = "18.12.0",
)

nodejs_esbuild_toolchain(
    name = "esbuild",
    version = "0.15.16",
)

remote_file(
    name = "node-v18.5.0-linux-x64",
    url = "https://github.com/vercel/pkg-fetch/releases/download/v3.4/node-v18.5.0-linux-x64",
    visibility = ["PUBLIC"],
)
