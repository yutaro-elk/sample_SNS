create/tweet:
	aws dynamodb create-table --cli-input-json file://./dynamodb/schema/create_tweets_table.json --endpoint-url http://localhost:8000

create/user:
	aws dynamodb create-table --cli-input-json file://./dynamodb/schema/create_users_table.json --endpoint-url http://localhost:8000

create/sequence:
	aws dynamodb create-table --cli-input-json file://./dynamodb/schema/create_sequences_table.json --endpoint-url http://localhost:8000

create/timeline:
	aws dynamodb create-table --cli-input-json file://./dynamodb/schema/create_timeline_table.json --endpoint-url http://localhost:8000

create/tables:
	make create/tweet
	make create/user
	make create/sequence
	make create/timeline

create/env:
	echo DYNAMO_ENDPOINT=http://localhost:8000 > ./api-server/.env

build/lambda:
	cd api-server/cmd/lambda && GOOS=linux go build -o ../../../bin/lambda

deploy:
	sls deploy --aws-profile test

