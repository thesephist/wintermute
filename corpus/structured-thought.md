---
title: "A runtime for structured thought"
date: 2020-05-11T11:14:25-04:00
location: "West Lafayette, IN"
---

When we write thoughts and ideas down with words for future recollection and synthesis, we take the stuff of abstract concepts, and collapse them down into a one-dimensional string of characters. For even the best, most careful writers, I think our ideas lose a lot of fidelity in this dramatic reduction in expressiveness of thoughts and their relationships to each other.

Our thoughts _in vivo_ take on all kinds of shapes, and are connected in a myriad of relationships -- by analogy, by example, by superset/subset relations, and so on. When we taxidermy them out onto paper or through the keyboard, we're forced to flatten these relationships out onto a single dimension, usually nested bulleted links with verbal references  that leap across the page arbitrarily to connect related ideas to each other. Just think of the last time you jot down an idea in your note-taking method of choice -- how much freedom did the medium afford you? At best, you're writing freeform, pen-on-paper, and you draw arrows between related ideas and underline the important parts; typing on a keyboard into a text box is even more restrictive.

The internal structure of individual ideas vary even more greatly. Some thoughts contain lists of things, others are about comparisons between categories of things, and there are probably many more that can't be named or enumerated or denoted with just a few lines going across the page between sentences. The state of the art in high-fidelity recording of our thoughts, in my opinion, is still handwriting on paper (or your digital handwriting app of choice). But here, we effectively lose our ability to search through information quickly, when we need it, because unlike bulleted lists, the information we hand-write isn't _structured_.

This yields the question, **how should we store and retrieve ideas in written form, so our workflows better preserve the living structure of thoughts?**

<p>
<img src="/img/structured-thought.jpg" alt="Structured thought" class="blend-multiply">
</p>

I've previously noted [the parallelism and contrast between programming and written languages](https://linus.coffee/note/abstraction/). I think the more precise, multi-dimensional language of computers may be better suited for articulating relationships between our ideas in the written form.

## Structured thought

Why might programming languages be better suited for recording relationships between ideas than the written word?

One reason is that programming languages are purpose-designed for building up living data structures. Programs, in their normal course of life, build up precise lists and trees and other complex structures of values. If there were a programming language as such, whose primitive values were units of thought or ideas instead of integers and strings and arrays, would we be able to record and retrieve our past ideas more effectively? In other words, what if there was some sort of a [domain-specific language](https://en.wikipedia.org/wiki/Domain-specific_language) for better note-taking? Can we view the macro-structure of our thoughts as a highly specialized, highly complex data structure that we can encode into programs? Probably not entirely, but for the more analytical ideas, I think it's a thought experiment we can entertain using new and interesting tools.

The most promising tool I've found exploring this direction is [Roam Research](https://roamresearch.com). It stores ideas in a fluid, cross-linkable and searchable database of [hypertext](https://en.wikipedia.org/wiki/Hypertext). But Roam is still pretty restricted to nested relationships and bidirectional links between ideas (though I admittedly haven't tried it in depth). I think there's lots of headroom here in the market for exploring more general data structures for ideas, and for allowing note-takers to create their own structures and layouts.

## A runtime for human abstractions

I think an ideal future extension of this idea isn't just a written format that can be queried like Roam or some kind of a programming language, but an omnipresent, ambient [_runtime_](https://en.wikipedia.org/wiki/Runtime_system) for our thoughts and ideas -- one that can interact fluidly with our environment and context, and maybe even preempt our recollections and augment our ability to sift through our ideas from the past and present.

Sci-fi literature has imagined an endless taxonomy of ways for humans to externalize our minds onto offboard computers and networks of mechanical brains, most of which yield the control of our memory storage and recollection to some artificial and autonomous system. But I think there are ways to extend our minds using more primitive tools like better notes, and a more carefully crafted runtime for structured thought. One where we continue to drive interactions with our memory, albeit with more powerful tools in hand to record what we think, and remember them at the right times, when we need them.
