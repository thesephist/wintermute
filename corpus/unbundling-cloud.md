---
title: "Unbundling cloud"
date: 2020-05-13T20:35:27-04:00
location: "West Lafayette, IN"
---

**The cloud compute industry is integrating vertically, and unbundling horizontally.**

Once upon a time, cloud compute was homogeneous. You clicked a few buttons on a website, and the host leased you access to a Linux virtual machine of the desired specifications. One VM of a particular size and shape was identical to another, even from a different provider, give or take the margin of error.

Cloud infrastructure has changed a lot in the last few years, and we've seen innovation up and down the technical stack. A single piece that used to be homogeneous is now heterogeneous, differentiated in a myriad of ways, and integrating vertically within workflows. I think this is a trend that'll continue, and a trend that will be difficult for existing Infrastructure-as-a-Service (IaaS) providers like Google and Amazon to reverse, despite their current market dominance.

<p>
<img src="/img/unbundling-cloud.jpg" alt="Unbundling cloud" class="blend-multiply">
</p>

## Horizontal unbundling

As companies build more of the global economy on public cloud infrastructure, previously small niche use cases and problem domains become relevant enough today, to warrant independent products of their own. As these markets grow larger, they attract the attention of both the incumbents (Amazon, Google, Microsoft) and founders of new startups. When the latter of these two are successful over the slower-moving or slower-to-market incumbents, that niche becomes unbundled from the big cloud trinity.

### Static hosting

I think the most mature of these innovative sub-markets as of early 2020 is static web hosting. AWS has their S3 service, which is still the most popular, unoffensive choice for static hosting on the web. But new companies are already pulling market share away from S3 by targeting smaller, more agile teams, with better developer experiences around their product for specific use cases and workflows.

