subinclude("//build/defs:nodejs")

nodejs_esbuild_bundle_browser(
    name = "react-typescript",
    srcs = [
        "tsconfig.json",
    ],
    entrypoints = [
        "app.tsx",
    ],
    public = [
        "index.html",
    ] + glob("assets/**"),
    deps = [
        "//test/esbuild/react-typescript/third_party/nodejs:@types_react-dom_18.0.9",
        "//test/esbuild/react-typescript/third_party/nodejs:@types_react_18.0.26",
        "//test/esbuild/react-typescript/third_party/nodejs:react-dom_18.2.0",
        "//test/esbuild/react-typescript/third_party/nodejs:react_18.2.0",
    ],
)
