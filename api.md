# Experimental

## PaymentFlows

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#PaymentFlow">PaymentFlow</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowNewResponse">ExperimentalPaymentFlowNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowGetResponse">ExperimentalPaymentFlowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowSearchResponse">ExperimentalPaymentFlowSearchResponse</a>

Methods:

- <code title="post /payment-flows">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowNewParams">ExperimentalPaymentFlowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowNewResponse">ExperimentalPaymentFlowNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment-flows/{payment_flow_ref}">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentFlowRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowGetResponse">ExperimentalPaymentFlowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment-flows/search">client.Experimental.PaymentFlows.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowSearchParams">ExperimentalPaymentFlowSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentFlowSearchResponse">ExperimentalPaymentFlowSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentGetResponse">ExperimentalPaymentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentSearchResponse">ExperimentalPaymentSearchResponse</a>

Methods:

- <code title="get /payments/{payment_ref}">client.Experimental.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentGetResponse">ExperimentalPaymentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/search">client.Experimental.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentSearchParams">ExperimentalPaymentSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExperimentalPaymentSearchResponse">ExperimentalPaymentSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExternalAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccount">ExternalAccount</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountNewResponse">ExternalAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountListResponse">ExternalAccountListResponse</a>

Methods:

- <code title="post /external-accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountNewParams">ExternalAccountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountNewResponse">ExternalAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /external-accounts">client.ExternalAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ExternalAccountListResponse">ExternalAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#Invoice">Invoice</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceNewResponse">InvoiceNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceGetResponse">InvoiceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceUpdateResponse">InvoiceUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceListResponse">InvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceListHistoryResponse">InvoiceListHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceSearchResponse">InvoiceSearchResponse</a>

Methods:

- <code title="post /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceNewParams">InvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceNewResponse">InvoiceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices/{id}">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceGetResponse">InvoiceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /invoices/{id}">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceUpdateParams">InvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceUpdateResponse">InvoiceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceListResponse">InvoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices/{id}/history">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.ListHistory">ListHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceListHistoryResponse">InvoiceListHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /invoices/search">client.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceSearchParams">InvoiceSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#InvoiceSearchResponse">InvoiceSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductNewResponse">ProductNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductGetResponse">ProductGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductNewParams">ProductNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductNewResponse">ProductNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{code}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, code <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductGetResponse">ProductGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Roles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#Role">Role</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleNewResponse">RoleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleListResponse">RoleListResponse</a>

Methods:

- <code title="post /roles">client.Roles.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleNewParams">RoleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleNewResponse">RoleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /roles">client.Roles.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#RoleListResponse">RoleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionNewResponse">TransactionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionGetResponse">TransactionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionUpdateResponse">TransactionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionListResponse">TransactionListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionListHistoryResponse">TransactionListHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchResponse">TransactionSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchAllocationsResponse">TransactionSearchAllocationsResponse</a>

Methods:

- <code title="post /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionNewParams">TransactionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionNewResponse">TransactionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/{transaction_ref}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionGetResponse">TransactionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /transactions/{transaction_ref}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionUpdateParams">TransactionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionUpdateResponse">TransactionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/{transaction_ref}/history">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.ListHistory">ListHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionRef <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionListHistoryResponse">TransactionListHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/search">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchParams">TransactionSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchResponse">TransactionSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/allocations/search">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionService.SearchAllocations">SearchAllocations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchAllocationsParams">TransactionSearchAllocationsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#TransactionSearchAllocationsResponse">TransactionSearchAllocationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#User">User</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserNewResponse">UserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserListResponse">UserListResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserNewResponse">UserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go">fragment</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/fragment-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
