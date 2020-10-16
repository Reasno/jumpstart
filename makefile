default:
		protoc --proto_path=. \
		--proto_path=`go list -m -f "{{.Dir}}" github.com/envoyproxy/protoc-gen-validate` \
		--proto_path=`go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2`/third_party/googleapis \
		--proto_path=`go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2` \
		--openapiv2_out ./doc \
		--openapiv2_opt logtostderr=true \
		--validate_out="lang=go:." \
		app.proto
		truss app.proto
clean:
		rm -rf protoc-gen-openapiv2
		rm -rf protoc-gen-validate
		rm -rf google
