# Rancher vs OpenShift

>Choose one service provider such as Rancher and compare it to another such as OpenShift.
>
>Decide arbitrarily which service provider is "better" and argue for it against the other service provider.
>
>For the submission a bullet point list is enough.

I have arbitrarily decided (for the purpose of this exercise) that Rancher is "better" because:

- OpenShift appears to be much more heavyweight being a Kubernetes distribution. This seems to have a negative effect on things such as installation time, upgrade time.
- Rancher aims more at doing just enough to simplify complex Kubernetes operations without tying your hands too much. OpenShift seems to ship as a platform that you are locked into fully after choosing to go with it. With Rancher you can test it out and get rid of it if it's not what you need or even import your current clusters to it.
- I've understood that OpenShift was originally developed to be a proprietary PaaS solution that was based on Linux containers that only later started to use Docker and Kubernetes as it's container and container orchestration technology. Rancher on the other hand (as I understand it) was specifically built for making Kubernetes management easier anywhere (in the cloud, edge or locally). 
- Because OpenShift is a more comprehensive solution it also lives in it's own space in a way. Learning to use OpenShift will not necessarily transfer to anything outside of OpenShift.
- OpenShift seems to also be much more expensive than Rancher.
- Rancher is fully open source (in contrast to OpenShift's Enterprise version)