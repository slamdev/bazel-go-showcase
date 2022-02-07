load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "bazel-go-showcase",
)

go_binary(
    name = "binary",
    embed = [":lib"],
)

go_image(
    name = "image",
    binary = ":binary",
)
