# dds(distribute-download-system) design doc

## what is dds

dds is a distribute download system that let user share there band width
main target:
  share band with

## concepts

- Cluster
  first of all, we've assume that all nodes(in a single cluster) are connected using LAN etc.
  The speed between each other can be regard as infinit, if you want connect 2 clusters,
  your speed between 2 cluster must can be regard as infinit, so we recommand you not using
  it under environment like MAN, WAN or internet, just LAN only.
  nodes between 2 cluster(if connected, we may regard it as 1 cluster) mapping is like below
    ------                                  ------
    | M1 |                                  | N1 |
    ------                                  ------
          \                                /
           \                              /
            \                            /
             ------  speed -> inf  ------
             | M3 | <------------> | N3 |
             ------                ------
            /                            \
           /                              \
          /                                \
    ------                                  ------
    | M2 |                                  | N2 |
    ------                                  ------
    Cluster M                               Cluster N
- Frends
  all nodes maintain a global nodes hashmap(not all nodes), add or remove node will change
  the mapping,node change frequency too high may lead cluster not stable, so we desgin a
  concept of frends, node change only spread between frends; besides, only alive node will
  be add to mapping, if node currently dead, syncing will not add it.
- Task
  all download request should be separete into many tasks, machine in our cluster will
  find tasks and peek a task to do and return to someone who made it.

## techs

- golang 1.8
- thrift 0.10.0
- go-restful

## what download protocol we support

- http
- https