bowdata.test.go_tcp_echo
=========================




Introduction
--------------------------


Setting up for development
--------------------------

Clone the repo and checkout a new dev branch. At the terminal run ``make develop`` in the root dir of the repo. The make
recipe assumes you have the ``AZURE_DEVOPS_PAT`` variable set and will use it to modify the git config to authenticate
imports of other Bow Data Go Modules. Then the local module (i.e. the one that is implemented in this repo) has its
references updated in ``go.mod`` so that any locally executed ``go`` commands see the local develop branch and not
``origin:master``. Finally, the cache is cleared and any dependencies of the local module are imported to the local
environment.





DevOps Pipeline
---------------

There is a devops pipeline (``azure-pipelines.yml``) which provides CD/CI via Azure DevOps based on the `Github Flow`
protocol.







------------

Repo created from bowdata.toolkits.repos (version 0.0.1a19)

