const dbus = require("dbus-next")

async function hardcoded() {
  try {
    const serviceName = "com.hiveio.vmmanagerhardcoded";
    const interfaceName = "com.hiveio.vmmanagerhardcoded";
    const objectPath = "/com/hiveio/vmmanagerhardcoded";
    const bus = dbus.systemBus();
    const proxyObject = await bus.getProxyObject(serviceName, objectPath);
    obj = proxyObject.getInterface(interfaceName)
    const c = await obj.CheckHostForMigration("test", "test")
    console.log("hardcoded worked", c);
  } catch (err) {
    console.error(err, "Inside hardcoded")
  }
}

async function generator() {
  try {
    const serviceName = "com.hiveio.vmmanager";
    const interfaceName = "com.hiveio.vmmanager";
    const objectPath = "/com/hiveio/vmmanager";
    const bus = dbus.systemBus();
    const proxyObject = await bus.getProxyObject(serviceName, objectPath);
    proxyObject.getInterface(interfaceName)
    const c = await obj.CheckHostForMigration("test", "test")
    console.log("generator worked", c);
  } catch (err) {
    console.error(err, "Inside generator")
  }
}

async function Main() {
  await Promise.all([
    hardcoded().catch(err => console.error(err, "from hardcoded")),
    generator().catch(err => console.error(err, "from generator"))
  ])
}

Main()
