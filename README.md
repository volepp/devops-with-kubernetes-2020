# devops-with-kubernetes-2020
Repository for publishing my exercise solutions for the course DevOps with Kubernetes 2020 at University of Helsinki.

## Google Cloud SQL vs DIY Postgres

Up to this point, self-hosted Postgres has been used in previous exercises which of course enables us to proceed with using it with somewhat more of a routine.
This would avoid the overhead that comes from learning Google Cloud SQL in the beginning since I have no experience of using it before. So, considering the cost and the time it would take to initialize the database in the first place, I suspect hosting Postgres by ourselves would be at least a slightly faster way to go.

We have to, however, consider other factors that could influence the decision as well. One of those factors is the maintenance of the database.
Reading the docs for running a Postgres instance with Google Cloud SQL, I have a feeling that after learning it initially, Google Cloud SQL would take a lot of weight off our shoulders in the long run. In this regard I will, with some level of certainty, say that setting up a Postgres instance with Google Cloud SQL would be the wiser choice.

I also think that creating backups of our data will be less bothersome using Google Cloud SQL. Again, I have no prior experience with the service but reading the docs gives me a feeling backing up the data shouldn't be too much of headache (of course I say that now but we'll see if that's the case when the time of setting it up actuallly comes).

In the end there is still one more aspect I think needs to be addressed. Since I've never used Google Cloud SQL, I see this as a learning opportunity.
Of course a case can be made for whether it's helpful to spend a lot of time (or very little, we'll see) learning provider specific technologies while you are not even certain you will stick with them over time.
The way I see it, however, is that there are surely similar services of other providers that can't be too far apart from Google Cloud SQL.
In that way, learning Google Cloud SQL would most likely be helpful when picking up similar services of other providers such as Amazon or Microsoft.
Learning new things never hurts.

These are the considerations over which I base my decision to use Google Cloud SQL hosted Postgres database. It remains to be seen how this turns out in the end.