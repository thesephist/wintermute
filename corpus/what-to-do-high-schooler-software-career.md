---
title: "What to do if you’re a high schooler interested in a career in software"
date: 2019-06-12T00:00:00-00:00
location: "Haight-Ashbury, SF"
---

Someone from the amazing [Hack Club](https://hackclub.com/) Slack asked me a few days ago what they should do, in regards to a career in software / computer science. In the process of jotting down my thoughts, I realized I had a lot to say. So I thought I’d share it in a longer form here.

I’m not an expert — I’m still navigating this space myself, but I’m a firm believer in the idea that there’s value in incomplete and flawed perspectives. So without further ado, here’s what I’d say, if you’re a high schooler interested in building software to make a living. Take it with a grain of salt — I’m only starting out, too.

## First, go for breadth

I think the most helpful approach in the long run is to _go for breadth first_. Explore what’s out there in terms of stacks, languages, types of products, size of company, kinds of development work. Some questions you might not have answered for yourself include…

- Do you like working on UI or more backend things, or infrastructure? What about networking / systems software / graphics? What about programming languages and compilers? Does AI interest you?

- Do you care about shipping things quickly? Or do you care about crafting software with attention to detail?

- What kinds of projects do you enjoy working on the most? Do you care about technical excellence and innovating in tools, or about building a product that has material impact to the end user?

- What size of team do you work best in? Do you want to be able to define exactly what you work on, or would you rather collaborate in an environment where there’s a set roadmap?

These are questions that you can’t answer on your own, but you can start to get a sense of these by either working on side projects like you’ve been doing, getting internships, joining clubs (especially in college), and trying to use the opportunities you have to explore. I think understanding how you want to use your software dev skills is critical to having a firmer feeling of what you want to do with it. Having a diverse set of experiences will help tremendously here, not only to give you perspective, but to help you find those one or two things that you can give deep on.

## Next, and only after that, go deep

The second thing I think is very helpful is to pick a topic in software that you find really interesting and _go deep_. Here are a few examples of what I mean by “going deep”.

- If you’re interested in React as a technology, for example, don’t just get good at using it in building user interfaces. Watch some of the ReactConf talks on YouTube. Really figure out and understand how it works under the hood. Read pull requests on GitHub, read what their maintainers say about JavaScript’s present and future. The great thing about open source software is that it’s not only distributed openly, but usually also built, discussed, disputed, and designed in the open, and being able to see how the sausage is made is an opportunity a lot of students leave on the table. You can go and read up on how Facebook engineers are pushing out new features to React _right now_. Take advantage of these resources.

- If you’re interested in low level operating systems stuff, there are stepping stones to be able to read or write code in the Linux kernel. Learn a bit of C. Install a Linux distribution and try to poke around in it to see how it works. How does a package manager work? Where are your shell commands and executables stored? What happens when you type those commands? At a lower level, how does your computer boot up? All these things are deep rabbit holes and they’re possible to learn very deeply with resources that are freely available online. Use them.

- There are side projects you can do to go really deep in any particular area. Some of these “areas” might be programming languages (PL) and compilers, web / JavaScript, graphics (fascinating topic in my opinion, look at things like path tracing and WebGL), firmware and assembly, hardware, and a lot more.

How you discover what you’re _really_ into will sort of be a searching process. Side projects are useful here, but I’d encourage you not to get stuck in the rut of just making the same kind of project (e.g. “web app”) multiple times because it’s familiar, unless that’s the thing you’re going deep on. But if you are, then don’t dwell on the same surface-level stuff. Be ambitious about what you want to make.

Pick an idea that’s ambitious enough that you find yourself doubting whether you’re going to be able to finish it through. I find that that’s the best way to go deep on interesting ideas in software.

## On leetcode and friends

Leetcode is good for one thing: preparing you for whiteboard/live coding interviews during interviews to large/medium tech companies. I think it helps you become better about thinking through tricky algorithms problems you get at interviews at companies like Uber / Google / Stripe, etc. But it’s not useful for much else. Unless your chief goal your freshman year in college is to ace a Google interview, I’d put Leetcode in the bucket of things that you should look at down the line. You certainly shouldn’t forget about them, but I personally look at Leet code and friends with some disdain.

If you do want to get better at the technical things, here’s what I recommend, from personal experience.

### 1. Read articles, watch talks.

Books are great, and some books are really timeless in software. Some great examples are _The C Programming Language_ by K&R and _The Art of Computer Programming_ by Knuth. If you want to learn a specific modern stack, O’Reilly books are decent and relatively up-to-date.

But beyond timeless classics in the industry, software moves too quickly to be documented well in published print books. At least 60% of what I know about JavaScript, which is my bread and butter, comes from reading blog posts, opinions, and watching talks about JS and web development from recent conferences and events. (The other 40% comes from having written a lot of very terrible JavaScript, which eventually became occasionally terrible JavaScript.)

I would try to find how you best learn — is it reading blog posts? Watching talks? Following people on Twitter? Whatever it is, find people who are making breakthroughs in whatever you want to be great at, and listen to what they’re talking about. Stick through, because these things take time.

### 2. Read source code.

In nearly every technology / language / stack, there are at least one or two codebases everyone points to for high standard of quality. In the C / systems world, this includes the Linux kernel and things like Postgres / Nginx. In JavaScript, I find React an excellent example. I also hear great things about TensorFlow and Folly, Facebook’s C++ library. [There’s a fuller list here](https://medium.com/@markpapadakis/interesting-codebases-159fec5a8cc).

The first time someone told me to start reading other people’s code, I didn’t know where to start or why it would be helpful. In retrospect, I think a more helpful thing to say would be “understand how good software solves hard problems.” If you’re a Postgres maintainer, how do you tackle the problem of sorting strings very quickly in a database? How does the Linux kernel implement task-scheduling? How does TensorFlow arrive at great API design?

These questions aren’t necessarily _just about code_ — they’re questions about how people think about designing software, but oftentimes reading well-written source code will help you understand these design decisions.

### 3. Collaborate

Working on a side project by yourself is so different from working on production-level systems in teams of experienced engineers, and I learn lots from working within a team, both technically and about how to make useful things with code.

There’s really no substitute here — if you have a chance to work in a good team of engineers, take it and see what you make of it.

### 4. Put your work out there and talk about it

If you do this consistently, you’ll find other interesting people making interesting things with technology. And that’s how lots of awesome things start.

### 5. Find interesting problems

When I was in a similar place to you a few years ago, I met a Google engineer at I/O and asked him a similar question. One of the things he said was to find novel and challenging problems, and try to solve them. I didn’t fully appreciate what he meant here then, but I think I know a little better today.

It’s easy to read “challenging problems” in the context of CS as research topics like advanced algorithms or networking. But in reality, software engineering has become such a complex discipline that there’s lots of low-hanging fruit everywhere — problems you can try to tackle even if you don’t know what a [Merkle Tree](https://en.wikipedia.org/wiki/Merkle_tree) is or how the [Raft algorithm](https://raft.github.io/) works.

I can’t tell you what these problems are, because I don’t know yet what topics or technologies you’ll really dive deep into, but there’s interesting problems everywhere, and many of them are recent enough that we’re constantly discovering better solutions. If you so fancy, I’d encourage you to keep an eye out for these types of “unsolved” or incomplete problems in software and see if you can try novel ideas.

### 6. Talk to strangers on the internet

There are many, many very smart, very experienced, very opinionated software engineers on the internet, and many of them like helping guide people to what they think are interesting topics to study. If you ever feel lost, don’t forget that the Internet is full of people with opinions about interesting topics to look at next.

## The right way != the best way

You specifically mentioned that you felt that you aren’t learning the “best” or “most efficient” ways of doing things. I wanted to answer this more fully.

My first and primary response here would be what I mentioned above in the “how to go deeper” part of this post. Performance and optimization happens to be one of my personal “deep dive” topics, and the way I learned about e.g. how V8 optimizes JavaScript is exactly what I mentioned above.

But beyond that, one of the things I learned as I transitioned into building more “real-world” software for companies and for libraries/frameworks is that software never exists in vacuum — there’s always a larger context that encompasses use cases, budgets, timelines, and all sorts of extraneous things like legacy code and availability requirements. As you write more and more code, your primary task is going to be less _writing code_ and more _balancing priorities_. Writing good software is less doing something the _best_ way and more finding one particular solution among many potential ones that balance priorities well. This is to say, I think it’s easy to get sucked into the feeling that performance and efficiency and algorithmic complexity rule software engineering, but people write slow code all the time.

What’s more important, as I’m also learning, seems to be domain-specific unsolved problems, like “how do we make websites load faster, when more people than ever in developing countries are getting online with slower connections?” or “how can we use hardware on battery-constrained devices like smartphones to render water textures more realistically in games?” These aren’t simple, make-this-loop-run-faster problems; these are much more complex and interesting, and this is where a lot of the really great software work happen these days. And how do you get started on finding and thinking about these problems? Explore for breadth, then go deep on a few things.

The last thing I want to impress on you — you may know this already, but I didn’t when I was where you are — is that college gives you a lot of time, and it’s also the _best_ time to search for new interests. So don’t run the risk of thinking what you’re into now will be what you’ll be into for the next many years. Give yourself some space to look for and get hooked on other interests.

I’ve spoken in super general terms here because I don’t want to bias you or influence the kinds of examples I gave, but if your interests are aligned in particular with web and UI development, I’ve always got more to say — so feel free to follow up.
