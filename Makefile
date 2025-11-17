BUILD_DEST_DIR ?= build
GIT_TAG=$(shell git rev-parse --abbrev-ref HEAD | grep HEAD >> /dev/null && git describe --tags || git rev-parse --abbrev-ref HEAD)
GIT_HASH=$(shell git rev-parse --short HEAD)
GIT_COMMIT_COUNT=$(shell git log --oneline|wc -l|sed s/[[:space:]]//g)
TIME=$(shell date -d today +"%Y%m%d%H%M%S")

VERSION:=$(shell git rev-parse --abbrev-ref HEAD | grep HEAD >> /dev/null && git describe --tags || echo "`git rev-parse --abbrev-ref HEAD`.${GIT_HASH}.${GIT_COMMIT_COUNT}.${TIME}")

.PHONY: install-goctl-openapi3
install-goctl-openapi3:
	go install github.com/wumitech-com/goctl-openapi3@latest

.PHONY: swagger
swagger: genFile := gen/swagger/operation.swagger.json
swagger:
	#goctl api format -dir apis
	# 生成 swagger.json 文件
	goctl api plugin -plugin goctl-openapi3="openapi -filename $(genFile)" -api apis/apis.api -dir .
	# 配置认证信息
	sed -i '' s'/^  "components": {/  "components": {\n    "securitySchemes": {"bearerAuth": {"type": "http","scheme": "bearer","bearerFormat": "JWT"}},/' gen/swagger/operation.swagger.json
	sed -i '' s'/^  "components": {/  "security": [{"bearerAuth":[]}],\n  "components": {/' gen/swagger/operation.swagger.json

.PHONY: swagger-gnu
swagger-gnu: genFile := gen/swagger/operation.swagger.json
swagger-gnu:
	goctl api format -dir apis
	goctl api plugin -plugin goctl-openapi3="openapi -filename $(genFile)" -api apis/apis.api -dir .
	sed -i s'/^  "components": {/  "components": {\n    "securitySchemes": {"bearerAuth": {"type": "http","scheme": "bearer","bearerFormat": "JWT"}},/' gen/swagger/operation.swagger.json
	sed -i s'/^  "components": {/  "security": [{"bearerAuth":[]}],\n  "components": {/' gen/swagger/operation.swagger.json


.PHONY: gen-go
gen-go:
	goctl api go --api apis/apis.api --dir . --remote git@github.com:wumitech-com/goctl-template.gitkam

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ${BUILD_DEST_DIR}/cancer-operation-api operation.go

.PHONY: model
model:
	goctl model mysql ddl --database ops -d "sql/model" --src "sql/ops/*.sql" -i ""
	goctl model mysql ddl --database report -d "sql/model" --src "sql/report/*.sql" -i ""
	goctl model mysql ddl --database cancer -d "sql/model" --src "sql/cancer/*.sql" -i ""
	goctl model mysql ddl --database sdkcancer -d "sql/sdkmodel" --src "sql/sdkcancer/*.sql" -i ""
	goctl model mysql ddl --database sdkcancer -d "sql/sdkmodel" --src "sql/dllpgdevents/*.sql" -i ""
	goctl model mysql ddl --database sdkh5jobs -d "sql/sdkmodel" --src "sql/dllpgdh5jobs/*.sql" -i ""
	goctl model mysql ddl --database sdkmain -d "sql/sdkmodel" --src "sql/dllpgdmain/*.sql" -i ""
	goctl model mysql ddl --database keywords_gg -d "sql/model" --src "sql/keywords_gg/*.sql" -i ""
	goctl model mysql ddl --database h5events -d "sql/sdkmodel" --src "sql/h5events/*.sql" -i ""
	goctl model mysql ddl --database proxycenter -d "sql/model/proxycenter" --src "sql/proxycenter/*.sql" -i ""

.PHONY: protoc
protoc:
	protoc -I . --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/dllpgd/common.proto
	protoc -I . --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative ./proto/signalingInternal/signaling_internal.proto




# self
# 数据库
#ps:   goctl model mysql ddl --database wfw_user -d "sql/model" --src "sql/users/*.sql" -i ""


#接口生成
#ps: goctl api go --api apis/apis.api --dir .