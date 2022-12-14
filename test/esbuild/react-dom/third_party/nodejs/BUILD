subinclude("///nodejs//build/defs:nodejs")

nodejs_npm_package(
    name = "react_18.2.0",
    package_name = "react",
    license = "MIT",
    registry = "https://registry.npmjs.org",
    version = "18.2.0",
    visibility = ["PUBLIC"],
    deps = ["//test/esbuild/react-dom/third_party/nodejs:loose-envify_1.4.0"],
)

nodejs_npm_package(
    name = "react-dom_18.2.0",
    package_name = "react-dom",
    license = "MIT",
    registry = "https://registry.npmjs.org",
    version = "18.2.0",
    visibility = ["PUBLIC"],
    deps = [
        "//test/esbuild/react-dom/third_party/nodejs:loose-envify_1.4.0",
        "//test/esbuild/react-dom/third_party/nodejs:scheduler_0.23.0",
    ],
)

nodejs_npm_package(
    name = "scheduler_0.23.0",
    package_name = "scheduler",
    license = "MIT",
    registry = "https://registry.npmjs.org",
    version = "0.23.0",
    visibility = ["PUBLIC"],
    deps = ["//test/esbuild/react-dom/third_party/nodejs:loose-envify_1.4.0"],
)

nodejs_npm_package(
    name = "loose-envify_1.4.0",
    package_name = "loose-envify",
    license = "MIT",
    registry = "https://registry.npmjs.org",
    version = "1.4.0",
    visibility = ["PUBLIC"],
    deps = ["//test/esbuild/react-dom/third_party/nodejs:js-tokens_4.0.0"],
)

nodejs_npm_package(
    name = "js-tokens_4.0.0",
    package_name = "js-tokens",
    license = "MIT",
    registry = "https://registry.npmjs.org",
    version = "4.0.0",
    visibility = ["PUBLIC"],
)
