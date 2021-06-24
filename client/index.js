const dbus = require("dbus-next")

async function hardcoded() {
  try {
    const serviceName = "com.hiveio.vmmanagerhardcoded";
    const interfaceName = "com.hiveio.vm.Managerhardcoded";
    const objectPath = "/com/hiveio/vmmanagerhardcoded";
    const bus = dbus.systemBus();
    const proxyObject = await bus.getProxyObject(serviceName, objectPath);
    proxyObject.getInterface(interfaceName)
    console.log("hardcoded worked");
  } catch (err) {
    console.error(err, "Inside hardcoded")
  }
}

async function generator() {
  try {
    const serviceName = "com.hiveio.vmmanager";
    const interfaceName = "com.hiveio.vm.Manager";
    const objectPath = "/com/hiveio/vmmanager";
    const bus = dbus.systemBus();
    const proxyObject = await bus.getProxyObject(serviceName, objectPath);
    proxyObject.getInterface(interfaceName)
    console.log("generator worked");
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
