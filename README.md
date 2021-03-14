# devops-with-kubernetes-2020
Repository for publishing my exercise solutions for the course DevOps with Kubernetes 2020 at University of Helsinki.

## Google Cloud SQL vs DIY Postgres

Up to this point, self-hosted Postgres has been used in previous exercises which of course enables us to proceed with using it with somewhat more of a routine.
This would avoid the overhead that comes from learning Google Cloud SQL in the beginning since I have no experience of using it before. So, considering the cost and the time it would take to initialize the database in the first place, I suspect hosting Postgres by ourselves would be a faster way to go.

We have to, however, consider other factors that could influence the decision as well. One of those factors is the burden of maintaining the database.
Reading the docs for running a Postgres instance with Google Cloud SQL, I have a feeling that after learning it initially, Google Cloud SQL would take a lot of weight off our shoulders in the long run. In this regard I will, with some level of certainty, say that setting up a Postgres instance with Google Cloud SQL would be the wiser choice.
On the other hand, the solution is only used for the remainder of this course. This means that the gain we can get later on from more effortless maintaining might not be all that important overall.

It seems creating backups shouldn't be a big problem with either option.

The choice was not easy but in the end I decided to go with a self-hosted Postgres database. That's because so far that is what I've used during the course and I feel it will save me a lot of time and work in the short term.
Since the scope for this decision is only the two parts left in the course, I think the time saved in the beginning overweighs the possible time we could save by letting Google Cloud SQL manage and maintain the database.
Also it might be the case that it's wiser to learn how to host the database by yourself since you're not guaranteed to always have Google Cloud SQL at your disposal (for example if you have to use a different provider such as Amazon).
Although I suspect other providers offer similar kinds of services that are more or less the same usage wise.
All in all I will try to go with a self-hosted Postgres database for now since it feels like a bit less hassle at the moment.