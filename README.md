A simple lambda authorizer written in go with dynamodb as backend.

This repo seprates code based upon branches , ref branches for specific implementation

## Prequites
1. To setup aws lambda and CI/CD we use sam, please download it from https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html

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

    To Use sam template (remember to change <region> to your region and <accountId> to your accountId in template.yaml )
    ```bash
        sam build && sam deploy --guide
    ```
    fill the details and it will create lambda and role for you
    now you can connect it as authorizer in your api gateway
    