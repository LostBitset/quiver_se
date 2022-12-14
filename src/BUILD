load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "libcode_true_hdrs",
    srcs = [],
    hdrs = glob(["lib/*_hdr.h"]),
)

cc_library(
    name = "libcode",
    srcs = [],
    hdrs = glob(["lib/*.hpp"]),
    deps = [
        ":libcode_true_hdrs",
    ],
)

cc_binary(
    name = "bincode",
    srcs = ["entry.cpp"],
    deps = [
        ":libcode",
    ],
)
