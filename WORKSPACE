workspace(name = "io_istio_api")

git_repository(
    name = "io_bazel_rules_go",
    commit = "eba68677493422112dd25f6a0b4bbdb02387e5a4",  # Aug 1, 2017
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")

go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")

go_proto_repositories()
