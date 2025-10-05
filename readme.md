
# How to

1. Run kafka (docker compose)
2. Create topics 'kyc-request' 'kyc-response'
> kafka-topics --create --topic kyc-responses --bootstrap-server localhost:29092 --partitions 1 --replication-factor 1
> kafka-topics --create --topic kyc-request --bootstrap-server localhost:29092 --partitions 1 --replication-factor 1
3. Run Go project
4. Submit new message to kyc-request kafka topic
> {"temp_id":0,"perm_id":0,"properties":{"Name":"John","Surname":"Peterson","DOB":"2025-01-01","SSN":"12345-12345"}}