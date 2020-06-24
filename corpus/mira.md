---
title: "Mira v0: my personal people notes app"
date: 2020-03-28T07:19:33-04:00
location: "West Lafayette, IN"
---

Last night, I built and started using [Mira](https://github.com/thesephist/mira), a sort of a hybrid contacts-app-CRM that's designed for my workflow around the people in my work and life. I'm calling it a "people notes" app.

![Mira on desktop](/img/mira.png)

I spend a lot of time trying to remember and keep track of people in my life -- what we talked about, what people are working on, what they're looking for, where they're trying to go. This was also a big part of my [goals in 2019](/posts/2019-goals/). Until today, I've used a few pages from [Ligature](https://github.com/thesephist/polyx), my generic text notes app, to keep track of people. But as I've built up a routine, I'm hitting a point where I'd really benefit from a customized tool to keep track of just the info I need in a searchable and organized way.

Hence, Mira was born. It's open-source and licensed under the permissive MIT License; you can go check it out on [GitHub](https://github.com/thesephist/mira).

## The timeline of a relationship

_But Linus, why don't you just use a normal contacts app or CRM like a normal person?_ I hear you asking.

First, for the same reason I wrote my own programming language and JavaScript UI framework and notes app -- because it's fun. But second, what I need is a tool for keeping track of my relationships with people that can handle more detail and history than a typical contacts app that comes pre-installed on a phone, but is not quite as fully featured as a CRM designed for sales teams and superconnectors.

Specifically, beyond the basics like phone numbers and emails, I needed my contacts app to be able to help me remember my _history_ with that person, not just how to get in touch with them. The most useful parts of what I jot down about people are topical or anecdotal. It's things like where we first met, what we talked about, what they were working on or what problem they were trying to solve when I last talked to them. I needed a tool designed around meetings, and connecting those conversations together and leading me to new conversations, rather than just getting and keeping in touch with people.

So Mira is going to be designed around building a story around each person in the list, through lists of conversations and topics, more detailed notes about people, and anything else I find myself needing along the way. There may be other apps that do this -- but this is mine, and it works for me. I quite enjoy the process of building custom tools like this.

## Mira is a rough draft

Mira was quickly designed and built around a few loose patterns I noticed myself using over the last year, but it still has lots of room for improvement, and my workflow is still growing and changing. So I'm calling the current version of Mira "v0" -- it's really just a quick, rough draft of a people notes app that I can use to nail down a better workflow and iterate on quickly over the next few months.

As a result, the architecture of the app is optimized for fast and painless iteration over safety or other cross-cutting software concerns. I can add and modify what kinds of fields each contact card stores, in 1-2 lines of code, live-reloading.

Once I find a more stable workflow that I can depend on with Mira, I'll probably rewrite it into the [Polyx](https://github.com/thesephist/polyx) software suite which also contains my notes app. For this reason, I'm not too concerned about the "scalability" of this particular piece of code or anything of that sort.

## Made for quick iteration

Although Mira isn't anything novel from a technical perspective, there's one aspect of the software design that I'm proud of, and it's how quickly I can iterate on the schema of each contact item to add new fields. It takes a single line of code added to the entire codebase to fully support a new field for a contact item, from data storage to editing UI.

The web app, which drives all interactions, uses a `Contact` record class as the source of truth for schema:

```
class Contact extends Record {

    singleProperties() {
        return [
            // format:
            // [
            //     label,
            //     column name,
            //     placeholder,
            //     (optional) use <textarea>?,
            // ]

            ['name', 'name', 'name'],
            ['place', 'place', 'place'],
            ['work', 'work', 'work'],
            ['twttr', 'twttr', '@username'],
            ['last', 'last', 'last met...'],
            ['notes', 'notes', 'notes', true],
        ];
    }

    multiProperties() {
        return [
            ['tel', 'tel', 'tel'],
            ['email', 'email', 'email'],
            ['mtg', 'mtg', 'meeting', true],
        ]
    }

}
```

When the UI renders, the app references this single source of truth to create and fill in any table rows in the UI, and dynamically create editable inputs, buttons, and key event callbacks. This means that, given this schema, the interface shown in the screenshot at the start of this post is entirely automatically generated by the UI. As I add new fields to the schema, the UI will automatically adapt to include them. This is the reason I can iterate so quickly on the contacts schema, which is important in this experimental, version-0 phase, as I play with which fields are useful for me, and which aren't.

I'm excited to upgrade my people-notes experience from a long, slow plain text file to something more structured, and keep hacking on this to something I can really find indispensible. As I've written elsewhere, this workflow isn't just a "CRM" type workflow to network. It genuinely helps me remember important things about people I meet for the first time, and people I care about, that I'd otherwise have forgotten. And having a better tool here is going to go a long way towards improving that experience for me.
