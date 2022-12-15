remote_file(
    name = "jq",
    binary = True,
    url = f"https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64",
    visibility = ["PUBLIC"],
)

remote_file(
    name = "pnpm",
    binary = True,
    url = f"https://github.com/pnpm/pnpm/releases/download/v7.17.1/pnpm-linux-x64",
    visibility = ["PUBLIC"],
)

remote_file(
    name = "go-deps",
    binary = True,
    url = f"https://github.com/Tatskaari/go-deps/releases/download/v1.4.2/go_deps_v1.4.2_linux_amd64",
)

remote_file(
    name = "caddy",
    binary = True,
    exported_files = ["caddy"],
    extract = True,
    url = f"https://github.com/caddyserver/caddy/releases/download/v2.6.2/caddy_2.6.2_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
)
