# ise-t
ISE Certificate toolbox is a small cli utility that performs some basic certificate management tasks for Cisco Identity Services Engine.  It aims to lower the friction level of managing ISE certificates and is designed to be used either standalone or as part of a pipeline.

Basic functionality:

- List nodes in a deployment
- List All system Certificates in a deployment by node
- Return a list of certificates expiring in x days
- One shot disaster recovery export of all certificates

Aspirational: ACME interface to provision certificates from:
    Letsencrypt
    Active Directory


