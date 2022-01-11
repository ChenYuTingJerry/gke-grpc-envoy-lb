GOOGLE_CLOUD_PROJECT=argon-acumen-337101
EXTERNAL_IP=35.185.162.62
PORT=8081

deploy:
	skaffold delete
	skaffold run --default-repo=gcr.io/$(GOOGLE_CLOUD_PROJECT)

test:
	grpcurl -plaintext -d '{"content": "echo"}' -proto proto/echo.proto \
	$(EXTERNAL_IP):$(PORT) api.EchoService/Echo