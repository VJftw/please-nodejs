subinclude("//build/defs:nodejs")

nodejs_bundle_browser(
    name = "react-typescript-tailwindcss",
    srcs = [
        "App.tsx",
        "contacts/Contact.tsx",
        "contacts/EditContact.tsx",
        "contacts/actions.tsx",
        "contacts/data.tsx",
        "tsconfig.json",
    ],
    build_cmd = """
    esbuild \\
        --bundle --minify --sourcemap \\
        --outdir="$OUT" \\
        App.tsx
    """,
    public = [
        "index.html",
        "//examples/esbuild/browser/react-typescript-tailwindcss/styles:app",
    ] + glob("assets/**"),
    visibility = ["//examples/esbuild/browser/react-typescript-tailwindcss/styles/..."],
    deps = [
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:@types_localforage_0.0.34",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:@types_match-sorter_6.0.0",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:@types_react-dom_18.0.9",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:@types_react_18.0.26",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:@types_sort-by_1.2.0",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:esbuild_0.16.8",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:localforage_1.10.0",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:match-sorter_6.3.1",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:react-dom_18.2.0",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:react-router-dom_6.4.5",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:react_18.2.0",
        "//examples/esbuild/browser/react-typescript-tailwindcss/third_party/nodejs:sort-by_1.2.0",
    ],
)
