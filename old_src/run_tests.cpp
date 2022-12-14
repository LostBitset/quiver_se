// <run_tests> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~noheader~~ TU
// REASON TESTING 

// ~~forced_include~~ NEXT_ITEM
// REASON IMPLICIT_ACCESS
// ~~include_cpp~~ NEXT_ITEM
// REASON PROPOGATE ~~nonstandard_tu~~
// ~~use_deferred_include_guard~~ NEXT_ITEM
// REASON CANNOT_LINK_OTHERWISE
// OVERRIDE_NEXT_ITEM SEGMENT
// begin

#ifndef __GUARD_TESTS
#include "tests.cpp"
#endif

// end

#include <cstdio>

// begin external includes

#include "CppUTest/CommandLineTestRunner.h"

// end external includes

int main(int argc, const char** argv) {
    return CommandLineTestRunner::RunAllTests(argc, argv);
}
