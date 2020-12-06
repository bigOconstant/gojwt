# Golang jwt service

Unhappy with python as it grows to larger sizes, this is a rewrite of my auth microservice in go. In go things mostly make since. 


## Requested and optional fields

**email**, **password**, **username** will be required fields for each user. There is another optional data field called data. Stored as josonb in postgres, can be retrieved with each token. Likely usage scenarios, store permissions/role , first, lastname. Small amounts of data because it will be sent on each Auth request. 