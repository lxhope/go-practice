module practice/thrift/demo

require (
	git.apache.org/thrift.git v0.0.0-20150427210205-dc799ca07862
	shared v0.0.0
	tutorial v0.0.0
)

replace (
	shared v0.0.0 => ./gen-go/shared
	tutorial v0.0.0 => ./gen-go/tutorial
)
