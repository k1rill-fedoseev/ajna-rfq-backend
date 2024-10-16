# Ajna RFQ orders backend

PoC backend implementation for tracking orders submitted to AjnaRFQ
contracts (https://github.com/k1rill-fedoseev/ajna-rfq-contracts)

## API spec

### GET /api/v1/{chainId}/orders

Fetches list of RFQ orders from the deployment on the particular chain ID. API will return a paginated list, first page
will contain 10 most recent orders that fit all filter criteria, next page can be fetched by setting a `createdBefore`
query param.

| Query Param     | Description                                                                                                                 | Default value |
|-----------------|-----------------------------------------------------------------------------------------------------------------------------|---------------|
| `lpOrder`       | `true` for fetching lp->quote orders, `false` for inverse orders                                                            | `true`        |
| `active`        | `true` to fetch only active orders (not yet filled, not expired, non-zero maker allowance), `false` for fetching all orders | `true`        |
| `maker`         | Optional address of the `maker` to filter orders for the particular `maker`                                                 | N/A           |
| `taker`         | Optional address of the `taker` to filter orders for the particular `taker`                                                 | N/A           |
| `pool`          | Optional list of addresses of Ajna pools to filter orders for, can be specified multiple times                              | N/A           |
| `createdBefore` | Optional pagination timestamp for fetching next page                                                                        | N/A           |

#### Example Response

```json5
{
  "items": [
    {
      "chainId": "8453",
      "rfq": "0xb6a68453b6509173836c20b4bcf66c139ca5ca3f",
      "hash": "0x9e6bc3e903cfbecd848629b11783031a2b36b2344a4b4c3d370b6f02cf6ef0f2",
      "remainingMakeAmount": "1045428212896713652",
      "approvedAmount": "1045428212896713652",
      "createdAt": 1729120027,
      "updatedAt": 1729510136,
      "signature": "0xd93de85a734abeb7954f629fd80428dbbe4e3615aa0543a8efbc7dc1cf3495660d8882562562c67d33d96718779e9368c3eb7b20dcc54747abe2ceb20bd204bf1b",
      "lpOrder": true,
      "maker": "0xd768ac37fe5a4c64121462fc98205a7129c1198a",
      "pool": "0x0b17159f2486f669a1f930926638008e2ccb4287",
      "index": 2619,
      "makeAmount": "1045428212896713652",
      "minMakeAmount": "1045428212896713652",
      "expiry": "2000000000",
      "price": "990000000000000000"
    },
    // ...
  ],
  // will be missing if current page is the last one
  "nextPage": {
    "createdBefore": 1729121678
  }
}
```

### POST /api/v1/{chainId}/orders

Submits signed RFQ order to the off-chain database. Submitted order will have to pass basic validation, order signature
should be valid, maker should have required balance and approval setup beforehand.

#### Request

```json5
{
  "lpOrder": true,
  "maker": "0xd768ac37fe5a4c64121462fc98205a7129c1198a",
  "pool": "0x0b17159f2486f669a1f930926638008e2ccb4287",
  "index": 2619,
  "makeAmount": "1045428212896713652",
  "minMakeAmount": "1045428212896713652",
  "expiry": "2000000000",
  "price": "990000000000000000",
  "signature": "0xd93de85a734abeb7954f629fd80428dbbe4e3615aa0543a8efbc7dc1cf3495660d8882562562c67d33d96718779e9368c3eb7b20dcc54747abe2ceb20bd204bf1b"
}
```

## Deployment info

Hosted version of the API is available at https://ajna-rfq.xyz