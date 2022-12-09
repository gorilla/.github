## Gorilla Toolkit

> ⚠️ **The Gorilla Toolkit is now in archive-mode, and is no longer actively maintained. You can read more below.**

_We’ll be putting the Gorilla project’s repositories into “archive mode” by the end of 2022._

It’s been a long run. The [first commit](https://github.com/gorilla/websocket/commit/273ecadfcae3306e1943ac6d4160ef02cac48247) on gorilla/mux was back in October 2012, just a few months after [Go itself reached 1.0](https://go.dev/blog/go1). [gorilla/websocket](https://github.com/gorilla/websocket/commit/273ecadfcae3306e1943ac6d4160ef02cac48247) started back in October 2013, and a number of other packages — forming the “Gorilla Toolkit” — sprung up around the same time.

The original author and maintainer, [moraes](https://github.com/moraes), had moved on a long time ago. [kisielk](https://github.com/kisielk) and [garyburd](https://github.com/garyburd) had the longest run, maintaining a mix of the HTTP libraries and gorilla/websocket respectively. I ([elithrar](https://github.com/elithrar)) got involved sometime in 2014 or so, when I noticed kisielk doing a lot of the heavy lifting and wanted to help contribute back to the libraries I was using for a number of personal projects. Since about ~2018 or so, I was the (mostly) sole maintainer of everything but websocket, which is about the same time garyburd put out an (effectively unsuccessful) [call for new maintainers](https://github.com/gorilla/websocket/issues/370) there too. 

Some of you might be reading this and thinking that we didn’t give a fair shot to potential new maintainers, or that the bar for new maintainers was too high. The problem there is two-fold:

* There were no active contributors even triaging issues. The call for maintainers made it clear we’d help merge and do a final review for anyone wanting to start contributing. Instead, a number of folks raised their hands (read: commented in the thread) and then were never seen again. Many OSS projects have a number of casual maintainers: we just never seemed to get anyone to stick. Maybe the “utilitarian” nature of the libraries didn’t help, or maybe it was more compelling to author your own?
* These are widely used libraries. As we said in the original call for maintainers: “no maintainer is better than an adversarial maintainer!” — just handing the reins of even a single software package that has north of 13k unique clones a week (mux) is just not something I’d ever be comfortable with. This has tended to play out poorly with other projects.

The call for maintainers lived well beyond the original 6 month window in an attempt to find someone who could responsibly take over the libraries. We didn’t find that person, people, or company, and here we are today.

I do believe that open source software is entitled to a lifecycle — a beginning, a middle, and an end — and that no project is required to live on forever. That may not make everyone happy, but such is life.

### Was it about money? 

No. I don’t think any of us were after money here. The Gorilla Toolkit was, looking back at the most active maintainers, a passion project. We didn’t want it to be a job. 

This isn’t a dig at maintainers who do want to be paid for their efforts, but a reminder that not everyone does things for money.

### What does “archiving” mean?

It means the repositories go into “read-only” mode ([read more here](https://docs.github.com/en/repositories/archiving-a-github-repository/archiving-repositories)). Anyone still using them can still clone them, go get them, and continue to build projects against them. In effect, there’s really no change here from the last 12 months, and it won’t break existing projects.

What it does signal is that there will be no future development on these libraries.

Folks are welcome to (as they always have been) fork them: all of the Gorilla libraries are permissively licensed (MIT, BSD-3, and Apache 2.0).

Thanks for all the fish,
Matt and Gary

---

🦍 A helpful toolkit for building HTTP-based applications with the [Go programming language](https://go.dev/).

## Projects

A few of the most popular libraries:

* [**mux**](https://github.com/gorilla/mux), a powerful request router for web applications
* [**sessions**](https://github.com/gorilla/sessions), making cookies and session management easy
* [**websocket**](https://github.com/gorilla/websocket), a standards-compliant and widely used websocket library
* [**handlers**](https://github.com/gorilla/handlers), a collection of useful middleware for Go HTTP applications.

## Help 

Open an issue on the relevant project. For security issues, see [SECURITY.md](https://github.com/gorilla/.github/blob/master/SECURITY.md).
