A simple lambda authorizer written in go with dynamodb as backend.

This repo seprates code based upon branches , ref branches for specific implementation


Branch Main:
    Go code is seprated into seprate directories and also introduces support for multiple db by using a interface based approach
    Also provides a methods to create custom authorizers


Branch SimpleAndFast:
    Single minimal function tightly integrated with aws dynamodb and apigateway
    You can fork the repo and create your custom logic suitable for your needs
    Expects minimal input as env variables
        Create Zip:
            ```bash
            GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap -tags lambda.norpc
            zip myFunction.zip bootstrap    
    
            ```