// <tests> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~noheader~~ TU
// REASON target("bincode")

#include <CppUTest/UtestMacros.h>
#include <cstdio>

// begin external includes

#include "CppUTest/TestHarness.h"
#include "CppUTest/CommandLineTestRunner.h"

// end external includes

// begin macro defines

#define NEW_TEST_GROUP(testGroup) TEST_GROUP(testGroup) {};

// end macro defines

// ~~forward_def_main~~ NEXT_ITEM
// REASON target("bincode")
int main(int argc, const char** argv);

// PROJECT_SCOPE TESTS SEGMENT
// begin tests

NEW_TEST_GROUP(SanityCheck)

TEST(SanityCheck, OnePlusOne) {
    CHECK_EQUAL(1 + 1, 2);
}

// end tests

int main(int argc, const char** argv) {
    return CommandLineTestRunner::RunAllTests(argc, argv);
}
