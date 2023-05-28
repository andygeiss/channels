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

- [x] Drain receives each value from an input channel and does something with it.
- [x] Generate sends each value to the out channel.
- [x] Merge sends each value from a list of input channels to the out channel.
- [x] Multiplex sends each value from an input channel to every output channel.
- [x] Process works on each value from an input channel and sends the result to an output channel.
- [x] Split sends each value from an input channel to one of the output channels.
