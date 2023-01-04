# confluent-kafka-go-veracode-stub

You can use this stub in place of the [Confluent's Golang Client for Apache KafkaTM](https://github.com/confluentinc/confluent-kafka-go) by using a replace when building your package for Veracode Static Analysis, for example like so in a file called `veracode.go.mod`:

```

// should be appended to go.mod before running `go mod vendor`
replace github.com/confluentinc/confluent-kafka-go => ../confluent-kafka-go-veracode-stub
```

This can then be used in a file called `veracode.sh` like so:

```bash
# Create /tmp/myapp
mkdir /tmp/myapp
# Copy everything from the current directory into /tmp/myapp
cp -Rf * /tmp/myapp
# Set up stub in /tmp
cd /tmp
git clone git@github.com:relaxnow/confluent-kafka-go-veracode-stub.git
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
rm -rf /tmp/confluent-kafka-go-veracode-stub
```

This results in `/tmp/myapp.zip` being created and packaged appropriately. You can them upload this file to Veracode Satic Analysis.

Please note this only deals with the CGo replace, you may still need other things like for example a `veracode.json` for this please see: 

[Go Application Packaging in the Veracode Documentation](https://docs.veracode.com/r/compilation_go)
