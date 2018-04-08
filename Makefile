
options := -var 'cidr_home="'`curl -s http://checkip.amazonaws.com`/32'"'

.PHONY: apply
apply:
	terraform apply $(options)

.PHONY: destroy
destroy:
	terraform destroy $(options)

ssh-config: terraform.tfstate
	gcloud compute config-ssh --ssh-config-file ./ssh-config

.PHONY: ssh
ssh: ssh-config
	gcloud compute ssh test-gce
