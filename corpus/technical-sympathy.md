---
title: "Technical sympathy"
date: 2020-04-18T06:52:59-04:00
location: "West Lafayette, IN"
---

>"Nobody on the founding team is technical, so I'm trying to recruit a developer. Since you have a technical background, do you have any advice?"

Whenever I've talked to a "non-technical" founder of a startup previously, I've answered this common question in one of two ways. Either learn to code well enough to build your first prototype yourself, or find a competent developer who also happens to be an amazing cofounder -- don't dilute your founding team for sake of technical competency, and _definitely_ don't outsource it in any way if it's core to your product.

But my answer has become more nuanced recently, to refer to something I call _technical sympathy_.

## Mechanical sympathy

I borrowed the terminology "technical sympathy" from the computer science lexicon, from an idea called _mechanical sympathy_.

I'm a web developer on most days, and mostly work on high level software. This means I write code at a pretty high level of abstraction, and concern myself with ideas like user interface elements, webpages, animations, and the like. It means I don't really spend time thinking about the low-level details of my software, like exactly which processor instructions will be executed in my program, or how many nanoseconds each command will take, or whether using a particular kind of micro-optimization to an algorithm will make a 2% difference in a particular routine.

But there are folks in the industry whose bread and butter is working much closer "to the metal," concerned with the details of how software objects get translated into machine instructions, and sweating the details of nanoseconds and tenths of percents of performance -- somebody has to worry about these details.

See, software engineering is a pyramid of abstractions. At the bottom of the pyramid are the hardware -- processor design and fabrication, low-level computing instructions, memory and cache, and so forth. A notch above that are lower level software ideas like processes, threads, mutexes, and system calls. Above that are libraries, frameworks, and toolkits that progressively enable whoever's writing the program to worry less about the lower-level details of what the computer's processors are doing, and spend those extra units of thought on software design and sustainability. This hierarchy of layered abstractions isolate developers who work higher-up in the pyramid from having to worry about exactly what's happening at the lowest levels of a computer system.

But even for developers like me who mostly work closer to the top of the computing-abstraction pyramid, there are times when we have to understand how the lower levels are built. For example, if a particular program I'm writing is slow and inefficient, I might want to understand exactly what's happening at the processor level to take advantage of inefficiencies in an existing design introduced by those layers of abstraction, to make the algorithm more efficient. By this thinking, even when we're working on a product at a level pretty removed from the technical "implementation details," it's helpful to have a basic feel for how the lower levels of the pyramid are structured, and what's more efficient, more predictable, or less error-prone. This kind of understanding, even if imprecise, helps me as an engineer make better high-level decisions.

Engineers call this kind of awareness that pierces beneath the immediate layer of abstraction _mechanical sympathy_ -- to have sympathy about the mechnical, inner workings of system is to be a more effective engineer, because the deeper understanding leads to better design and decisions, with fewer mistakes.

## Technical sympathy

A diverse founding team is, if I may stretch the metaphor a bit, sort of a collection of abstractions. A more technical founder may be in charge of product development and distribution, while a more sales-savvy founder may be in charge of marketing and lead generation. The sales details are abstracted away from the developer, and the code is abstracted away from the salesperson. But the abstraction shouldn't stand in the way of each founder understanding the domains of other teammates in a way that ultimately lead to better decisions and fewer mistakes.

I think technical founders get the short end of the stick more often than other types of founders here. A technical founder is usually expected to have an awareness about how the sales, marketing, or fundraising operations of a startup are going. But a non-technical founder is easily excused for having only a surface-level understanding of the reasons particular technical decisions were made, let alone understanding deeper technical details about why certain libraries or frameworks were used, why certain architectures are preferred, and why one feature was quicker to bring to market than another.

If you're a non-technical founder looking to recruit technical founding talent, also know that having a developer is the baseline. You'll be able to make much more salient decisions about which features to prioritize, which product differentiators to bet on, and when to expand the technical team if your understanding pierces through the technical-nontechnical barrier to have the _technical sympathy_ to understand how the engineered systems that your product depends on work, and where the constraints are.

So if you're building a technical product, and nobody on the team is technical, you actually have two problems to solve. One, find a great technical cofounder. Two, lean on the technical cofounder not only to build, but to impart their technical reasoning and decision-making tools to the rest of the team. Invest in your collective technical sympathy.
