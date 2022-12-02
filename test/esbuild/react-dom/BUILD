subinclude("//build/defs:nodejs")

nodejs_esbuild_bundle(
    name = "react-dom",
    srcs = [
        "app.jsx",
    ],
    # browser = True,
    deps = [
        # TODO: these deps still install their dependencies
        "//test/esbuild/react-dom/third_party/nodejs:react",
        "//test/esbuild/react-dom/third_party/nodejs:react-dom",
    ],
    toolchain = "//third_party/nodejs:nodejs",
)
