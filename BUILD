genrule(
    name = "architectures",
    srcs = ["ARCHITECTURES"],
    outs = ["architectures.build_defs"],
    cmd = """
    set -Eeuo pipefail
    sed -i 's#/#_#g' "$SRCS"
    grep -v "^#" "$SRCS" > "${SRCS}.new"
    mv "${SRCS}.new" "$SRCS"
    echo "ARCHITECTURES = $($TOOLS -ncR '[inputs]' $SRCS)" > $OUT
    """,
    tools = ["//third_party/binary:jq"],
    visibility = ["PUBLIC"],
)

genrule(
    name = "version",
    outs = ["version.build_defs"],
    cmd = """
    echo "VERSION = \\"$(git describe --always)\\"" > $OUTS
    """,
    visibility = ["PUBLIC"],
)
