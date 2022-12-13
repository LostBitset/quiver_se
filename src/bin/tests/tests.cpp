// <tests> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~no_header~~ TU
// REASON target("test")

// begin external includes

#include "CppUTest/TestHarness.h"

// end external includes

// PROJECT_SCOPE TESTS SEGMENT
// begin tests

TEST_GROUP(SanityCheck) {};

TEST(SanityCheck, OnePlusOne) {
    CHECK_EQUAL(1 + 1, 2);
}

// end tests
