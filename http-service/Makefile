GOOGLE_CLOUD_PROJECT=argon-acumen-337101
EXTERNAL_IP=35.185.162.62
PORT=8080

deploy:
	skaffold delete
	skaffold run --default-repo=gcr.io/$(GOOGLE_CLOUD_PROJECT)

test:
	curl -v http://$(EXTERNAL_IP):$(PORT)/api/v1/demo