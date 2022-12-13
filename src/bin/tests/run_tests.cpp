// <run_tests> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~noheader~~ TU
// REASON target("test")

// ~~forced_include~~ NEXT_ITEM
// REASON IMPLICIT_ACCESS
// ~~include_cpp~~ NEXT_ITEM
// REASON PROPOGATE ~~nonstandard_tu~~
#include "tests.cpp"

#include <cstdio>

// begin external includes

#include "CppUTest/CommandLineTestRunner.h"

// end external includes

int main(int argc, const char** argv) {
    return CommandLineTestRunner::RunAllTests(argc, argv);
}
