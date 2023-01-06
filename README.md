# Veracode Static (Analysis) Go Stub: Confluent Kafka

Veracode Static Analysis does not currently support CGo. [Confluent's Golang Client for Apache KafkaTM](https://github.com/confluentinc/confluent-kafka-go) uses CGo. This means that if your application directly uses this library it will be unable to analyze the application.

As a workaround you can use this stub in place of the library by using a replace when building your package for Veracode Static Analysis, for example like so in a file called `veracode.go.mod`:

```

// should be appended to go.mod before running `go mod vendor`
replace github.com/confluentinc/confluent-kafka-go => ../veracode-static-go-stub-confluent-kafka
```

This can then be used in a file called `veracode.sh` like so:

```bash
#!/bin/sh
# Create /tmp/myapp
mkdir /tmp/myapp
# Copy everything from the current directory into /tmp/myapp
cp -Rf * /tmp/myapp
# Set up stub in /tmp
cd /tmp
git clone git@github.com:relaxnow/veracode-static-go-stub-confluent-kafka.git
# Go to myapp and add the replace to the go.mod
cd /tmp/myapp
cat veracode.go.mod >> /tmp/myapp/go.mod
# Run go mod vendor
go mod vendor
# Zip everything up in myapp.zip
cd /tmp
zip -r myapp.zip myapp
# Clean up - Remove /tmp/myapp
rm -rf /tmp/myapp
# Clean up - Remove stub
rm -rf /tmp/veracode-static-go-stub-confluent-kafka
```

This results in `/tmp/myapp.zip` being created and packaged appropriately. You can them upload this file to Veracode Satic Analysis.

Please note this only deals with the CGo replace, you may still need other things like for example a `veracode.json` for this please see: 

[Go Application Packaging in the Veracode Documentation](https://docs.veracode.com/r/compilation_go)

For an example on how to use this please see: https://github.com/relaxnow/veracode-static-example-go-kafka-stub
