#include "quiver.h"

#include <algorithm>
#include <iterator>
#include <utility>

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNode<N, E, C>* QuiverNodeRef::find_in_quiver(Quiver<N, E, C> *quiver) {
    quiver->get_node(this->index);
}

inline bool QuiverNodeRef::operator<(const QuiverNodeRef& other) {
    return this->index < other.index;
}

template <typename E, typename R>
SimpleQuiverEdge<E, R>::SimpleQuiverEdge() {}

template <typename E, typename R>
static SimpleQuiverEdge<E, R> empty() {
    SimpleQuiverEdge<E, R> res;
    return res;
}

template <typename E, typename R>
void SimpleQuiverEdge<E, R>::foreach_key(std::function<void(E)> func) {
    for (std::pair<E, R> kv : this->backing_map) {
        func(kv.first);
    }
}

template <typename E, typename R>
void SimpleQuiverEdge<E, R>::insert(E edge, R node_ref) {
    if (!this->backing_map.contains(edge)) {
        this->backing_map[edge] = node_ref;
    }
}

template <typename E, typename R>
R* SimpleQuiverEdge<E, R>::fwd_lookup(E edge) {
    return this->backing_map.at(&edge);
}

template <typename E, typename R>
std::vector<E> SimpleQuiverEdge<E, R>::rev_lookup(R node_ref) {
    std::vector<E> res;
    auto pred = [=](std::pair<E, R> kv) {
        return kv->second.index == node_ref.index;
    };
    auto xform = [=](std::pair<E, R> kv) {
        return kv->first;
    };
    std::copy_if(
        this->backing_map.begin(),
        this->backing_map.end(),
        std::back_inserter(res),
        pred
    );
    std::transform(
        res.begin(),
        res.end(),
        res.begin(),
        xform
    );
    return res;
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNode<N, E, C>::QuiverNode(N set_value)
    : value(set_value)
    , edge_container(C::empty())
{}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
N QuiverNode<N, E, C>::get_value() {
    return this->value;
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
C QuiverNode<N, E, C>::get_edge_container() {
    return this->edge_container;
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
std::set<QuiverNodeRef> QuiverNode<N, E, C>::get_parents() {
    return this->parents;
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNodeRef* QuiverNode<N, E, C>::follow_edge_fwd(E edge) {
    return this->edge_container.fwd_lookup(edge);
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
Quiver<N, E, C>::Quiver() {}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
QuiverNodeRef Quiver<N, E, C>::insert_node(N node_value) {
    QuiverNode<N, E, C> new_node {node_value};
    this->arena.push_back(new_node);
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
void Quiver<N, E, C>::insert_edge(QuiverNodeRef src, QuiverNodeRef dst, E edge) {
    src.find_in_quiver(this)->get_edge_container().insert(edge, dst);
    dst.find_in_quiver(this)->get_parents().insert(src);
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
std::vector<std::pair<QuiverNodeRef, E>> Quiver<N, E, C>::follow_all_fwd(
    QuiverNodeRef node_ref
) {
    std::vector<std::pair<QuiverNodeRef, E>> res;
    QuiverNode<N, E, C>* node = node_ref.find_in_quiver(this);
    auto xproc = [=, &res](E edge) {
        res.push_back(std::make_pair(
            *(node->get_edge_container().lookup_fwd(edge)),
            edge
        ));
    };
    node->get_edge_container().foreach_key(xproc);
    return res;
}

template <typename N, typename E, typename C>
requires ReversibleAssoc<C, E, QuiverNodeRef>
std::vector<std::pair<QuiverNodeRef, E>> Quiver<N, E, C>::follow_all_rev(
    QuiverNodeRef node_ref
) {
    std::vector<std::pair<QuiverNodeRef, E>> res;
    QuiverNode<N, E, C>* node = node_ref.find_in_quiver(this);
    for (QuiverNodeRef parent : node->get_parents()) {
        QuiverNode parent_node = parent.find_in_quiver(this);
        std::vector<E> edges = parent_node.get_edge_container().rev_lookup(node_ref);
        auto xform = [=](E edge) {
            return std::make_pair(parent, edge);
        };
        std::transform(
            edges.begin(),
            edges.end(),
            edges.begin(),
            xform
        );
        res.reserve(edges.size());
        res.insert(res.end(), edges.begin(), edges.end());
    }
    return res;
}
