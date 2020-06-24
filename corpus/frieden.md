---
title: "Frieden: my personal availability calendar"
date: 2020-06-02T17:45:11-04:00
location: "West Lafayette, IN"
---

Today, I added another small item to my growing arsenal of software tools to make my day-to-day work smoother. I built and launched [Frieden](https://github.com/thesephist/frieden), my personal, public read-only availability calendar. You can see it running at [free.linus.zone](https://free.linus.zone). It pulls my calendar data in real time from my actual personal and work calendars, and aggregates my schedule into a single public view. It looks like this:

<p><img src="/img/frieden.png" alt="Screenshot on desktop" class="blend-multiply"></p>

There are services out there that allow strangers to schedule meetings with me from a limited set of time slots, like Uncalendar or Calendly. I didn't want those -- I want to have full control of who gets to add events on my calendar where, so I still want the booking process to be manual. But when I'm scheduling things onto my calendar, I waste a lot of back-and-forths and everyone's time finding a time slot that's open for everyone, and serializing the empty slots in my own calendars to others through email and text.

So Frieden is designed to alleviate that small pain point. Instead of me texting a friend "I'm free on X day from Y-Z time and at W other hour," they have an always-up-to-date public place to check what times I might be available. I think it'll also help me personally check my free slots whenever I need to, since opening this page is faster than checking my calendar app sometimes.

It's also small and easy to fix and update as needed -- the entire application is under 150 lines of loose Go on the backend and about 400 lines of frontend [Torus](https://github.com/thesephist/torus) UI code. It fits well in my head, and there's few places for things to break, so I can keep it running without much headache. Like most of my web projects, Frieden runs on my 1-vCPU DigitalOcean server.

Frieden is designed to be deployable and customizable for other people to use, with a little configuration under your Google Calendar settings. If you're interested, feel free to get in touch or reference the [README](https://github.com/thesephist/frieden/blob/master/README.md) for the project.

---

PS - Why's it called _Frieden_? Frieden is a German word for "peace" -- as in, the opposite of war and chaos, which is what I want in my schedule. It's also a bit of a pun and a play on words on the English word "free," as in "free time".
