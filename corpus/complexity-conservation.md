---
title: "Conservation of complexity"
date: 2020-05-14T14:16:13-04:00
location: "West Lafayette, IN"
---

I want to share and expand on a blog post I read this week by Fred Hebert on software complexity -- ["Complexity has to live somewhere"](https://ferd.ca/complexity-has-to-live-somewhere.html). I think it embodies the main philosophy I carry through the design of my side projects and libraries like [Torus](https://github.com/thesephist/torus).

## Complex problems require complex solutions

The leading idea of the blog is that if a system is modeling a problem that is inherently complex, that complexity needs to live somewhere -- we cannot simply eliminate complexity in the problem domain with good engineering; instead, good engineering isolates complexity to the parts of a system where that complexity is easiest for the team to manage over time.

In other words, **complexity is conserved**. A complex problem will have a complex solution. Fred's blog discusses this at length, so I won't belabor this point. Here, it's enough to note that real world problems have no shortage of exceptions, edge cases, and situational differences that the software that models it has to know about.

A classic case of this conservation of complexity is [software internationalization](https://en.wikipedia.org/wiki/Internationalization_and_localization). Human language is complicated and littered with exceptions, like corner cases for pluralizing nouns and for conjugating verbs to match pronouns. That complexity about human language has to live somewhere -- either in the translator's mind, in the software handling translation, or in some library buried deep underneath the technical stack that everyone depends on.

Good engineering, then, isolates and marshals complexity inherited from the problem to the parts of a system where it's easiest to manage. But more than simply the ease of managing complexity now, I think we should optimize for encoding problem complexity in places that are easiest to _change and adapt_ as the real world changes around us, and as our requirements change.

Where might that be?

## Complex interfaces > complex components

I think we should try to **push the conserved complexity from our problems towards the interfaces between our software components**, rather than try to contain and abstract it away into idealized reusable parts.

Software is [stacks of abstraction layers](/posts/dof/), and we task each layer or component in our software design with encoding away a small bit of domain-specific knowledge in the underlying problem.

Sometimes, I think we get carried away in this way of thinking and try to hide or ignore complexity in pursuit of ever more elegant or minimal interfaces for our components.

Fred's blog takes a build system as a case study: the simpler and more elegant the interface for expressing build constraints, the more susceptible that layer of the stack is to weird edge cases that engineers need to work around. Conversely, if the system has no constraints and an infinitely flexible API, you might as well simply have a library of useful functions, and little else that's helpful.

I think the best solution in times like these is to take care that the complexity "leaks out" through the interfaces of software components in the right ways. Complexity always leaks out, if not through well-defined interfaces, then through users' workarounds and edge cases and error checks. So I think it's best to take control of how they manifest in our design.

Let's not shy away from complex interfaces, if the problem we're solving is itself complex. The best software interfaces aren't as simple or elegant as possible. The best interfaces accurately reflect the complexity of the problem, while hiding away unnecessary detail.
