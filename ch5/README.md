## How to install protobuf compiler

#### Mac

```
brew install protobuf
```

#### Linux

```
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

# Unzip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc3

# Move only protoc* to /usr/bin/
sudo mv protoc3/bin/protoc /usr/bin/protoc
```

#### Windows

```
On Windows, we can just copy the executable (.exe) from https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-win32.zip 
to the PATH environment variable.
```
