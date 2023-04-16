# Backend for portfolio website

- Allows retrieval of projects (synced daily from Github, as well as custom ones without repositories).
- Defines demonstration implementations of projects.
- Accepts email requests.
- Organises payment flows for donations.

- Built entirely using AWS Lambda functions, sitting behind an API Gateway.


<br> </br>
## Todo:
- [ ] Use terraform to create lambdas that don't already exist.
- [ ] Setup Stripe flow.
- [ ] Setup email flow.

<br></br>
# Structure

`cmd/`

- Contains all lambda handlers

`cmd/projects`

- All handlers related to projects.

`cmd/payments`

- All handlers related to Stripe payments.

`cmd/emails`

- All handlers related to receiving emails.

`internal/`

- Contains all integration code.

`internal/dynamo`

- Integration for DynamoDB.

`internal/stripe`

- Integration for Stripe.

`pkg/`

- General library code.