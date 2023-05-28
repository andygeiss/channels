# channels

This package provides reusable functions for creating, splitting, merging and multiplexing channels of any type to build effective pipelines.
The functions automatically use the available CPU cores and create and close the channels.

                          3)                               6)
                         ┌───────┐                        ┌───────┐
                   ┌─────►Process├─────┐          ┌───────►Process├─────┐
                   │     └───────┘     │          │       └───────┘     │
     1)          2)│                 4)│      5)  │                   7)│      8)
    ┌────────┐  ┌──┴──┐  ┌───────┐  ┌──▼──┐  ┌────┴────┐  ┌───────┐  ┌──▼──┐  ┌─────┐
    │Generate├──►Split├──►Process├──►Merge├──►Multiplex├──►Process├──►Merge├──►Drain│
    └────────┘  └──┬──┘  └───────┘  └──▲──┘  └────┬────┘  └───────┘  └──▲──┘  └─────┘
                   │                   │          │                     │
                   │     ┌───────┐     │          │       ┌───────┐     │
                   └─────►Process├─────┘          └───────►Process├─────┘
                         └───────┘                        └───────┘

