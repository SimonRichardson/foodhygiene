# Service

This service folder contains a Service interface, which helps abstract the
following in a very nice and modular way:

 1. Caching
 2. Mock testing

### Caching

The caching is a very simple cache, but helps improve manual testing when using
the UI. It's not in anyway clever and just stores the data during the
applications life cycle. Once the application is closed, all the data with in
the application is released.

### Mock testing

Mock testing service helps test various parts of the system without the need to
hit the real endpoint. The mocking is provided via gomock and generates the
`mock_service/service.go` file using mockgen.

See examples of this in `pkg/query/api_test.go`