[Vercel](https://vercel.com) (recently rebranded from Zeit) and [Netlify](https://www.netlify.com) lead this area. S3 behaves like a bare-bones utility on top of which developers need to build their own integrations. Actions like deploying new changes to static sites and managing how the S3 assets connect to the rest of their infrastructure requires custom work. By contrast, products from Vercel and Netlify specialize in making the experience of static deploys as easy as possible. They integrate with source control systems like Git and GitHub, and automate setup processes for the required infrastructure. They also connect to tools from the rest of the static web ecosystem to bring the experience much closer to "one-click" than S3 has ever been, and will ever be. This is the virtue of specialization. You can be much, much better in a narrower domain with more specific best practices and workflows.

That Vercel and Netlify are stealing this market from S3 isn't a sign of AWS's decline or inferiority. AWS provides generic infrastructure, and smaller players specialize on top, adding more domain-specific value and reaping commensurate returns. I expect we'll see more players build more specific innovations atop generic cloud infrastructure to unbundle incumbents' general use cases and create new lucrative markets from niches.

### Differentiated hardware

Compute hardware has become more heterogeneous in the last few years. The annual performance and efficiency gains in general-purpose processors are slowing down, and industry focus is shifting to special-purpose processors for specific problems, like cryptography, accelerated machine learning, high end networking, and ray-traced graphics. And the emergence of these niche hardware domains is creating room for new players here, too.

The best-known "differentiated hardware" in the cloud is probably [Google TPUs](https://cloud.google.com/tpu). TPUs are hardware designed for accelerating machine learning workloads. They're faster and more efficient than running training and inference using conventional hardware, and available exclusively on Google's cloud at a price premium.

While the machine learning market is big enough, and Google is enough of a market leader that Google's deep pockets and investments paid off in developing custom hardware, most smaller markets won't get the same treatment. There are dozens of much smaller, emerging hardware niches for specialized use cases that'll open up opportunities for new players to come in and unbundle incumbent services with custom hardware. I'm keeping an eye on growing software markets that benefit from specialized hardware for common computations, like cryptography, video encoding/decoding, specific machine learning and networking tasks, and perhaps even hardware optimized for things like low power usage and high assurance (e.g. finance, healthcare).

In fact, there's developments happening today in many of these fields. Hardware accelerated video codecs are commonplace, but a service optimized for such things may attract the attention of companies that encode at high volumes like Netflix and Mux. Amazon's recently announced [Graviton2](https://aws.amazon.com/blogs/aws/new-m6g-ec2-instances-powered-by-arm-based-aws-graviton2/) chip, provides better safety with encrypted memory. Gaming over the Internet, like Google Stadia, are most efficient when paired with specific gaming hardware rather than conventional "server" processors. And of course, many early-stage companies today are working on dedicated hardware for machine learning and cryptographic applications like cryptocurrencies, for availability in the cloud.

### Managed databases

The last emerging niche I'll mention here is the managed database space. This market is more mature, and already stands on its own. Each of the big cloud players have several services offering fully managed databases on their infrastructure (like AWS's Dynamo and Google's Cloud SQL), but for those who need more specific requirements or prefer to use another kind of a database, new players are already carving out their places in the market. [Datomic](https://www.datomic.com), for example, offers an IaaS version of their proprietary database on the cloud, and uses that software to differentiate.

For developers in markets with specific or unusual shapes of data to store, like high resolution geography, medical and financial information, or high-throughput time series data, I think there's ample room for new players to come in and carve out their places in the managed database space.

## Vertical integration

The unbundling of a homogeneous cloud by smaller innovators comes hand-in-hand with tighter, more domain specific integrations vertically through the technical stack, from hardware and firmware up to programming languages and libraries used by developers.

I think the machine learning space has one of the best-developed vertically integrated technical stacks today. Google builds the hardware (Cloud TPUs), the service (Google Cloud, AutoML), and the libraries and models (TensorFlow, Keras) that lots of developers use to build and deliver machine learning models and inferences. What would have been small differentiators in each of the individual layers of the stack are now pieces of a stronger moat because of the tighter integration. Some companies even use Amazon's web services, but Google's infrastructure purely for their machine learning workloads -- Google Cloud ML products are effectively its own niche and product, and shows the potential in vertically integrated products.

I think Google Cloud, one of the three old guards of the industry, is a bit of an anti-trend borne of the popularity that machine learning holds today. So I'll also note two other cases where smaller companies are offering vertically integrated cloud products.

[Shopify](https://www.shopify.com) and [Stripe](https://stripe.com/connect) probably aren't the first examples that come to mind when you think of cloud infrastructure, but they compete in the same market as many of the incumbents' cloud offerings. If you have an online shop, you can either (1) build and deploy a website on the generic AWS EC2 platform, or (2) just create a Shopify site and connect Stripe, to have a fully-managed e-commerce store. Although their novelty might be wearing off, they've also unbundled the functionality of old web hosts -- hosting simple websites, often used for shops -- and built a vertically integrated product for that specific use case and market. They're the go-to solution for launching an e-commerce store in 2020.

[Mux](https://mux.com) specializes in video delivery and live video streaming. They don't just provide managed hardware like AWS or GCP, but also provide libraries and software integrations that do much of the heavy lifting for developers wanting to build streaming-video experiences on the web.

There's also technologies that are just emerging that might offer even closer integration and lock-in between developer tooling and infrastructure. [Cloudflare Workers](https://workers.cloudflare.com) offers JavaScript and WebAssembly tools to deploy dynamic code to the edges of their CDN with their API, and [Dark](https://darklang.com) promises to offer continuous delivery and effectively-zero downtime with a programming language and version control environment that's tailor-made for their infrastructure.

Startups that are unbundling old or existing general-purpose cloud hosts are integrating more tightly at the same time, and I think this trend will continue as smaller cloud markets are played out.

## The heterogeneous cloud

Differentiation and unbundling signals a maturing industry, and I think cloud infrastructure has come of age. Startups and specialized teams in existing companies are playing to their strengths, building more vertically-integrated products that offer better experiences in small, growing niches. And consequently, they're stealing smaller, faster-moving markets away from large incumbents (or in some cases, building on top of them but capturing more of the value).

I think we're still very early in this trend, and the specialization of cloud infrastructure is going to make cloud really interesting again, reminiscent of the explosion of variety and experimentation in the "web 2.0" era. There's still a lot of value to be captured by building more domain-specific cloud infrastructure, and I'm excited to watch the space grow, both as a developer and a consumer.
