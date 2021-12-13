load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "bazel-go-confluent-kafka-showcase",
    deps = ["@com_github_confluentinc_confluent_kafka_go//kafka"],
)

go_binary(
    name = "binary",
    embed = [":lib"],
)

go_image(
    name = "image",
    binary = ":binary",
)
