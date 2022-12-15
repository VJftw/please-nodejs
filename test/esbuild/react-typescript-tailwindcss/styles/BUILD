genrule(
    name = "app",
    tools = ["//test/esbuild/react-typescript-tailwindcss/third_party/binary:tailwindcss"],
    srcs = {
        "config": ["tailwind.config.js"],
        "input": ["app.css"],
        "content": ["//test/esbuild/react-typescript-tailwindcss:react-typescript-tailwind"],
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
    visibility = ["//test/esbuild/react-typescript-tailwindcss/..."],
)
