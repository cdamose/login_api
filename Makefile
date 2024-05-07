include .env
export

.PHONY: openapi
openapi: openapi_http openapi_js


.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh core internal/core/ports ports