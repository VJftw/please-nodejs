subinclude("//build/defs:nodejs")

nodejs_bundle_browser(
    name = "react",
    srcs = [
        "app.jsx",
    ],
    build_cmd = """
    esbuild \\
        --bundle --minify \\
        --outdir="$OUT" \\
        app.jsx
    """,
    public = [
        "index.html",
    ] + glob("assets/**"),
    deps = [
        "//examples/esbuild/binary/react-server/third_party/nodejs:esbuild_0.16.8",
        "//examples/esbuild/binary/react-server/third_party/nodejs:react-dom_18.2.0",
        "//examples/esbuild/binary/react-server/third_party/nodejs:react_18.2.0",
    ],
)
