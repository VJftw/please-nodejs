genrule(
    name = "app",
    srcs = {
        "config": ["tailwind.config.js"],
        "input": ["app.css"],
        "content": ["//examples/esbuild/browser/react-typescript-tailwindcss:react-typescript-tailwindcss"],
    },
    outs = ["app.min.css"],
    cmd = """
    $TOOLS \\
        --config "$SRCS_CONFIG" \\
        --minify \\
        --input "$SRCS_INPUT" \\
        --output "$OUTS" \\
        --content "**/*"
    """,
    tools = ["//examples/esbuild/browser/react-typescript-tailwindcss/third_party/binary:tailwindcss"],
    visibility = ["//examples/esbuild/browser/react-typescript-tailwindcss/..."],
)
