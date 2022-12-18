subinclude("//build/defs:nodejs")

nodejs_bundle_binary(
    name = "react-server",
    srcs = [
        "app.jsx",
    ],
    build_cmd = """
    esbuild \\
        --bundle --minify \\
        --outdir="$OUT" \\
        app.jsx
    """,
    deps = [
        "//examples/esbuild/binary/react-server/third_party/nodejs:esbuild_0.16.8",
        "//examples/esbuild/binary/react-server/third_party/nodejs:react-dom_18.2.0",
        "//examples/esbuild/binary/react-server/third_party/nodejs:react_18.2.0",
    ],
)
