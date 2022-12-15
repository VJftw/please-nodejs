subinclude("//build/defs:nodejs")

nodejs_esbuild_bundle_binary(
    name = "binary",
    srcs = [
        "app.jsx",
    ],
    deps = [
        "//test/esbuild/react-dom/third_party/nodejs:react-dom_18.2.0",
        "//test/esbuild/react-dom/third_party/nodejs:react_18.2.0",
    ],
)
