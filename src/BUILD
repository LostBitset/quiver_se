load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "libcode",
    srcs = [],
    hdrs = glob(["lib/*.hpp"]),
)

cc_binary(
    name = "bincode",
    srcs = ["entry.cpp"],
    deps = [
        ":libcode",
    ],
)
