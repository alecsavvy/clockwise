# core entities, repositories, interactors, data sources, transport layers

In here are the structs and adapters that let all the different services we have communicate through predictable channels. This is so database models don't bubble up to the API and chain transaction structs aren't directly used in db logic. Etc...

Also so we can have multiple implementations of transport layers without changing chain and db interfaces.
