# go build command
gobuild:
	@go build -v -o chatserver cmd/chatserver/*.go

# go run command
gorun:
	make gobuild
	@./chatserver --config_path='./config/tkp-app.{TKPENV}.yaml'

# deploy command
deploy:
	@echo "deploying"