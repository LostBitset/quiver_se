// <tests> -*- C++ -*-
// ~~nonstandard_tu~~

// ~~no_header~~ TU
// REASON target("test")

#include "../../lib/quiver.h"

// begin external includes

#include "CppUTest/TestHarness.h"
#include <CppUTest/UtestMacros.h>
#include <utility>

// end external includes

// PROJECT_SCOPE TESTS SEGMENT
// begin tests

TEST_GROUP(SanityCheck) {};

TEST(SanityCheck, OnePlusOne) {
    CHECK_EQUAL(1 + 1, 2);
}

TEST_GROUP(QuiverTU) {};

TEST(QuiverTU, SimpleFwdCheck) {
    Quiver<int,int> q;
    auto n1 = q.insert_node(45);
    auto n2 = q.insert_node(46);
    auto n3 = q.insert_node(47);
    q.insert_edge(n1, n2, 0);
    q.insert_edge(n1, n2, 1);
    q.insert_edge(n1, n2, 2);
    q.insert_edge(n1, n3, 77);
    q.insert_edge(n3, n1, 73);
    
    std::pair<QuiverNodeRef, int> expected[4] {
        std::make_pair(n2, 0),
        std::make_pair(n2, 1),
        std::make_pair(n2, 2),
        std::make_pair(n3, 77)
    };
    int seen_bitmask = 0x0;
    int expected_seen_bitmask = 0x0;
    auto fwd_vector = q.follow_all_fwd(n1);
    const int expected_fwd_size =
        sizeof(expected) == 0 ? 0 : (sizeof(expected)/sizeof(expected[0]));
    CHECK_EQUAL(expected_fwd_size, fwd_vector.size());
    for (int i = 0; i < expected_fwd_size; i++) {
        expected_seen_bitmask |= (1 << i);
        for (int j = 0; j < expected_fwd_size; j++) {
            if (
                true
                && fwd_vector[i].first.index == expected[j].first.index
                && fwd_vector[i].second == expected[j].second
            ) {
                int field_bit = 1 << j;
                CHECK_EQUAL(0, seen_bitmask & field_bit);
                seen_bitmask |= field_bit;
            }
        }
    }
    CHECK_EQUAL(expected_seen_bitmask, seen_bitmask);
};

// end tests
