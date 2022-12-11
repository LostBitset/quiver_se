// <quiver> -*- C++ -*-

/*!
    A doubly-linked quiver, abstracted over the node, edge, and edge container types. Edge
    containers must be some form of reversible associative data structure, as they are placed inside
    each node and map edges to the nodes they lead to. 
*/

#pragma once
#include <concepts>
#include <cstddef>
#include <functional>
#include <iterator>
#include <map>
#include <set>
#include <sys/types.h>
#include <utility>
#include <vector>

/*! A concept representing a reversible (think doubly-linked) associative data structure. */
template <typename T, typename K, typename V>
concept ReversibleAssoc = requires(T item, K key, V value, std::function<void(K)> func) {
    { T::empty() } -> std::convertible_to<T>;
    { item.foreach_key(func) } -> std::same_as<void>;
    { item.insert(key, value) } -> std::same_as<void>;
    { item.fwd_lookup(key) } -> std::convertible_to<V*>;
    { item.rev_lookup(value) } -> std::convertible_to<std::vector<K>>;
};

// begin forward declarations

struct QuiverNodeRef;

template <typename E, typename R = QuiverNodeRef>
class SimpleQuiverEdge;

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
class QuiverNode;

template <typename N, typename E, typename C = SimpleQuiverEdge<E>>
requires ReversibleAssoc<C, E, QuiverNodeRef>
class Quiver;

// end forward declarations

/*! A simple set of quiver edges from a single node. Reverse lookups are slow. */
template <typename E, typename R>
class SimpleQuiverEdge {

    public:
    explicit SimpleQuiverEdge();
    SimpleQuiverEdge(SimpleQuiverEdge<E, R>& other) = default;
    SimpleQuiverEdge(SimpleQuiverEdge<E, R>&& other) = default;

    // begin ReversibleAssoc methods
    static SimpleQuiverEdge<E, R> empty();
    void foreach_key(std::function<void(E)> func);
    void insert(E edge, R node_ref);
    R* fwd_lookup(E edge);
    std::vector<E> rev_lookup(R node_ref);
    // end ReversibleAssoc methods

    private:
    std::map<E, R> backing_map;

};

/*! A doubly-linked quiver, abstracted over the node, edge, and edge container types. */
template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
class Quiver {
    
    public:
    explicit Quiver();
    Quiver(Quiver<N, E, C>& other) = default;
    Quiver(Quiver<N, E, C>&& other) = default;

    QuiverNodeRef insert_node(N node);
    void insert_edge(QuiverNodeRef src, QuiverNodeRef dst, E edge);
    std::vector<std::pair<QuiverNodeRef, E>> follow_all_fwd(QuiverNodeRef node_ref);
    std::vector<std::pair<QuiverNodeRef, E>> follow_all_rev(QuiverNodeRef node_ref);

    private:
    std::vector<QuiverNode<N, E, C>> arena;

};

/*! A reference to a node inside a quiver. This is just an index, and you need the Quiver object. */
struct QuiverNodeRef {

    size_t index;

    template <typename N, typename E, typename C>
    requires ReversibleAssoc<C, E, QuiverNodeRef>
    QuiverNode<N, E, C>* find_in_quiver(Quiver<N, E, C>* quiver);

    /*! Needed for storing QuiverNodeRef objects in an std::set. */
    inline bool operator<(const QuiverNodeRef& other);

};

/*! A node inside a quiver. Not usable without the Quiver object. */
template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
class QuiverNode {

    public:
    explicit QuiverNode(N set_value);
    QuiverNode(QuiverNode<N, E, C>& other) = default;
    QuiverNode(QuiverNode<N, E, C>&& other) = default;

    QuiverNodeRef* follow_edge_fwd(E edge);

    N get_value();
    C get_edge_container();
    std::set<QuiverNodeRef> get_parents();

    private:
    N value;
    C edge_container;
    std::set<QuiverNodeRef> parents;
};
