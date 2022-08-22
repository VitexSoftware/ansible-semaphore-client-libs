deffile = ansible-semaphore-openapi3.yaml
outdir = clients

get-def:
	curl "https://converter.swagger.io/api/convert?url=https://ansible-semaphore.com/api/api-docs.yml" -H "Accept: application/yaml" -o ${deffile}

install-dev-deps:
	wget -c https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh -O openapi-generator-cli
	chmod u+x openapi-generator-cli
	./openapi-generator-cli version

generate: java php python ruby golang

java:
		mkdir -p ${outdir}/libjava-semaphore-client
		./openapi-generator-cli generate -i ${deffile} -g java --additional-properties artifactId=semaphore -o ${outdir}/libjava-semaphore-client

php:
		mkdir -p ${outdir}/libphp-semaphore-client
		./openapi-generator-cli generate -i ${deffile} -g php --additional-properties namespace=semaphore -o ${outdir}/libphp-semaphore-client

python:
		rm -rf ${outdir}/libpython-semaphore-client
		./openapi-generator-cli generate -i ${deffile} -g python  -t templates/python --config=templates/python/config.yml --additional-properties=debugSupportingFiles,namespace=semaphore,generatedBy=Vitex,packageName=semaphore_client,projectName=semaphore_client, -o ${outdir}/libpython-semaphore-client

ruby:
		mkdir -p ${outdir}/libruby-semaphore-client
		./openapi-generator-cli generate -i ${deffile} -g ruby --additional-properties namespace=semaphore -o ${outdir}/libruby-semaphore-client

golang:
		mkdir -p ${outdir}/libgolang-semaphore-client
		./openapi-generator-cli generate -i ${deffile} -g go --additional-properties namespace=semaphore -o ${outdir}/libgolang-semaphore-client


