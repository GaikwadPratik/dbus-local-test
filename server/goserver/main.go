package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/rs/zerolog/log"
)

func main() {
	go hardcoded()
	go generator()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	log.Warn().Msg("Adios!")
}

func hardcoded() error {
	const serviceName = "com.hiveio.vmmanagerhardcoded"
	const objectPath = "/com/hiveio/vmmanagerhardcoded"
	const intro = `
<node>
	<interface name="com.hiveio.vmmanagerhardcoded">
		<method name="CheckHostForMigration">
		<arg name="guestName" direction="in" type="s"/>
		<arg name="cpuxml" direction="in" type="s"/>
		<arg name="result" direction="out" type="b"/>
	</method>
	<method name="RecoverGuest">
		<arg name="guestName" direction="in" type="s"/>
		<arg name="reason" direction="in" type="s"/>
	</method>
	<method name="RecoverUservolume">
		<arg name="guestName" direction="in" type="s"/>
		<arg name="username" direction="in" type="s"/>
	</method>
</interface>` + introspect.IntrospectDataString + `</node> `
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Error().Err(err).Msg("While initiating connection to system bus in initiateServer")
		return err
	}

	iface := VMManagerDbusInterface{}
	conn.Export(iface, objectPath, "com.hiveio.vmmanagerhardcoded")
	conn.Export(introspect.Introspectable(intro), objectPath, "org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName(serviceName, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Error().Err(err).Stack().Msg("While checking name on system bus in initiateServer")
		return err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Error().Err(err).Stack().Interface("reply", reply).Msg("Service name is already taken")
		return errors.New("name already taken")
	}
	log.Info().Msg(fmt.Sprintf("Listening on %s - %s ...", serviceName, objectPath))
	select {}
}

func generator() error {

	const serviceName = "com.hiveio.vmmanager"
	const objectPath = dbus.ObjectPath("/com/hiveio/vmmanager")

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Error().Err(err).Msg("While initiating connection to system bus in initiateServer")
		return err
	}
	iface := VMManagerDbusInterface{}
	err = ExportComHiveioVmmanager(conn, objectPath, iface)
	if err != nil {
		log.Error().Err(err).Stack().Msg("While exporting dbus interface")
		return err
	}

	reply, err := conn.RequestName(serviceName, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Error().Err(err).Stack().Msg("While checking name on system bus in initiateServer")
		return err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Error().Err(err).Stack().Interface("reply", reply).Msg("Service name is already taken")
		return errors.New("name already taken")
	}

	log.Info().Msg(fmt.Sprintf("Listening on %s - %s ...", serviceName, objectPath))
	select {}

}

type VMManagerDbusInterface struct {
}

func (server VMManagerDbusInterface) CheckHostForMigration(guestName string, cpuxml string) (result bool, err *dbus.Error) {
	log.Info().Str("guestName", guestName).Msg("Received request for checking host compatibility")
	return true, nil
}

func (server VMManagerDbusInterface) RecoverGuest(guestName string, reason string) (err *dbus.Error) {
	log.Info().Str("guestName", guestName).Str("reason", reason).Msg("Received request for recovering guest record")
	return nil
}

func (server VMManagerDbusInterface) RecoverUservolume(guestName string, username string) (err *dbus.Error) {
	log.Info().Str("guestName", guestName).Str("username", username).Msg("Received request for recovering guest record")
	return nil
}
