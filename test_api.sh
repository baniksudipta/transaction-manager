#!/bin/bash

BASE_URL="http://localhost:8080"

echo "=== 1. Create an Account ==="
curl -s -w "\n" -X POST $BASE_URL/accounts \
  -H "Content-Type: application/json" \
  -d '{"document_number": "12345678900"}'
echo ""

echo "=== 2. Get the Account ==="
curl -s -w "\n" -X GET $BASE_URL/accounts/1
echo ""

echo "=== 3. Create a Debit Transaction (Normal Purchase) ==="
curl -s -w "\n" -X POST $BASE_URL/transactions \
  -H "Content-Type: application/json" \
  -d '{"account_id": 1, "operation_type_id": 1, "amount": 50.50}'
echo ""

echo "=== 4. Create a Credit Transaction (Credit Voucher) ==="
curl -s -w "\n" -X POST $BASE_URL/transactions \
  -H "Content-Type: application/json" \
  -d '{"account_id": 1, "operation_type_id": 4, "amount": -100.0}'
echo ""

echo "=== 5. Validation Error: Missing Document Number ==="
curl -s -w "\n" -X POST $BASE_URL/accounts \
  -H "Content-Type: application/json" \
  -d '{"document_number": ""}'
echo ""

echo "=== 6. Validation Error: Transaction for Non-Existent Account ==="
curl -s -w "\n" -X POST $BASE_URL/transactions \
  -H "Content-Type: application/json" \
  -d '{"account_id": 999, "operation_type_id": 1, "amount": 50.50}'
echo ""

echo "=== 7. Validation Error: Transaction with Zero Amount ==="
curl -s -w "\n" -X POST $BASE_URL/transactions \
  -H "Content-Type: application/json" \
  -d '{"account_id": 1, "operation_type_id": 1, "amount": 0}'
echo ""

echo "=== 8. Not Found Error: Get Non-Existent Account ==="
curl -s -w "\n" -X GET $BASE_URL/accounts/999
echo ""