subinclude("//build/defs:nodejs")

nodejs_toolchain(
    name = "nodejs",
    version = "18.12.0",
)

nodejs_esbuild_toolchain(
    name = "esbuild",
    version = "0.15.16",
)

nodejs_npm_package(
    name = "pkg",
    version = "5.8.0",
    toolchain = "//third_party/nodejs:nodejs",
)
