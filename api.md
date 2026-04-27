# Experimental

## PaymentFlows

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#PaymentFlow">PaymentFlow</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowNewResponse">ExperimentalPaymentFlowNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowGetResponse">ExperimentalPaymentFlowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowSearchResponse">ExperimentalPaymentFlowSearchResponse</a>

Methods:

- <code title="post /payment-flows">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowNewParams">ExperimentalPaymentFlowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowNewResponse">ExperimentalPaymentFlowNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment-flows/{payment_flow_ref}">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentFlowRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowGetResponse">ExperimentalPaymentFlowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment-flows/search">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowSearchParams">ExperimentalPaymentFlowSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentFlowSearchResponse">ExperimentalPaymentFlowSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentGetResponse">ExperimentalPaymentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentSearchResponse">ExperimentalPaymentSearchResponse</a>

Methods:

- <code title="get /payments/{payment_ref}">client.Experimental.Payments.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentGetResponse">ExperimentalPaymentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/search">client.Experimental.Payments.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentSearchParams">ExperimentalPaymentSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExperimentalPaymentSearchResponse">ExperimentalPaymentSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccount">ExternalAccount</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountNewResponse">ExternalAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountListResponse">ExternalAccountListResponse</a>

Methods:

- <code title="post /external-accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountNewParams">ExternalAccountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountNewResponse">ExternalAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external-accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ExternalAccountListResponse">ExternalAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#Invoice">Invoice</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceNewResponse">InvoiceNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceGetResponse">InvoiceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceUpdateResponse">InvoiceUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceListResponse">InvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceListHistoryResponse">InvoiceListHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceSearchResponse">InvoiceSearchResponse</a>

Methods:

- <code title="post /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceNewParams">InvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceNewResponse">InvoiceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices/{id}">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceGetResponse">InvoiceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoices/{id}">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceUpdateParams">InvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceUpdateResponse">InvoiceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceListResponse">InvoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices/{id}/history">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.ListHistory">ListHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceListHistoryResponse">InvoiceListHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /invoices/search">client.Invoices.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceSearchParams">InvoiceSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#InvoiceSearchResponse">InvoiceSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductNewResponse">ProductNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductGetResponse">ProductGetResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductNewParams">ProductNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductNewResponse">ProductNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{code}">client.Products.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, code <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductGetResponse">ProductGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Roles

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#Role">Role</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleNewResponse">RoleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleListResponse">RoleListResponse</a>

Methods:

- <code title="post /roles">client.Roles.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleNewParams">RoleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleNewResponse">RoleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /roles">client.Roles.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#RoleListResponse">RoleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionNewResponse">TransactionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionGetResponse">TransactionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionUpdateResponse">TransactionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionListResponse">TransactionListResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionListHistoryResponse">TransactionListHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchResponse">TransactionSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchAllocationsResponse">TransactionSearchAllocationsResponse</a>

Methods:

- <code title="post /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionNewParams">TransactionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionNewResponse">TransactionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/{transaction_ref}">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionGetResponse">TransactionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /transactions/{transaction_ref}">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionUpdateParams">TransactionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionUpdateResponse">TransactionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/{transaction_ref}/history">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.ListHistory">ListHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionListHistoryResponse">TransactionListHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/search">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchParams">TransactionSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchResponse">TransactionSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/allocations/search">client.Transactions.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionService.SearchAllocations">SearchAllocations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchAllocationsParams">TransactionSearchAllocationsParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#TransactionSearchAllocationsResponse">TransactionSearchAllocationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#User">User</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserNewResponse">UserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserListResponse">UserListResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserNewResponse">UserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/fragment-dev/golang">fragment</a>.<a href="https://pkg.go.dev/github.com/fragment-dev/golang#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
