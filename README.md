# Backend for portfolio website

- Allows retrieval of projects (synced daily from Github, as well as custom ones without repositories).
- Defines demonstration implementations of projects.
- Accepts email requests.
- Organises payment flows for donations.

Built entirely using AWS Lambda functions, sitting behind an API Gateway.

Todo:
- [ ] Use terraform to create lambdas that don't already exist.
- [ ] Setup Stripe flow.
- [ ] Setup email flow.