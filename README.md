dbus-codegen-go -camelize -system -package=main -output=goserver/server.go -server-only service.xml
dbus-codegen-go -camelize -system -package=main -output=goclient/client.go -client-only service.xml