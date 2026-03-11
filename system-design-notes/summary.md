
page 25 message queue
page 43 example if you are asked
page 53 Besides the client and server
page 97 Inconsistency resolution: versioning
----------------------------
interview tips from this book:



Write down your assumptions. It is a good idea to write down your assumptions to be
referenced later.

Over-engineering is a real disease of many
engineers as they delight in design purity and ignore tradeoffs



Other red flags include narrow mindedness, stubbornness, etc.

Step 1 : understand the problem and establish dessign scope
DON'T be like Jimmy.
In a system design interview, giving out an answer quickly without thinking gives you no
bonus points.

So, do not jump right in to give a solution. Slow down. Think deeply and ask questions to
clarify requirements and assumptions. This is extremely important.

When you ask a question, the interviewer either answers your question directly or asks you to
make your assumptions. If the latter happens, write down your assumptions on the
whiteboard or paper. You might need them later.

What kind of questions to ask? Ask questions to understand the exact requirements. Here is a
list of questions to help you get started:
• What specific features are we going to build?
• How many users does the product have?
• How fast does the company anticipate to scale up? What are the anticipated scales in 3
months, 6 months, and a year?
• What is the company’s technology stack? What existing services you might leverage to
simplify the design?

Step 2: Porpose high level design and get buy in
initial blue print of the design 
draw componnents 
do back of envelope 
back-of-the-envelope calculations are estimates you create using a
combination of thought experiments and common performance numbers to get a good feel for
which designs will meet your requirements
communicate with interviewer
get feedback and buy in

Step 3: Design deep dive
at this step:
agreed on the overall goals and feature scope
sketched out a high level overall design
feedback 
initial ideas

Step 4: wrap up
The interviewer might want you to identify the system bottlenecks.
Never say your design is perfect
Error Cases are interesting to talk about.
operation issues are worth mentioning. monitor metrics and error logs. how to roll out the system.
how to handle the next scale curve
if you had more time, propose other refinements.


Dos
• Always ask for clarification. Do not assume your assumption is correct.
• Understand the requirements of the problem.
• There is neither the right answer nor the best answer. A solution designed to solve the
problems of a young startup is different from that of an established company with millions
of users. Make sure you understand the requirements.
• Let the interviewer know what you are thinking. Communicate with your interview.
• Suggest multiple approaches if possible.
• Once you agree with your interviewer on the blueprint, go into details on each
component. Design the most critical components first.
• Bounce ideas off the interviewer. A good interviewer works with you as a teammate.
• Never give up.

Don’ts
• Don't be unprepared for typical interview questions.
• Don’t jump into a solution without clarifying the requirements and assumptions.
• Don’t go into too much detail on a single component in the beginning. Give the highlevel design first then drills down.
• If you get stuck, don't hesitate to ask for hints.
• Again, communicate. Don't think in silence.
• Don’t think your interview is done once you give the design. You are not done until your
interviewer says you are done. Ask for feedback early and often.

time allocation on each step
Step 1 Understand the problem and establish design scope: 3 - 10 minutes
Step 2 Propose high-level design and get buy-in: 10 - 15 minutes
Step 3 Design deep dive: 10 - 25 minutes
Step 4 Wrap: 3 - 5 minutes


--------------------------------






NoSQL databases are grouped into four categories: key-value stores, graph stores, column stores, and document stores. 
Vertical scaling, referred to as “scale up”, means the process of adding more power (CPU,RAM, etc.) to your servers.
Horizontal scaling, referred to as “scale-out”, allows you to scale by adding more servers into your pool of resources.



A service level agreement (SLA) is a commonly used term for service providers. This is an
agreement between you (the service provider) and your customer, and this agreement
formally defines the level of uptime your service will deliver.

worth looking at this:
like multi-masters and circular replication
celebrity problem , hotspot key problem 

---------------------

chapter 4: design a rate limiter

in microservices nowadays rate limiting usually implemented within a component called API Gateway.
API Gateway is a fully managed service that supports rate limiting, ssl termination authentication ip whitelisting servicing static content . 
API Gateway is a middleware that supports rate limiting.
make sure your cuurent programming language is efficient to implement rate limiting on the server-side.
building your own api limiter service takes time. if you dont have enough engineering resources to implement a rate limiter. a commercial one is a better option.

rate limiting algorithms:
- token bucket. { bucket size, refil rate}
- leaking bucket. if FIFO . { bucket size, outflow rate}
- fixed window counter
- sliding window log
- sliding window counter

for high level architeture:
keep tracks data should be stored in a in-memory cache and also support time-based expiration. Redis is a popular one. (INCR, Expire)

• Avoid being rate limited. Design your client with best practices:
• Use client cache to avoid making frequent API calls.
• Understand the limit and do not send too many requests in a short time frame.
• Include code to catch exceptions or errors so your client can gracefully recover from
exceptions.
• Add sufficient back off time to retry logic.


-------------------------------
Chapter 5:  Design Consistent Hashing

- simple hashing problem is that if one server gets removed, all the data will be relocated.
- consistent hashing we can  minimize the data relocation only to a limited range of data.

the benefits of consistent hashing include:
- minimized keys are redistributed when servers are added or removed.
- it is easy to scale horiziontally because data are more evenly distributed.
- mitigate hotspot key problem. Excessive access to a specific shard could cause server overload. imagine data for katty perry, justin bieber and lady gaga all end up on the same shard. consistent hashing helps to mitigate the probelm by distributing the data more evenly.

consistent hashing read world :
partitioning component of Amazon's dynamo database
data partitioning across the cluster in apache cassandra 
discord chat application
Akamai content delivery network. 
Maglev network load balancer.

virtual nodes will be added to evenly distribute the data across more servers to reduce the imbalanced data distribution.

-------------------------
Chapter 6: design a key-value store

understand the problem and establish design scope:

for a single server key-value store:
    sotre pairs in a hash table
    everything in memory is fast but impossible for all the data
    two optimizations can be done:
        data compression
        store only frequently used data in memory and the rest on the disk.

distributed key-value store:
    distributed hash table

CAP theorem:
    Consistency: consistency means all clients see the same data at the same time no matter
which node they connect to

Availability: availability means any client which requests data gets a response even if some
of the nodes are down.

Partition Tolerance: a partition indicates a communication break between two nodes.
Partition tolerance means the system continues to operate despite network partitions.

CP (consistency and partition tolerance) systems: a CP key-value store supports
consistency and partition tolerance while sacrificing availability.
AP (availability and partition tolerance) systems: an AP key-value store supports
availability and partition tolerance while sacrificing consistency.

Core components to build a key-value store:
- Data Partition
- Data Replication
- Consistency
- Inconsistency resolution
- Handling Failures
- System Architecture Diagram
- Write Path
- Read Path


Consistency Models:
Strong consistency: any read operation returns a value corresponding to the result of the
most updated write data item. A client never sees out-of-date data.

Weak consistency: subsequent read operations may not see the most updated value.

Eventual consistency: this is a specific form of weak consistency. Given enough time, all
updates are propagated, and all replicas are consistent.


vector clock for inconsistency solving
Bloom filter
SSTables
Twitter’s unique ID generation system called “snowflake” 


"https://www.allthingsdistributed.com/2026/02/a-chat-with-byron-cook-on-automated-reasoning-and-trust-in-ai-systems.html"


URL Frontier page 136