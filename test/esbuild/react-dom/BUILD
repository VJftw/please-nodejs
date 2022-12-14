subinclude("//build/defs:nodejs")

nodejs_esbuild_bundle(
    name = "react-dom",
    srcs = [
        "app.jsx",
    ],
    browser = True,
    toolchain = "//third_party/nodejs:nodejs",
    deps = [
        "//test/esbuild/react-dom/third_party/nodejs:react-dom_18.2.0",
        "//test/esbuild/react-dom/third_party/nodejs:react_18.2.0",
    ],
)
