subinclude("//build/defs:nodejs")

nodejs_package_json(
    name = "package_json",
    extra = {
        "jest": {
            "verbose": True,
            "reporters": [
                "jest-junit",
            ],
        },
    },
    visibility = [],
    deps = [
        "//examples/testing/jest/third_party/nodejs:jest-junit_15.0.0",
        "//examples/testing/jest/third_party/nodejs:jest_29.4.0",
    ],
)

nodejs_test(
    name = "jest_test",
    srcs = [
        "sum.js",
        "sum.test.js",
    ],
    package_json = ":package_json",
    test_cmd = """
    cd $PKG
    set +e; jest; ec=$?
    set -e
    mv junit.xml $RESULTS_FILE
    exit $ec
    """,
)
